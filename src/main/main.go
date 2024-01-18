package main

import (
	"fuck-go/src/main/db"
	"fuck-go/src/main/routers"
)

func main() {
	db.ConnectMysql()
	routers.CreateRouter()
}
