package server

import (
	"database/sql"
	"fmt"
	"log"
)

type Word struct {
	contents map[string]string
}

func (word *Word) Init() {
	word.contents = make(map[string]string)
	word.contents["userid"] = ""
	word.contents["word"] = ""
    word.contents["weight"] = ""
}

func (word Word) QueryAll() map[int] *Word {

    // Execute the query
    rows, err := db.Query("SELECT * FROM " + wordTable)
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

    var result map[int] *Word
    result = make(map[int] *Word)
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
        var word Word
        word.Init()
        for i, col := range values {
            // Here we can check if the value is nil (NULL value)
            if col == nil {
                value = ""
            } else {
                value = string(col)
            }
            // fmt.Println(columns[i], ": ", value)
            word.contents[columns[i]] = value
        }
        // fmt.Println("-----------------------------------")
        result[index] = &word
        index = index + 1
    }
    return result
}

func (word Word) QueryWord(userid, wordStr string) *Word {

    // Execute the query
    rows, err := db.Query("SELECT * FROM " + wordTable + " where userid = ? AND word = ?", userid, wordStr)
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
    var temp Word
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
    if  index > 0{
        return &temp
    }else{
        return nil
    }
    
}

func (word Word) QueryUserWord(userid string) map[int] *Word {

    // Execute the query
    rows, err := db.Query("SELECT * FROM " + wordTable + " where userid = ? ", userid)
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

    var result map[int] *Word
    result = make(map[int] *Word)
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
        var word Word
        word.Init()
        for i, col := range values {
            // Here we can check if the value is nil (NULL value)
            if col == nil {
                value = ""
            } else {
                value = string(col)
            }
            // fmt.Println(columns[i], ": ", value)
            word.contents[columns[i]] = value
        }
        // fmt.Println("-----------------------------------")
        result[index] = &word
        index = index + 1
    }
    return result
}

func (word *Word) insert() bool {
	stmt, err := db.Prepare("INSERT INTO word (userid, word, weight)" + 
		" VALUES(?, ?, ?)")
	defer stmt.Close()
	checkErr(err)

	_, err = stmt.Exec(word.contents["userid"], word.contents["word"], word.contents["weight"])
	checkErr(err)
    if err != nil {
        fmt.Println(err)
        return false
    }
    return true
}

func (word *Word) delete(){
	_, err := db.Exec("DELETE FROM word where userid = ? AND word = ?", word.contents["userid"], word.contents["word"])
	if err != nil {
		log.Println(err)
		return
	}
}

func (word *Word) update() bool {
	stmt, err := db.Prepare("update word set weight=? where userid = ? AND word = ?")
    checkErr(err)
    
    _, err = stmt.Exec(word.contents["weight"], word.contents["userid"], word.contents["word"])

    checkErr(err)
    if err != nil {
        fmt.Println(err)
        return false
    }
    return true
}