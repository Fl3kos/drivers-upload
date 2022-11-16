package sqlite

import (
	"database/sql"
	"drivers-create/methods/file"
	"drivers-create/methods/log"

	sqlite3 "github.com/mattn/go-sqlite3"
)

type Activities struct {
	db *sql.DB
}

func InsertSqlite(query, database string) error {
	act, err := openDatabase(database)

	if err != nil {
		log.Errorf("Error:%v", sqlite3.ErrError)
		log.Errorf("Error opening database%v", database)
		return err
	}
	db := act.db

	res, err := db.Exec(query)

	if err != nil {
		log.Errorln("Error insert database")
		return err
	}

	log.Debugf("The id: %v", res)

	return nil
}

func openDatabase(database string) (*Activities, error) {
	log.Debugln("Opening the database")
	databaseRoute := file.ReadSqliteFile(database)
	db, err := sql.Open("sqlite3", databaseRoute)

	if err != nil {
		log.Errorf("Error opening the database %v", databaseRoute)
	}

	act := Activities{
		db: db,
	}

	return &act, err
}
