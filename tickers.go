package main
import "time"
import "fmt"

func main() {
	//打点器和定时器非常相似
	ticker := time.NewTicker(time.Second * 1)
	
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Second * 5)
	fmt.Println("Ticker stop!!")
}
