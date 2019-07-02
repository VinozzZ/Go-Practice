package	main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//Write a multi-threaded C program that gives readers priority over writers concerning a shared (global) variable. Essentially, if any readers are waiting, then they have priority over writer threads -- writers can only write when there are no readers. This program should adhere to the following constraints:
//
//    Multiple readers/writers must be supported (5 of each is fine)
//    Readers must read the shared variable X number of times
//    Writers must write the shared variable X number of times
//    Readers must print:
//        The value read
//        The number of readers present when value is read
//    Writers must print:
//        The written value
//        The number of readers present were when value is written (should be 0)
//    Before a reader/writer attempts to access the shared variable it should wait some random amount of time
//        Note: This will help ensure that reads and writes do not occur all at once
//    Use pthreads, mutexes, and condition variables to synchronize access to the shared variable

func main() {
	var mu sync.RWMutex
	var counter int
	var rList int

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go reader(&counter, &rList, &mu, &wg)
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go writer(&counter, &rList, &mu, &wg)
	}

	wg.Wait()
}

func reader(counter *int, rList *int, m *sync.RWMutex, wg *sync.WaitGroup) {
	m.RLock()
	sleep()
	*rList++
	fmt.Printf("Reader: %v, Value: %v\n", *rList, *counter)
	m.RUnlock()
	*rList--
	wg.Done()
}

func writer(counter *int, rList *int, m *sync.RWMutex, wg *sync.WaitGroup) {
	m.Lock()
	sleep()
	for *rList > 0 {
		sleep()
	}
	*counter++
	m.Unlock()
	fmt.Printf("Writer: %v, Value: %v\n", *rList, *counter)
	wg.Done()
}

func sleep() {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
}