package fanin

import (
	"log"
	"testing"
)

func Test_FanInBlocking (t *testing.T) {

	chan1 := make(chan interface{})
	chan2 := make(chan interface{})

	output := func(a int, c chan interface{}) {
		for i := 0; i < 2; i++ {
			c <- struct {
				a int
				b string
			}{
				a: a,
				b: "hi",
			}
		}
		close(c)
	}

	go output(1, chan1)
	go output(2, chan2)

	resChan := fanin(0, chan1, chan2)
	for c := range resChan {
		log.Printf("%v", c)
	}

}

func Test_FanIn_NonBlocking (t *testing.T) {

	chan1 := make(chan interface{}, 3)
	chan2 := make(chan interface{}, 3)

	output := func(a int, c chan interface{}) {
		for i := 0; i < 2; i++ {
			c <- struct {
				a int
				b string
			}{
				a: a,
				b: "hi",
			}
		}
		close(c)
	}

	go output(1, chan1)
	go output(2, chan2)

	resChan := fanin(6, chan1, chan2)
	for c := range resChan {
		log.Printf("%v", c)
	}

}
