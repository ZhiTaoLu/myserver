package mytcp

import (
	"net"
	"log"
	"time"
)



func Server() {
	listener, err := net.Listen("tcp", ":2233")
	if err != nil {
		log.Println("error listen:", err)
		return
	}
	defer listener.Close()
	log.Println("listen ok")

	var i int
	for {
		//time.Sleep(time.Second * 10)
		if conn, err := listener.Accept(); err != nil {
			log.Println("accept error:", err)
			break
		} else {
			go clientHandle(conn)
		}
		i++
		log.Printf("%d: accept a new connection\n", i)
	}
}

//单独处理客户端的请求
func clientHandle(conn net.Conn) {
	//设置当客户端3分钟内无数据请求时，自动关闭conn
	conn.SetReadDeadline(time.Now().Add(time.Minute * 1))
	defer conn.Close()

	//循环的处理客户的请求
	for {
		data := make([]byte, 256)
		//从conn中读取数据
		n, err := conn.Read(data)
		//如果读取数据大小为0或出错则退出
		if n == 0 || err != nil {
			break
		}

		log.Printf("recv client msg:%v", string(data[0:n]))

		//发送给客户端的数据
		rep := "hello,client \r\n"

		//发送数据
		_, err2 := conn.Write([]byte(rep))
		if err2 == nil {
			log.Printf("send client msg:%v", rep)
		}
	}
}
