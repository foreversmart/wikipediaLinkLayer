package spider

import (
	"github.com/wangbin/jiebago"
	"github.com/wangbin/jiebago/analyse"
	"regexp"
	// "strings"
)

func InitSegment() {
	jiebago.SetDictionary("../../github.com/wangbin/jiebago/dict.txt")
	analyse.SetIdf("../../github.com/wangbin/jiebago/analyse/idf.txt")
}

func Parse(sentence string) []string{
	 // 设定字典
	// fmt.Printf("【精确模式】: %s\n\n", strings.Join(jiebago.Cut(sentence, false, true), "/ "))
	result := jiebago.CutForSearch(sentence, true)
	return result
}

func GetKeyWord(sentence string) []string {
	result := analyse.ExtractTags(sentence, len(sentence)/20)
	reg := regexp.MustCompile("^[0-9.\t\r\n]+$")
	var newtemp = make([]string, 0, 10)
	for _, value := range result{
		if reg.MatchString(value) {
			//数字
		}else{
			newtemp = append(newtemp, value)
		}
	}
	return newtemp
}