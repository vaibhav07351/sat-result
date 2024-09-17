package main

import (
	"fmt"
	satcliservice "sat-result/sat-cli-service"
)

func main() {
	fmt.Println()
	fmt.Println("Welcome to the SAT Results Application")
	satcliservice.Run()
	fmt.Println("Thank you for visting SAT Results Application!")
	fmt.Println()
}
