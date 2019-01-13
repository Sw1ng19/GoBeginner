package main
import "fmt"

func main() {
    nums := []int{2,3,4}

    //traversal key & value
    for k,v := range nums {
        fmt.Println("key", k)
        fmt.Println("value", v)
    }
}

