package main

import (
	"fmt"
	"sync"
)

/*
1.启动一个协程将1-2000的数放入一个channel中，比如numChan
2.启动8个协程从channel中取出数，并计算1+2+...+n的值,存放到resChan
3.8个协程工作完后，在遍历resChan
*/

var num int
var wg sync.WaitGroup

//生成1-2000个数
func putNum(numChan chan int) {
	for i := 1; i <= 1000; i++ {
		numChan <- i
	}
	close(numChan)
}

//计算
func calNum(numChan chan int, resChan chan int) {
	defer wg.Done()
	for {
		v, ok := <-numChan
		if !ok {
			break
		}
		num += v
		resChan <- num
	}

}

func main() {
	numChan := make(chan int, 1000)
	resChan := make(chan int, 1000)
	num2 := 0

	go putNum(numChan)

	for i := 0; i < 8; i++ {
		go calNum(numChan, resChan)
		wg.Add(1)
	}
	wg.Wait()

	close(resChan)

	for {
		res, ok := <-resChan
		if !ok {
			break
		}
		num2++
		fmt.Printf("[%v]值为%v\n", num2, res)
	}

	//for {
	//	select {
	//	case res:= <-resChan :
	//		num2++
	//		fmt.Printf("[%v]值为%v\n", num2, res)
	//	default:
	//		fmt.Println("没了，读取不到了")
	//		return
	//	}
	//}

}
