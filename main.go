// Rest api program
package main

import "fmt"

func hello() {
	return 
}

func main() {
	var name string = "DC"

	if len(name) != 0 {
		fmt.Println("more than 0")
	}

	fmt.Printf("Hello %s", name)
}
