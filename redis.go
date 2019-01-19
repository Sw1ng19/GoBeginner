package main

import (
    "fmt"
    "github.com/garyburd/redigo/redis"
)

func newPool() *redis.Pool {
return &redis.Pool{
        MaxIdle: 80,
        MaxActive: 120,
        Wait: true,
        Dial: func() (redis.Conn, error) {
            c, err := redis.Dial("tcp", ":6379")
            if err != nil {
                panic(err.Error())
            }
            return c, err
        },
    }
}


var pool = newPool()

func main() {
    c := pool.Get()
    defer c.Close()

    test, _ := c.Do("set", "hello", "world!")
    fmt.Println(test)
}

func test() {
    c := pool.Get()
    defer c.Close()

    test, _ := c.Do("get", "hello")
    fmt.Println(test)
}

