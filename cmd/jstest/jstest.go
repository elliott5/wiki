package main

import (
	"fmt"

	"github.com/elliott5/wiki"
)

func main() {
	req, err := wiki.NewRequest("https://%s.wikipedia.org/w/api.php", "Cubism", "simple")
	if err != nil {
		panic(err.Error())
	}
	resp, err := req.Execute(true)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Result: %#v\n", *resp)
}
