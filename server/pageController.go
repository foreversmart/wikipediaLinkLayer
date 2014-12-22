package server

import(
	"bytes"
	"strconv"
	"strings"
	"Cohesion/server/spider"
)

type PageModel struct{
	name string
	keyword []string
	link *spider.ValSorter
}

func (pageMode *PageModel)GetDataBaseModel() *Page {
	var page = &Page{}
	page.Init()
	page.contents["name"] = pageMode.name
	page.contents["keyword"] = getDataArray(pageMode.keyword)
	page.contents["link"] = getDataMap(pageMode.link)
	return page
}

func (page *Page)GetPageModel() *PageModel {
	var pm = &PageModel{}
	pm.name = page.contents["name"]
	pm.keyword = getArray(page.contents["keyword"])
	pm.link = getMap(page.contents["link"])
	return pm
}

func getDataArray(keyword []string) string{
	tempContent := bytes.Buffer{}
	count := 0
	for _,value := range keyword{
		tempContent.WriteString(value)
		count++
		if count>= len(keyword){

		}else{
			tempContent.WriteString(",")			
		}
	}
	return tempContent.String()
}

func getDataMap(link *spider.ValSorter) string{
	tempContent := bytes.Buffer{}
	count := 0
	for key, value := range link.Keys{
		tempContent.WriteString(value)
		tempContent.WriteString(":")
		tempContent.WriteString(strconv.Itoa(link.Vals[key]))
		if count>= len(link.Keys){

		}else{
			tempContent.WriteString(",")			
		}
	}
	return tempContent.String()
}

func getArray(keyword string) []string{
	return strings.Split(keyword, ",")
}

func getMap(link string)  *spider.ValSorter{
	array := strings.Split(link, ",")
	vs := &spider.ValSorter{
        Keys: make([]string, 0, 10),
        Vals: make([]int, 0, 10),
    }
	for _, value := range array{
		temp:=strings.Split(value, ",")
		vs.Keys = append(vs.Keys, temp[0])
		tempInt,_ := strconv.ParseInt(temp[1], 10, 64)
        vs.Vals = append(vs.Vals, int(tempInt))
	}
	return vs
}

