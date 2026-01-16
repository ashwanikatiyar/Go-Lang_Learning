package main

import "fmt"

/*Generics allow you to parameterize types, just like functions parameterize values.

Eg
		func Print[T any](value T) {
			fmt.Println(value)
		}

*/

//Basic Example
func showSlice[T int | string](s []T){
	for _ , v := range s{
		fmt.Println(v)
	}
}

//Eg 2
//Real world Stack implementation
type Stack[T any] struct{
	items []T
}

func (s *Stack[T]) Push(item T ){
	s.items = append(s.items , item)
}

func (s *Stack[T])Pop()T{
	last := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return last
}


func main() {
	//Eg1
	//Now you are not bound to send slice of only type.
	nums:= []int{1 ,2 ,3 ,4}
	showSlice(nums)

	languages := []string{"Go-Lang" ,"Javascript" , "TypeScript" , "C++"}
	showSlice(languages)


	//Eg2
	intStack := Stack[int]{}
	intStack.Push(10)
	intStack.Push(20)
	intStack.Push(30)

	fmt.Println(intStack)

	intStack.Pop()
	fmt.Println(intStack)


	//You may try this for a string stack as well
	stringStack := Stack[string]{}
	stringStack.Push("Hi")
	stringStack.Push("am")
	stringStack.Push("Ashwani")

	fmt.Println(stringStack)
}

