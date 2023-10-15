package main
import "fmt"
import "time"
import "sync"
var wg sync.WaitGroup

func async(num int){
	defer wg.Done()
	fmt.Println("async code first", num)
	time.Sleep(time.Second)
	fmt.Println("async code second", num)
}

func main(){
	wg.Add(1)
	wg.Add(1)
	go async(1)
	go async(2)
	fmt.Println("sync code")
	wg.Wait()
}

