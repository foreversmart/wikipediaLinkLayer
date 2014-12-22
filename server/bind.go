package server


func bind(){
	//search bind
	mymux.HandleFunc("/searchkey", searchKey)
    mymux.HandleFunc("/pullkey", pullKey)

    //user map bind
    mymux.HandleFunc("/pullmap", pullMap)
    mymux.HandleFunc("/createkey", createId)
}