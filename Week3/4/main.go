package main

import "fmt"

func main() {
	fmt.Print("input : ")
	var name string
	var age int
	var height float32
	var height float64
	n, err := fmt.Scan(&name, &age, &weigth, &heihgt)
	fmt.Println(name, age, weigth, height)
	fmt.Println(`number of argument `, n)
	fmt.Println(`error `, err)
}
