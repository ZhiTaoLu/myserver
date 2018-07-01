package main

import (
	"net"
	"log"
	"myserve/mytcp"
)

func main() {
	go client()
	mytcp.Server()
	//time.Sleep(time.Second * 10)
}



func client() {
	conn, err := net.Dial("tcp", ":2233")
	if err != nil {
		log.Printf("dial error: %s", err)
		return
	}
	log.Println("connect to server ok")
	conn.Write([]byte("haha"))

	for  {
		data := make([]byte, 256)
		_, err2 := conn.Read(data)
		if err2 != nil {
			log.Println(err2)
		} else {
			log.Printf("recv server msg:%v",string(data))
		}
	}
}
