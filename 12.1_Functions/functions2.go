package main

import "fmt"

/*
1 In Go, functions are first-class citizens.
That means a function can be:

-Stored in a variable
-Passed as an argument
-Returned from another function


2️ Why Would We Return a Function?
Returning a function lets you:
✅ 1. Create customized behavior
You can create a function that’s configured once and reused many times.
✅ 2. Hide implementation details
The caller doesn’t need to know how the logic works.
✅ 3. Maintain state without global variables
Closures allow state to persist safely.
✅ 4. Build clean, reusable APIs
Very common in:
Middleware
Callbacks
Validation
Logging
Retry logic
*/

//3️ A Gentle Example (Without Closures Yet)
func makeAdder() func(int ,int)int{
	return func(a int , b int) int {
		return a + b
	}
}

/*
4️⃣ What Is a Closure? (This Is the Key Idea)
A closure is a function that remembers variables from its surrounding scope, even after that scope has finished executing.

Simple definition:
A closure = function + remembered variables

*/

func makeCustomAdder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}//Check Main function


//Lets try another eg for clarity
func counter() func() int {
	count := 0
	return func () int  {
		count ++
		return count
	}
}



func main() {
	add := makeAdder()
	fmt.Println(add(5 , 7)) //12

	//Custom Adder using Closure
	add10 := makeCustomAdder(10)
	add20 := makeCustomAdder(20)

	fmt.Println(add10(20)) // 30
	fmt.Println(add20(30)) // 50

	//Counter example
	c:= counter()
	fmt.Println(c()) //1
	fmt.Println(c()) //2
	fmt.Println(c()) //3
/*
This is because the inner function remember the "count" and increament it.
Lets see Multiple Closure , Independent State
*/

	fmt.Println("Multiple Closure yet Independent State")

	c1 := counter()
	c2 := counter()

	fmt.Println(c1()) //1
	fmt.Println(c1()) //2
	fmt.Println(c2()) //1





}