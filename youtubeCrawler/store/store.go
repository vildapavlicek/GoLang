package store

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"youtubeCrawler/config"
	"youtubeCrawler/models"
)

// Manager manages data storing
type Manager struct {
	StorePipe        chan models.NextLink // chan to receive data to store from
	StoreDestination Storer               // destination where to store data, DB or file
	Shutdown         chan bool
}

type Storer interface {
	Store(link models.NextLink) error
	Close()
}

// DbStore holds DB configuration
type DbStore struct {
	User   string
	Pwd    string
	DbUrl  string
	DbName string
	DbPool *sql.DB
}

type FileStore struct {
	destFile *os.File
}

// Returns new *Manager
func New(config config.StoreConfig) *Manager {
	storeDestination, err := decideStoreTarget(config)
	if err != nil {
		fmt.Printf("Failed to resolve store destination. Reason: %s", err)
		panic(err)
		return nil
	} else {
		return &Manager{
			StorePipe:        make(chan models.NextLink, 500),
			StoreDestination: storeDestination,
			Shutdown:         make(chan bool, 1),
		}
	}

}

// Decides target to store data to. If opening connection to DB fails, saves data to file links.dat
func decideStoreTarget(c config.StoreConfig) (Storer, error) {
	db := DbStore{
		User:   c.DbUser,
		Pwd:    c.DbPwd,
		DbUrl:  "tcp(" + c.DbUrl + ")",
		DbName: c.DbName,
	}

	err := db.OpenConnection()

	if err == nil {
		return db, nil
	}

	file, err := os.Create(c.FilePath)
	if err != nil {
		file.Close()
		return nil, err
	}
	path, err := filepath.Abs(filepath.Dir(file.Name()))
	fmt.Printf("Created file at '%v'\n", path)
	return FileStore{destFile: file}, nil
}

// openConnection opens connection to db
func (db *DbStore) OpenConnection() error {
	var err error
	connectionString := db.User + ":" + db.Pwd + "@" + db.DbUrl + "/" + db.DbName
	db.DbPool, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Println("Failed to connect to DB. Reason: ", err)
		return err
	}
	if err = db.DbPool.Ping(); err != nil {
		return err
	}
	return nil

}

//TODO stores data to DB
func (db DbStore) Store(link models.NextLink) error {
	stmt, err := db.DbPool.Prepare("insert into videoData (`suffix`, `title`) values (?,?)")
	if err != nil {
		log.Printf("Failed to prepare stmt %s", err)
	}
	insertId, err := stmt.Exec(link.Link, link.Title)

	if err != nil {
		log.Printf("Insert failed: %s", err)
	}

	id, err := insertId.LastInsertId()
	if err != nil {
		log.Printf("Failed to get insert id: %s", err)
	}
	fmt.Printf("Inserted link %v; title: %v; with id %v\n", link.Link, link.Title, id)

	return nil
}

func (db DbStore) Close() {
	db.DbPool.Close()
}

//TODO store data to file
func (f FileStore) Store(link models.NextLink) error {
	s := "[ID: '" + link.Id + "', Link: '" + link.Link + "', Title: '" + link.Title + "', no.: '" + strconv.Itoa(link.Number) + "']\n"
	_, err := f.destFile.Write([]byte(s))
	if err != nil {
		return err
	}
	return nil
}

// stores data to configured destination
func (m *Manager) StoreData() {

	for {
		select {
		case data, ok := <-m.StorePipe:
			if !ok {
				fmt.Println("Store channel closed, shutting down")
				m.Shutdown <- true
				close(m.Shutdown)
				return
			}
			err := m.StoreDestination.Store(data)
			if err != nil {
				fmt.Printf("Failed to store data [ID: %v], iteration %v, reason: %s", data.Id, data.Number, err)
			}
		default:

		}

	}
	fmt.Println("Storing data finished")
}

func (f FileStore) Close() {
	f.destFile.Close()
}
