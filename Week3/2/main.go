package main

import "fmt"

func main() {
	n, e := fmt.Print("Hello", "World", 123, 456, "\n")
	fmt.Println("number of bytes written :", n, "\n")
	fmt.Println("write erro encountered :", e, "\n")

}
