package main

import (
	"fmt"
	"net/http"
	"sync"
	"net/url"
)


func main () {
	var wg sync.WaitGroup
	wg.Add(1)
	testUrl := "http://localhost:8080/api/lib/1/app-launched"
	resp, err := http.PostForm(testUrl, url.Values{"key": {"Value"}, "id": {"123"}})
	if err != nil {
		fmt.Println("Error : ", err)
	}
	fmt.Println(resp)
	wg.Done()
	wg.Wait()
}
