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
	var sytmRating = random()

	fmt.Println("Hello", input, "Your rating is ", sytmRating)

	fmt.Printf("type of input is %T", input)

	fmt.Printf("New Rating from admin is %T", string(sytmRating)+"Arun")

}

func random() int {
	return rand.Intn(5) + 1
}
