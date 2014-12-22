package main 

import (
	"Cohesion/server/spider"
	"log"
)

func main(){
	spider.InitSegment()
	link, word := spider.DealPage("软件")
	log.Printf("ling:%v", link)
	log.Printf("words:%v", word)
}