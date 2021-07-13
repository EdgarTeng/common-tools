package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
)

var (
	hFlag string
	pFlag int
	tFlag int
)

func init() {
	flag.StringVar(&hFlag, "s", "localhost", "server host(or IP)")
	flag.IntVar(&pFlag, "p", 8080, "port")
	flag.IntVar(&tFlag, "t", 1, "threads number")
	flag.Parse()
}

func main() {
	println("host: ", hFlag)
	println("port: ", pFlag)
	println("threads: ", tFlag)

	wg := new(sync.WaitGroup)
	addr := fmt.Sprintf("%s:%d", hFlag, pFlag)
	for i := 0; i < tFlag; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c := connect(addr)
			readAndWrite(c)
		}()
	}
	wg.Wait()
	log.Println("all finished")

}

func connect(addr string) net.Conn {
	c, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%s connect ok\n", addr)
	return c

}

func readAndWrite(c net.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">> ")
		text, _ := reader.ReadString('\n')
		fmt.Fprintf(c, text+"\n")

		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("->: " + message)
		if strings.TrimSpace(text) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}

	}
}
