package database 

import (
    "database/sql"

    _ "github.com/go-sql-driver/mysql" // MySQL connection driver
)

// Open the connection with the database

func Connect() (*sql.DB, error) { 
    stringConnection := "thiago:@/supercars?charset=utf8&parseTime=True&loc=Local" 
    db, error := sql.Open("mysql", stringConnection) 
    
    if error != nil { 
        return nil, error
    }

    if error = db.Ping(); error != nil { 
        return nil, error 
    }

    return db, nil 
}