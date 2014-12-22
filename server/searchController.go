package server

import(
	"net/http"
	"encoding/json"
	"fmt"
	"strconv"
	"Cohesion/server/spider"
	"math"
	"bytes"

)

type Sirity struct {
	key string
	value float64
}

func searchKey(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST"{
		id := r.FormValue("id")
		key := r.FormValue("key")
		var word Word
		tempWord := word.QueryWord(id, key)

		if tempWord ==nil {
			//query is not exsit
			word.Init()
			word.contents["weight"] = "1";
			if word.insert() {
				//insert success
				link, keyword := spider.DealPage(key)
				count := 0
				var searchout = make([]*Sirity, 0, 10)
				for key1, value := range link.Keys{
					count ++
					if(count<20){
						_, keyword1 := spider.DealPage(value)
						newSirity := &Sirity{}
						newSirity.key = value
						newSirity.value = float64(link.Vals[key1]) * calcSrity(keyword, keyword1)
						searchout = append(searchout, newSirity)
					}
				}

				for i:=0; i<len(searchout); i++{
					for j:=i; j<len(searchout); j++{
						if(searchout[i].value < searchout[j].value){
							temp := searchout[i]
							searchout[i] = searchout[j]
							searchout[j] = temp
						}
					}
				}
				output := bytes.Buffer{}
				for _,value := range searchout{
					output.WriteString(value.key)
					output.WriteString(",")
					// output.WriteString(strconv.FormatFloat(value.value, 'g', 4, 64)+",")
				}
				result := map[string]string{"result": output.String()}
				strResult,_ := json.Marshal(result)
				fmt.Fprintf(w, string(strResult))
			} else{
				//insert error

				result := map[string]string{"result": "data base error"}
				strResult,_ := json.Marshal(result)
				fmt.Fprintf(w, string(strResult))
			}
		}else{
			//query is exsit
			number, err := strconv.ParseInt(tempWord.contents["weight"], 10, 32)
			checkErr(err)
			tempWord.contents["weight"] =  strconv.Itoa(int(number) + 1);
			if tempWord.insert() {
				//insert success
				result := map[string]string{"result": "1"}
				strResult,_ := json.Marshal(result)
				fmt.Fprintf(w, string(strResult))
			} else{
				//insert error
				result := map[string]string{"result": "data base error"}
				strResult,_ := json.Marshal(result)
				fmt.Fprintf(w, string(strResult))
			}
		}
		
	}else{
		//network wrong
		result := map[string]string{"result": "network wrong"}
		strResult,_ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	}
}


func pullKey(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST"{
		key := r.FormValue("key")
		result := map[string]string{"result": key}
		strResult,_ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
		
	}else{
		//network wrong
		result := map[string]string{"result": "1"}
		strResult,_ := json.Marshal(result)
		fmt.Fprintf(w, string(strResult))
	}
}

func calcSrity(link1 []string, link2 []string) float64{
	link1 = link1[:10]
	link2 = link2[:10]
	var totalup float64 = 0
	for key, value := range link1 {
		for key2, vlaue2 := range link2 {
			if value == vlaue2 {
				totalup = totalup + math.Abs(float64(key-20)) * math.Abs(float64(key2-20))
			}
		}
	}
	return totalup / 101.6857
}
