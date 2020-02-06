package main

import (
	"fmt"
)

func main() {
	fmt.Println("****~~~~ A Simple Calculator built using Go lang ~~~~****")
	fmt.Println("Enter the operation you want to use:\n[a]Addition\n[s]Subtraction\n[m]Multiplication\n[d]Divivion\n:")
	var option string
	fmt.Scanf("%s", &option)
	switch option {
		case "a":
			num1, num2 := getUserInputs()
			fmt.Printf("The result is %f\n", num1 + num2)
		case "s":
			num1, num2 := getUserInputs()
			fmt.Printf("The result is %f\n", num1 - num2)
		case "m":
			num1, num2 := getUserInputs()
			fmt.Printf("The result is %f\n", num1 * num2)
		case "d":
			num1, num2 := getUserInputs()
			fmt.Printf("The result is %f\n", num1 / num2)
		default:
			fmt.Println("Invalid option")
	}
}

func getUserInputs() (float32, float32) {
	var (
		num1 float32
		num2 float32
	)
	fmt.Println("Enter first number: ")
	_, err1 := fmt.Scanf("%f", &num1)
	if err1 != nil {
		panic("Invalid Input")
	}
	fmt.Println("Enter first number: ")
	_, err2 := fmt.Scanf("%f", &num2)
	if err2 != nil {
		panic("Invalid Input")
	}
	return num1, num2
}