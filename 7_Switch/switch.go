//Normal , Multiple , Type Switch Case

package main


import (
	"fmt"
	"time"
)

func main() {
	//Normal Switch Case

	// i := 5
	// switch i {
	// case 1:
	// 	fmt.Println("One")
	// case 2:
	// 	fmt.Println("Two")

	// case 3:
	// 	fmt.Println("Three")

	// case 4:
	// 	fmt.Println("Four")

	// case 5:
	// 	fmt.Println("Five")
	// }

	//Multiple Condition Switch

	// switch time.Now().Weekday() {
	// 	case time.Saturday, time.Sunday:
	// 		fmt.Println("Its a Weekend")
	// 	default :
	// 		fmt.Println("Its a Workday")
	// }

	//Type Switch

	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("Integer type")
		case string:
			fmt.Println("String type")
		case bool:
			fmt.Println("String type")
		default:
			fmt.Println("Other", t)
		}
	}

	// whoAmI(time.Now())
	whoAmI("Hello")

	/*
	No functional difference: Using "any" or "interface{}" in your code will behave exactly the same.
	Readability: Many developers prefer any because it’s shorter and clearer that the function accepts any type.
	*/

	//One more thing , if you dont want to use t variable you can simply remove it
	whoAmI2 := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("Integer type")
		case string:
			fmt.Println("String type")
		case bool:
			fmt.Println("String type")
		case float64:
			fmt.Println("Float type")
		default:
			fmt.Println("Other")
		}
	}

	whoAmI2(time.Now())


}
