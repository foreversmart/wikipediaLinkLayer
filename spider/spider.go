package spider

import(
	"bytes"
	"log"
	"net/http"
	"io/ioutil"
	"fmt"
	"strings"
	query "github.com/PuerkitoBio/goquery"
)

const(
	BaseUrl string = "http://zh.wikipedia.org/zh-cn/"
)

func DealPage(word string) (*ValSorter, []string){
	doc, err := newDocumentFromString(getPageBody(BaseUrl + word))
	checkErr(err)
	tempContent := bytes.Buffer{}
	//段落
	doc.Find(".mw-content-ltr p").Each(func(i int, s * query.Selection){
		tempContent.WriteString(s.Text())

	})
	//ul
	doc.Find(".mw-content-ltr ul").Each(func(i int, s * query.Selection){
		pText := s.Parent().Text()
		sText := s.Text()
		if strings.Trim(pText, "\t\n") == strings.Trim(sText, "\t\n") {

		}else {
			log.Printf("parent:%v\n", strings.Trim(pText, "\t\n"))
			log.Printf("ul", s.Text())
			tempContent.WriteString(sText)
		}
	})
	//ol
	doc.Find(".mw-content-ltr ol").Each(func(i int, s * query.Selection){
		
		// log.Printf("line: %v\n", s.Text())
		pText := s.Parent().Text()
		sText := s.Text()
		if strings.Trim(pText, "\t\n") == strings.Trim(sText, "\t\n") {

		}else {
			log.Printf("parent:%v\n", strings.Trim(pText, "\t\n"))
			log.Printf("ul", s.Text())
			tempContent.WriteString(sText)
		}
		// log.Printf("parent:, %v\n", html)
	})
	//link
	var tempMap = make(map[string]int)
	doc.Find(".mw-content-ltr p a").Each(func(i int, s * query.Selection) {
		res:= s.Text()

		if strings.HasPrefix(res, "[") && strings.HasSuffix(res, "]"){

		}else{
			if value, ok := tempMap[res]; ok {
				tempMap[res] = value + 1
			}else{
				tempMap[res] = 1
			}
		}		
	})

	//content effect
	contentArray := Parse(tempContent.String())
	for key, value := range tempMap {
		for _, value1 := range contentArray {
			if key == value1 {
				tempMap[key] = value + 1
			}
		}
	}
	//sort
	sortArray := NewValSorter(tempMap)
	sortArray.Sort()

	return sortArray, GetKeyWord(tempContent.String())
	// log.Printf("keyword:", GetKeyWord(tempContent.String()))

}

func getPageBody(link string) string {
	res, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("body:", string(body))
	return string(body)

}

func newDocumentFromString(content string) (*query.Document, error) {
	buf := bytes.NewBuffer(nil)
	buf.WriteString(content)
	doc, err := query.NewDocumentFromReader(buf)
	if (err != nil) {
		return nil, err
	}
	return doc, nil
}

func checkErr(err error) {
    if err != nil {
    	log.Println(err)
        fmt.Println(err)
        panic(err)
    }
}