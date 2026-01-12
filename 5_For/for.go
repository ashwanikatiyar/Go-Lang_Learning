package main

import "fmt"

func main() {
	//Since there is no "while" keyword in Go , so we use "for"
	//in order implement while loop

	i := 0
	for i < 4 {
		fmt.Println(i)
		i++ //or i+=1 or i = i + 1
	}

	//Infinite loop

	// for {
	// 	fmt.Println(9)
	// }

	// Classic For loop
	for i := 10; i > 0; i-- {
		if i == 6 {
			println("Continuing the loop at i = ", i)
			continue
		}
		if i == 3 {
			println("Breaking the loop at i = ", i)
			break
		}
		println(i)

	}

	//Range function in Go , Range 5 = 0 to 4

	for j:= range 5{
		println("Using Range function  j = ",j)
	}
}
