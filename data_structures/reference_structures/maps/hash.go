package main

import (
	"fmt"
	"log"
	"net/http"
	//	"io/ioutil"
	"bufio"
	"os"
)

func main() {
	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEn.txt")
	if err != nil {
		log.Fatalln(err)
	}

	words := make(map[string]string)
	sc := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		words[sc.Text()] = ""
	}

	if err := sc.Err(); err != nil {
		fmt.Fprint(os.Stderr, "reading error", err)
	}

	i := 0
	for f, _ := range words {
		fmt.Println(f)
		if i > 200 {
			break
		}
		i++
	}
	//bs, _ := ioutil.ReadAll(res.Body)
	//str := string(bs)
	//	fmt.Println(str)
}
