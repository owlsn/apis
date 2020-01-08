package datasource

import (
	"errors"
	"github.com/owlsn/apis/src/utils/database"
	"github.com/jinzhu/gorm"
)

// Engine is from where to fetch the data, in this case the users.
type Engine uint32

const (
	// Memory stands for simple memory location;
	// map[int64] datamodels.User ready to use, it's our source in this example.
	Memory Engine = iota
	// PostgreSQL compatible source
	PostgreSQL
	// MySQL for mysql-compatible source location.
	MySQL
	// File data
	File
	// Sqlite3 data
	Sqlite3
)

// Load : Load mysql database
func Load(engine Engine) ( *gorm.DB, error) {
	if(engine != MySQL){
		return nil, errors.New("must use mysql as data source")
	}
	db, err := database.Instance()
	if err != nil{
		return nil, err
	}
	return db, nil
}