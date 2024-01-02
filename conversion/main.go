package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	fmt.Print("Welcome to Arun's App")
	fmt.Print("Enter your name: ")
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("Hello", input, "Your rating is ", random())

}

func random() int {
	return rand.Intn(5) + 1
}
