package main

import (
	"fmt"
	"sync"
)

//1st run

// func worker(i int){
// 	fmt.Println("I am worker ", i , "starting the arrangement")
// 	fmt.Println("I am worker ", i , "arrangements are done ")

// }

// func main() {

// 	fmt.Println("Arrangements started ")

// 	for i:=1 ;i<=3;i++{
// 		worker(i)
// 	}

// 	fmt.Println("Party got started")

// }

/*Output

Arrangements started
I am worker  1 starting the arrangement
I am worker  1 arrangements are done
I am worker  2 starting the arrangement
I am worker  2 arrangements are done
I am worker  3 starting the arrangement
I am worker  3 arrangements are done
Party got started

*/

//2nd Run

// func worker(i int){
// 	fmt.Println("I am worker ", i , "starting the arrangement")
// 	fmt.Println("I am worker ", i , "arrangements are done ")

// }

// func main() {

// 	fmt.Println("Arrangements started ")

// 	for i:=1 ;i<=3;i++{
// 		go worker(i)
// 	}

// 	fmt.Println("Party got started")

// }

/*
Output (Since we didnt add time.Sleep() in the main function):

Arrangements started
Party got started
*/

//3rd Run


func worker(i int , wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("I am worker ", i , "starting the arrangement")
	fmt.Println("I am worker ", i , "arrangements are done ")

}


func main() {

	var wg sync.WaitGroup
	fmt.Println("Arrangements started ")

	for i:=1 ;i<=3;i++{
		wg.Add(1)
		go worker(i , &wg)
	}

	wg.Wait()
	fmt.Println("Party got started")

}

/*
Output


Arrangements started 
I am worker  3 starting the arrangement
I am worker  3 arrangements are done 
I am worker  1 starting the arrangement
I am worker  1 arrangements are done 
I am worker  2 starting the arrangement
I am worker  2 arrangements are done 
Party got started
ashwanikatiyar@Ashwanis-MacBook-Air Go Lessons % go run "/Users/ashwanikatiyar/Documents/Lessons/Go Lessons/19_WaitGroups/Concept/concept.go"
Arrangements started 
I am worker  3 starting the arrangement
I am worker  3 arrangements are done 
I am worker  2 starting the arrangement
I am worker  2 arrangements are done 
I am worker  1 starting the arrangement
I am worker  1 arrangements are done 
Party got started


As you can see workers arent in sequence but , 
*/