package main

import (
	"fmt"
	"io"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		fmt.Println("等待接收消息..", conn.RemoteAddr())
		n, err := conn.Read(buf) // 阻塞在这一直等待,
		if err != nil {
			fmt.Println(err)
			if err == io.EOF {
				fmt.Println("客户端已退出..我也退出")
			}
			return
		}
		fmt.Println(string(buf[:n]))
	}
}

func main() {

	listen, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}
	fmt.Println("listen success:", listen.Addr())
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		fmt.Println("accept conn:", conn.RemoteAddr())
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		} else {
			fmt.Println("conn", conn, conn.RemoteAddr())
			fmt.Printf("%T\n", conn)
			go process(conn)
		}

	}

}
