package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Payload struct {
	Stuff Data
}

type Data struct {
	Fruit   Fruits
	Veggies Vegetables
}

type Vegetables map[string]int
type Fruits map[string]int

func main() {
	url := "http://localhost:1337"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	var p Payload

	err = json.Unmarshal(body, &p)

	if err != nil {
		panic(err)
	}

	fmt.Println(p.Stuff.Fruit, p.Stuff.Veggies)
}
