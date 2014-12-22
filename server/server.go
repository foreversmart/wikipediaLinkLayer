package server

import (
	"net/http"
	"log"
)

var mymux *http.ServeMux

func Run() {
	
	mymux = http.NewServeMux()
    //绑定路由
    bind()

    err := http.ListenAndServe(":1281", mymux) //设置监听的端口

    if err != nil {

        log.Fatal("ListenAndServe: ", err)
    }
}

