package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func onMessage(conn net.Conn) {
	for {
		reader := bufio.NewReader(conn)
		msg,_ := reader.ReadString('\n')
		fmt.Println(msg)
	}
}
func main()  {
	conn, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(conn)
	fmt.Print("Your name:")
	nameReader := bufio.NewReader(os.Stdin)
	nameInput, _ := nameReader.ReadString('\n')
	nameInput = nameInput[:len(nameInput) -1]
	fmt.Println("============= message ==================")

	go onMessage(conn)
	for {
		msgReader := bufio.NewReader(os.Stdin)
		msg, err := msgReader.ReadString('\n')
		if err != nil {
			break
		}

		msg = fmt.Sprintf("%s: %s\n", nameInput, msg[:len(msg)-1])
		conn.Write([]byte(msg))

	}

	conn.Close()


	fmt.Println(nameInput)
}
