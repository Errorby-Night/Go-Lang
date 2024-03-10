package main

import "fmt"

func main() {
	i := 3
	for i < 6 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 11; j++ {
		if j%2 == 0 {
			fmt.Println(j)
		}
	}
	for {
		fmt.Println("Hola")
		break
	}
	for k := range 8 {
		fmt.Println(k)
	}
}
