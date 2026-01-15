package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

const (
	TCP  = "tcp"
	TCP4 = "tcp4"
	UDP  = "udp"
	UDP4 = "udp4"
	IP   = "ip"
	IP4  = "ip4"
)

var openPorts = make(map[int]string)

func main() {
	ScanPorts()
	printMap(openPorts)
}

func ScanPorts() {
	start, end := grabRange()
	for i := start; i <= end; i++ {
		portNum := strconv.Itoa(i)
		conn, err := net.Dial("tcp", "localhost:"+portNum)
		if err != nil {
			fmt.Printf("Port %d: Closed\n", i)
		} else {
			openPorts[i] = "Open"
			fmt.Printf("Port %d: Open\n", i)
			conn.Close()
		}
	}
}

func grabRange() (int, int) {
	input := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a range of Ports (like 80-100): ")
	input.Scan()
	line := input.Text()
	s := strings.Split(line, "-")
	start, _ := strconv.Atoi(s[0])
	end, _ := strconv.Atoi(s[1])
	return start, end
}

func printMap(mapOfPorts map[int]string) {
	fmt.Println("List of Open Ports:")
	fmt.Println("=======================================")
	for k, v := range mapOfPorts {
		fmt.Printf("Port %d is %v\n", k, v)
	}
}
