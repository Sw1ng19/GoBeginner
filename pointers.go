package main
import "fmt"

func zval(ival int) {
    ival = 0
}

func zptr(iptr *int) {
    *iptr = 3 
}

func main() {
    i := 1
    fmt.Println(i)

    zval(i)
    fmt.Println(i)
    
    zptr(&i)
    fmt.Println(i)
}
