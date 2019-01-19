package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.baidu.com")
	if err != nil {
		fmt.Println("Request error!")
	} 

	defer res.Body.Close()
	
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))
}
