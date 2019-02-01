package shared

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"youtubeCrawler/models"
)

// MyDB struct that holds config data and pointer to *sql.DB
type MyDB struct {
	User   string
	Pwd    string
	DbUrl  string
	DbName string
	DbPool *sql.DB
}

//TODO: load data from config and assign them to the struct
func init() {

}

// openConnection opens connection to db
func (db *MyDB) OpenConnection() error {
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
func (db *MyDB) NewJobId() (int, error) {
	//TODO implement function logic
	return -1, nil
}

//InsertNextVideoData inserts data to the DB taken from <-chan c
func (db *MyDB) InsertNextVideoData(c <-chan structs.VideoData) error {
	//TODO implement function logic
	return nil
}

func (db *MyDB) TestInsertSuffixUrl(c chan structs.VideoData) error {
	stmt, err := db.DbPool.Prepare("insert into videoData (`suffix`, `title`) values (?,?)")
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
	}

	return nil
}
