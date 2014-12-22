package main 

import(
	"log"
	"Cohesion/server/server"
	"Cohesion/server/spider"
)

func main() {
	log.Printf("server starting...")
	spider.InitSegment()
	server.DatabaseInit()
	server.Run()
}