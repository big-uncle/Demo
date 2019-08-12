package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

var block sync.WaitGroup

func main() { //客户端
	flag := make(chan bool) //这里有俩种方式实现，1是使用sync的同步等待，2是使用无缓冲的chan来进行异步等待
	con, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		log.Println("连接服务器失败，即将退出程序", err)
		time.Sleep(3 * time.Second)
		return
	}
	block.Add(1)
	go send(con)
	go revice(con, flag)
	log.Printf("成功连接%s可以进行通信了\n", con.RemoteAddr())
	block.Wait()

	select {
	case <-flag: //这里是负责判断错误的，也可以用sync.WaitGrup
		log.Println("连接服务器失败，即将退出程序", con.RemoteAddr())
		time.Sleep(3 * time.Second)
		return
	}

}

func send(con net.Conn) {
	defer con.Close()
	reader := bufio.NewReader(os.Stdin)
	for {
		line, _ := reader.ReadString('\n') //由于什么都不输入长度就为2，所以在此作出判定长度
		str := strings.Trim(line, " ")     //去除首位的空格
		str = strings.Trim(str, "\n\r")    //也就是没有值的情况不能发送，因为\r\n也是占字节的，（空值不能发送）
		if len(str) > 0 {
			_, err := con.Write([]byte(str))
			if err != nil {
				log.Println("发送失败", err)
				return
			}
		}
	}
}
func revice(con net.Conn, flag chan bool) {
	defer con.Close()

	for {
		buf := make([]byte, 1024) //未填充的数组，他都给你填充默认值0，由于字节数组是数组，所以就是0
		_, err := con.Read(buf)
		if err != nil {
			log.Println("服务器已断开", err)
			flag <- true //这俩种使用一种既可   要么无缓冲的chan，要么异步等待sync.WaitGroup
			block.Done()
			return
		}
		str := strings.Replace(string(buf), string([]byte{0}), "", -1) //因为缓冲的字节切片，那么没有达到最大范围，那么他会默认添加为0，所以此时需要替换掉0
		log.Printf("接受到来自%s的数据为----:%s\n", con.RemoteAddr(), str)
	}
}
