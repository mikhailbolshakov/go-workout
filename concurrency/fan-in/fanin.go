package fanin

import "sync"

func fanin(cap int, chans ...<-chan interface{}) <-chan interface{} {

	resChan := make(chan interface{}, cap)
	var wg sync.WaitGroup
	wg.Add(len(chans))

	f := func(ch <-chan interface{}) {
		defer wg.Done()
		for c := range ch {
			resChan <- c
		}
	}

	for _, ch := range chans {
		go f(ch)
	}

	go func() {
		wg.Wait()
		close(resChan)
	}()

	return resChan
}
