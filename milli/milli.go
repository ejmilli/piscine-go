// practice channles and waitgroups

package main

import (
	"fmt"
	"sync"
	"time"
)

func downloading(id int, c chan string, wg *sync.WaitGroup) {

	defer wg.Done()
  
	time.Sleep(time.Second)
	c <- fmt.Sprintf("file %d downloading is complete", id)
}

 
func main() {

	var wg sync.WaitGroup
   c := make(chan string, 5)

  for i := 0; i <= 5; i++ {
		wg.Add(1)
		go downloading(i, c, &wg)
	}

	go func ()  {
		wg.Wait()
		close(c)
	}()

	for msg := range c {
		fmt.Println(msg)
	}
}