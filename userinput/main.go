package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var a, b int
	fmt.Println("Enter two numbers")
	fmt.Scan(&a, &b)
	fmt.Println("Sum of two numbers is", a+b)

	// read whole line
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Your full name: ")
	name, _ := reader.ReadString('\n')
	fmt.Println("Your name is", strings.TrimSpace(name))

}
