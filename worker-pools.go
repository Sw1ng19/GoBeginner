package main
import "fmt"
import "time"

func worker(id int, jobs<-chan int, results chan<-int) {
    for j := range jobs {
        fmt.Println("worker", id, "processing job", j)
        time.Sleep(time.Second)
        results <- j * 2
    }
}

func main() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)

    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
	}
	
    for i := 1; i <= 9; i++ {
        jobs <- i
    }
    close(jobs)

    for k := 1; k <= 9; k++ {
        <-results
    }
}