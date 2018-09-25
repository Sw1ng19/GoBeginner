package main
import "fmt"

func main() {
    two := make([][]int, 3)

    for i:=0; i<3; i++ {
        innerLen := i+1
        two[i] = make([]int, innerLen)
        for j:=0; j<innerLen; j++ {
            two[i][j] = i+j
        }
    }
    fmt.Println(two)
}
