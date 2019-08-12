package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"strings"
)

func main() { //服务端
	listen, err := net.Listen("tcp", "127.0.0.1:8888")
	defer listen.Close()
	if err != nil {
		log.Println("服务器启动失败", err)
		return
	}
	log.Println("服务器启动成功")

	for {
		con, err2 := listen.Accept()
		if err2 != nil {
			log.Println("连接请求失败")
		}
		log.Printf("一个客户端%s连接成功\n", con.RemoteAddr())
		go send(con)   //发送
		go revice(con) //接受
	}
}

func send(con net.Conn) {
	defer con.Close()
	reader := bufio.NewReader(os.Stdin)
	for {
		line, _ := reader.ReadString('\n')
		str := strings.Trim(line, " ")
		str = strings.Trim(str, "\n\r")
		if len(str) > 0 {
			_, err := con.Write([]byte(str))
			if err != nil {
				log.Println("发送失败", err)
				break
			}
		}
	}

}

func revice(con net.Conn) {

	defer con.Close()
	buf := make([]byte, 1024)
	for {
		_, err := con.Read(buf)
		if err != nil {
			log.Println("接受失败，对方已断开", err)
			return
		}
		str := strings.Replace(string(buf), string([]byte{0}), "", -1)
		log.Printf("接受到来自%s的数据为----:%s\n", con.RemoteAddr(), str)
	}
}
