package main

import (
	"bufio"
	"fmt"
	// "time"
	"omega-gos/syslib"
	"os"
	"strings"
)

func main() {

	fmt.Printf("Hello World!\n")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")

	for {
		// time.Sleep(time.Second)
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)

		if strings.Compare("hi", text) == 0 {
			fmt.Println("hello, Yourself")
		}
		if strings.Compare("reboot", text) == 0 {
			syslib.Restart2(false)
		}
	}
}
