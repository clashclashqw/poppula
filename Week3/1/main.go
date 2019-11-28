package main

import "fmt"

func main() {
	fmt.Println("My name is %s, I am %d year old\n", "Goku", 18)

	fmt.Println("type = %T \n", 3.14159265359)
	fmt.Println("pi = %f \n", 3.14159265359)
	fmt.Println("pi = %2.2f \n", 3.14159265359)
	fmt.Println("pi = %9.f \n", 3.14159265359)
	fmt.Println("pi = %-9.f \n", 3.14159265359)
	fmt.Println("pi = %09.f \n", 3.14159265359)
	fmt.Println("pi = %09.2f \n", 3.14159265359)
	fmt.Println("pi = %E \n", 3.14159265359)

	fmt.Println("1 > 2 = %t \n", 1 > 2)
	fmt.Println("10 (base 2) = %b \n", 10)
	fmt.Println("10 (base 8) = %o \n", 10)
	fmt.Println("10 (base 10) = %d \n", 10)
	fmt.Println("10 (base 16) = %x \n", 10)
}
