package main

import "fmt"

func main() {
	// for i := 1; i <= 10; i++ {
	// 	if i == 5 {
	// 		break
	// 	}
	// 	fmt.Print(i)

	// }

	for i := 1; i <= 10; i++ {
		if i % 2 != 0 {
			continue
		}
		fmt.Print(i)
	}

}
