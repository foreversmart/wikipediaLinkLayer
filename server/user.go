package server

import (
	"database/sql"
	"fmt"
	"log"
)

type User struct {
	contents map[string]string
}

func (user *User) Init() {
	user.contents = make(map[string]string)
	user.contents["id"] = ""
}

func (user User) QueryAll() map[int] *User {

    // Execute the query
    rows, err := db.Query("SELECT * FROM " + userTable)
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    // Get column names
    columns, err := rows.Columns()
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    // Make a slice for the values
    values := make([]sql.RawBytes, len(columns))

    // rows.Scan wants '[]interface{}' as an argument, so we must copy the
    // references into such a slice
    // See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
    scanArgs := make([]interface{}, len(values))
    for i := range values {
        scanArgs[i] = &values[i]
    }

    var result map[int] *User
    result = make(map[int] *User)
    var index int
    index = 0
    // Fetch rows
    for rows.Next() {
        // get RawBytes from data
        err = rows.Scan(scanArgs...)
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }

        // Now do something with the data.
        // Here we just print each column as a string.
        var value string
        var user User
        user.Init()
        for i, col := range values {
            // Here we can check if the value is nil (NULL value)
            if col == nil {
                value = ""
            } else {
                value = string(col)
            }
            // fmt.Println(columns[i], ": ", value)
            user.contents[columns[i]] = value
        }
        // fmt.Println("-----------------------------------")
        result[index] = &user
        index = index + 1
    }
    return result
}

func (user User) QueryId(id string) User {

    // Execute the query
    rows, err := db.Query("SELECT * FROM " + userTable + " where id = ? ", id)
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    // Get column names
    columns, err := rows.Columns()
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    // Make a slice for the values
    values := make([]sql.RawBytes, len(columns))

    // rows.Scan wants '[]interface{}' as an argument, so we must copy the
    // references into such a slice
    // See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
    scanArgs := make([]interface{}, len(values))
    for i := range values {
        scanArgs[i] = &values[i]
    }
    var temp User
    index := 0
    // Fetch rows
    for rows.Next() {
        // get RawBytes from data
        err = rows.Scan(scanArgs...)
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }

        // Now do something with the data.
        // Here we just print each column as a string.
        var value string
        temp.Init()
        for i, col := range values {
            // Here we can check if the value is nil (NULL value)
            if col == nil {
                value = ""
            } else {
                value = string(col)
            }
            // fmt.Println(columns[i], ": ", value)
            temp.contents[columns[i]] = value
        }
        // fmt.Println("-----------------------------------")
        index = index + 1
    }
    return temp
}

func (user *User) insert() (int, bool) {
	stmt, err := db.Prepare("INSERT INTO user (id)" + 
		" VALUES(?)")
	defer stmt.Close()
	checkErr(err)

    var index interface {}
    if user.contents["id"] == "" || user.contents["id"] == "NULL" {

    }else {
        index = user.contents["id"]
    }
	result, err1 := stmt.Exec(index)
    id, err2 := result.LastInsertId()
	checkErr(err2)
    if err1 != nil {
        fmt.Println(err1)
        return int(id),false
    }
    return int(id),true
}

func (user *User) delete(){
	_, err := db.Exec("DELETE FROM user where id = ?", user.contents["id"])
	if err != nil {
		log.Println(err)
		return
	}
}