package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	dial, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		return
	}
	defer dial.Close()

	fmt.Println("conn success", dial)

	reader := bufio.NewReader(os.Stdin)
	for {

		cont, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		cont = strings.Trim(cont, " \r\n")
		if cont == "exit" {
			fmt.Println("bye")
			break
		}

		write, err := dial.Write([]byte(cont))
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("write", write)

	}

}
