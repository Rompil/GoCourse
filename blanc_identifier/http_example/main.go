package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//always check errors in a production code
	//res, _ := http.Get("http://www.ya.com")
	res, err := http.Get("http://www.ya.com")
	if err != nil {
		log.Fatal(err)
	}
	//page, _ := ioutil.ReadAll(res.Body)
	page, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", page)
}
