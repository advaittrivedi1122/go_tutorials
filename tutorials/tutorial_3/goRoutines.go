package main
import "fmt"
import "time"
import "sync"
var wg sync.WaitGroup

func async(){
	defer wg.Done()
	fmt.Println("async code 1")
	time.Sleep(time.Second)
	fmt.Println("async code 2")
}

func main(){
	wg.Add(1)
	go async()
	go async()
	fmt.Println("sync code")
	wg.Wait()
}

