package main

import "fmt"

func main() {
	var name = "Srijan"
	var req = 131
	var cgpa = 9.09
	fmt.Println("Name: " + name + "\nRegistration Number: 23MCA0" + fmt.Sprintf("%d", req) + "\nCGPA: " + fmt.Sprintf("%.2f", cgpa))
}
