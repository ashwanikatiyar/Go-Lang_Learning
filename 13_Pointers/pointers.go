package main

import "fmt"



//!~Rule #1
// * means different things depending on context
// In a type: *int → pointer to int
// In code: *p → value at that address


//!~Rule #2
// Go passes everything by value — pointers let you share access

//!~Rule #3
// Never dereference a nil pointer
// *p = 10 // ❌ panic

//!~Rule #4
// Use pointers for structs when you want mutation or efficiency

func changeWithoutPointers(x int) int{
	x = 22
	fmt.Println("Changed variable without using Pointers: ",x)
	return x
}

func changeWithPointers(x *int) int{
	*x = 22
	fmt.Println("Changed variable without using Pointers: ",*x)
	return *x

}


//Pointers with Struct 
type User struct{
	Name string
	Age int
	email string
}
func renameWithoutPointers(u User){
	u.Name = "ABC"
}
func renameWithPointers(u *User){
	u.Name = "XYZ"
}

func main() {
	num := 10
	fmt.Println(changeWithoutPointers(num))
	fmt.Println("Num : ",num) //Not actually changed ,just the copy got changed

	p := &num
	fmt.Println(changeWithPointers(p))
	fmt.Println("Num : ",num) //Got actually changed using the memory address


	//Pointers with struct
	var u User
	u.Name = "ABC"
	fmt.Println("Name before rename : ", u.Name)
	renameWithoutPointers(u)
	fmt.Println("Name after rename : ", u.Name)
	renameWithPointers(&u)
	fmt.Println("Name after Pointer rename : ", u.Name)



	//Common mistake done
	nums := make([]int , 6 )
	for _ , v:= range nums{
		p:= &v
		fmt.Println(*p) 
	}

	for i:= range nums{
		p := &nums[i]
		fmt.Println(*p)
	}
	/*
	Both snippets print the values correctly in your example, but they handle memory addresses very differently. The first one takes the address of the temporary loop variable, while the second one takes the address of the actual slice element.
	*/
}