package server

import (
	"database/sql"
	"fmt"
	"log"
)

type Page struct {
	contents map[string]string
}

func (page *Page) Init() {
	page.contents = make(map[string]string)
	page.contents["id"] = ""
	page.contents["name"] = ""
	page.contents["link"] = ""
	page.contents["keyword"] = ""
}

func (page Page) QueryAll() map[int] *Page {

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

    var result map[int] *Page
    result = make(map[int] *Page)
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
        var page Page
        page.Init()
        for i, col := range values {
            // Here we can check if the value is nil (NULL value)
            if col == nil {
                value = ""
            } else {
                value = string(col)
            }
            // fmt.Println(columns[i], ": ", value)
            page.contents[columns[i]] = value
        }
        // fmt.Println("-----------------------------------")
        result[index] = &page
        index = index + 1
    }
    return result
}

func (page Page) QueryId(id string) Page {

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
    var temp Page
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

func (page Page) QueryPage(name string) *Page {

    // Execute the query
    if db == nil {
        fmt.Println("db is null")
    }
    rows, err := db.Query("SELECT * FROM " + userTable + " where name = ? ", name)
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
    var temp Page
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
    return &temp
}

func (page *Page) insert() bool {
	stmt, err := db.Prepare("INSERT INTO page (id, name, link, keyword)" + 
		" VALUES(?, ?, ?, ?)")
	defer stmt.Close()
	checkErr(err)

    var index interface {}
    if page.contents["id"] == "" || page.contents["id"] == "NULL" {

    }else {
        index = page.contents["id"]
    }
	_, err = stmt.Exec(index, page.contents["name"], page.contents["link"], page.contents["keyword"])
	checkErr(err)
    if err != nil {
        fmt.Println(err)
        return false
    }
    return true
}

func (page *Page) delete(){
	_, err := db.Exec("DELETE FROM page where id = ?", page.contents["id"])
	if err != nil {
		log.Println(err)
		return
	}
}

func (page *Page) update() bool {
	stmt, err := db.Prepare("update page set name=?, link=?, keyword=? where id = ?")
    checkErr(err)
    
    var index interface {}
    if page.contents["id"] == "" || page.contents["id"] == "NULL" {

    }else {
        index = page.contents["id"]
    }
    
    _, err = stmt.Exec(page.contents["name"], page.contents["link"], page.contents["keyword"], index)

    checkErr(err)
    if err != nil {
        fmt.Println(err)
        return false
    }
    return true
}


