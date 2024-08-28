package main

import (
	"awesomeProject1/server"
	"log"
)

// アップグレーダーの設定

func main() {
	log.Fatalf("%+v", server.Run())
}
