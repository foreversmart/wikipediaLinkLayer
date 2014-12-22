package server

import (
	"net/http"
	"strconv"
	"encoding/json"
	"fmt"

)

func createId(w http.ResponseWriter, r *http.Request){
	if r.Method == "GET"{
		user := &User{}
		user.Init()
		id, ok := user.insert()
		if ok {
			//network wrong
			result := map[string]string{"result": strconv.Itoa(id)}
			strResult,_ := json.Marshal(result)
			fmt.Fprintf(w, string(strResult))
		}else {
			//sql wrong
			result := map[string]string{"result": "sql wrong!"}
			strResult,_ := json.Marshal(result)
			fmt.Fprintf(w, string(strResult))
		}

	}else{
		//network wrong
		result := map[string]string{"result": "net work wrong!"}
		strResult,_ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	}
}


func pullMap(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST"{
		id := r.FormValue("id")
		var word Word
		words := word.QueryUserWord(id)
		var tempMap = make(map[string]string)
		for _, value := range words {
			tempMap[value.contents["word"]] = value.contents["weight"]
		}
		result := map[string]interface{}{"result": tempMap}
		strResult,_ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	}else{
		//network wrong
		result := map[string]string{"result": "1"}
		strResult,_ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	}
}
