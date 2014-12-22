package main  
  
import (  
    "fmt"  
    "net/http"
    "net/url" 
    "io/ioutil"
)  
  
func main() {  
    httpSearchKey()
    // httpPostFetchFavor()
    // httpPostFetchComment()
    // httpPostCheckId()
}

func httpSearchKey() {
    resp, err := http.PostForm("http://127.0.0.1:1281/searchkey",
        url.Values{"id": {"1"}, "key": {"软件"}})
    if err != nil {
        fmt.Println(err) 
    }
 
    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        // handle error
    }
    fmt.Println(string(body))
    fmt.Println("tail")

}
