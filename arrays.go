package main
import "fmt"

func main() {
    //定义数组
    var a[5] int
    fmt.Println("emp:", a)
    
    //赋值
    a[4] = 100
    fmt.Println("set:", a)
}
