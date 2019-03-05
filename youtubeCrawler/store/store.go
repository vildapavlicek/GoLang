package store

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"youtubeCrawler/models"
)

// DbStore struct that holds config data and pointer to *sql.DB

type Manager struct {
	StorePipe        chan models.NextLink
	StoreDestination Storer
}

type Storer interface {
	store(link models.NextLink) error
	Close()
}

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

//TODO: load data from config and assign them to the struct
func init() {

}

func New() *Manager {
	storeChan := make(chan models.NextLink, 500)
	storeDestination, err := decideStoreTarget()
	if err != nil {
		fmt.Printf("Failed to resolve store destination. Reason: %s", err)
		return nil
	} else {
		return &Manager{
			StorePipe:        storeChan,
			StoreDestination: storeDestination,
		}
	}

}

func decideStoreTarget() (Storer, error) {
	db := DbStore{
		User:   "root",
		Pwd:    "1111",
		DbUrl:  "tcp(127.0.0.1:3306)",
		DbName: "testdb",
	}

	err := db.OpenConnection()

	if err == nil {
		return db, nil
	}

	file, err := os.Create("links.dat")
	if err != nil {
		file.Close()
		return nil, err
	}
	thepath, err := filepath.Abs(filepath.Dir(file.Name()))
	fmt.Printf("Created file at '%v'\n", thepath)
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

//NewJobID inserts new job and returns its ID
func (db *DbStore) NewJobId() (int, error) {
	//TODO implement function logic
	return -1, nil
}

//InsertNextVideoData inserts data to the DB taken from <-chan c
func (db *DbStore) InsertNextVideoData(c <-chan models.NextLink) error {
	//TODO implement function logic
	return nil
}

func (db DbStore) store(link models.NextLink) error {
	//TODO implement
	/*stmt, err := db.DbPool.Prepare("insert into videoData (`suffix`, `title`) values (?,?)")
	if err != nil {
		log.Printf("Failed to prepare stmt %s", err)
	}
	for data := range c {
		insertId, err := stmt.Exec(data.Link, data.Title)
		if err != nil {
			log.Printf("Insert failed: %s", err)
		}

		id, err := insertId.LastInsertId()
		if err != nil {
			log.Printf("Failed to get insert id: %s", err)
		}
		fmt.Printf("Inserted link %v; title: %v; with id %v\n", data.Link, data.Title, id)
	}*/

	return nil
}

func (db DbStore) Close(){
	db.DbPool.Close()
}

func (f FileStore) store(link models.NextLink) error {
	s := "[ID: '" + link.Id + "', Link: '" + link.Link + "', Title: '" + link.Title + "', no.: '" + strconv.Itoa(link.Number) + "']\n"
	_, err := f.destFile.Write([]byte(s))
	if err != nil {
		return err
	}
	return nil
}

func (m *Manager) StoreData() {

	for {
		select {
		case data := <-m.StorePipe:
			err := m.StoreDestination.store(data)
			if err != nil {
				fmt.Printf("Failed to store data [ID: %v], iteration %v, reason: %s", data.Id, data.Number, err)
			}
		default:

		}

	}
	fmt.Println("Storing data finished")
}

func (f FileStore) Close(){
	f.destFile.Close()
}