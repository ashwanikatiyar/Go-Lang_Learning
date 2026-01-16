package main

import (
	"fmt"
	"sync"
)

/*
4.Goroutines Need Coordination
	-Starting goroutines is easy.
	-Waiting for them is the real problem.
 sync.WaitGroup (MOST COMMON TOOL)
	Example
			var wg sync.WaitGroup

			for i := 0; i < 5; i++ {
				wg.Add(1)
				go func(i int) {
					defer wg.Done()
					fmt.Println(i)
				}(i)
			}

			wg.Wait()
Rules
*Add() before starting goroutine
*Done() exactly once
*Wait() blocks

*/

func task(id int , wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("Doing task" , id)
}


func main() {
	var wg sync.WaitGroup

	for i:=0 ; i<10; i++{
		wg.Add(1)
		go task(i, &wg)
	}
	wg.Wait()

}