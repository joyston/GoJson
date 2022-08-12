package main

import (
	"encoding/json"
	"fmt"
)

func ExecuteMarshaling() {
	p := Person{
		FirstName: "Joyston",
		LastName:  "Fernandes",
		Age:       27,
		email:     "dummy@gmail.com",
	}

	json, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(json))
}

func ExecuteMarshalingWIthTags() {
	p := sPerson{
		FirstName: "Joyston",
		LastName:  "Fernandes",
		Age:       27,
		email:     "dummy@gmail.com",
	}

	json, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(json))
}
