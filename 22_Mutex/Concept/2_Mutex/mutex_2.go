package main

//!~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

import (
	"fmt"
	"sync"
)

type Post struct {
	view int
}

func (p *Post) inc(wg *sync.WaitGroup) {
	defer wg.Done()
	p.view += 1
}

func main() {

	myPost := Post{
		view: 0,
	}
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg)
	}

	wg.Wait()
	fmt.Println(myPost.view)

}

/*Output :

Sometimes : 100
Sometimes : 95
Sometimes : 82
.
.
*/

//!~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

//*2nd Run

// type Post struct {
// 	views int
// 	mu    sync.Mutex
// }

// func (p *Post) inc(wg *sync.WaitGroup) {

// 	defer func() {
// 		p.mu.Unlock()
// 		wg.Done()
// 	}()

// 	p.mu.Lock()
// 	p.views += 1

// }

// func main() {

// 	myPosts := Post{
// 		views: 0,
// 	}

// 	var wg sync.WaitGroup

// 	for i := 0; i < 100; i++ {

// 		wg.Add(1)
// 		go myPosts.inc(&wg)
// 	}

// 	wg.Wait()
// 	fmt.Println(myPosts.views)

// }

/*
Output : Always same
100
*/
