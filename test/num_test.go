package test

import (
	"fmt"
	"testing"
)

func TestNum(t *testing.T) {
	c := make(chan int, 10)
	z := make(chan int, 1)
	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()
	go func() {
		for {
			select {
			case _, b := <-c:
				if !b {
					z <- 1
				} else {
					fmt.Println(len(c))
				}
			}
		}
	}()
	<-z
}
