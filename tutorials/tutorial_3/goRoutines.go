package main
import "fmt"
import "sync"
var wg sync.WaitGroup

func async(num int){
	defer wg.Done()
	fmt.Println("async code first", num)
	fmt.Println("async code second", num)
}

func main(){

	wg.Add(1)
	go async(11)
	wg.Wait() // will wait for 11
	
	wg.Add(1)
	go async(22)
	
	fmt.Println("sync code")
	
	wg.Add(1)
	go async(33)
	wg.Wait() // will wait for all pending added wg till here (22,33)
}

