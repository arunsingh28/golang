package main

import "fmt"

func main() {
	var username string = "Arun"
	var age int = 25
	fmt.Print("My name is ", username, " and my age is ", age, "\n")
	fmt.Printf("Type of username is %T and type of age is %T", username, age)

	var isLoggedIn bool = true
	fmt.Println("\nIs user logged in? ", isLoggedIn)

	var smallVal uint8 = 255
	var smallVal1 int = int(smallVal)
	smallVal1 = 5992
	var floatValue float32 = float32(smallVal1)
	floatValue = 455.21

	fmt.Println("Max value of unit8 is ", smallVal)
	fmt.Printf("Type of smallVal is %T", smallVal)
	fmt.Println("new values of smallVal1 ", smallVal1)
	fmt.Println("new values of floatValue ", floatValue)

	// default value
	var defaultStr string
	fmt.Print("default value of defaultStr", defaultStr, "\n")
	fmt.Printf("type of defaultStr is %T \n", defaultStr)

	// no var style

	withoutVar := "Arun Pratap Singh"
	fmt.Print("withoutVar value is ", withoutVar, "\n")
}
