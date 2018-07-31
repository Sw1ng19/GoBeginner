package main
import "fmt"

type person struct {
    name string
    age int
}

func main() {
    fmt.Println(person{name:"bob", age:20})
    fmt.Println(person{"bob", 20})
    fmt.Println(person{name:"fred"})
    s := person{name:"Ann", age:40}

    sp := &s
    fmt.Println(sp.age)
}
