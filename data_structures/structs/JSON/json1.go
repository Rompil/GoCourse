package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	Firstname string `json:"First Name"` //this how tags work, check json output
	Lastname  string `json:"-"`
	Age       int
	isvalid   bool //this field isn't converted to JSON, this is not exported field.
}

func main() {
	p := Person{"Jim", "Johnson", 21, false}
	bs, _ := json.Marshal(p)
	fmt.Printf("%T - %v\n", bs, bs)
	ss := string(bs)
	fmt.Println(ss)

	var ubs Person
	bs = []byte(`{"First Name":"Jim","Age":21}`) // `` is used to take a raw string. we have "" in this string
	json.Unmarshal(bs, &ubs)
	fmt.Println(ubs)

	json.NewEncoder(os.Stdout).Encode(p)

	json.NewDecoder(strings.NewReader(ss)).Decode(&ubs)
	fmt.Printf("%T - %v\n", ubs, ubs)
}
