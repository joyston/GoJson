package main

import (
	"encoding/json"
	"fmt"
)

func ExecuteUnmarshaling() {
	b := `{"FirstName":"Joyston","LastName":"Fernandes","Age":27, "email":"dumy@g.com"}`
	var P Person
	err := json.Unmarshal([]byte(b), &P)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", P)
}

func ExecuteUnmarshalingWIthTags() {
	b := `{"first_name":"Joyston","last_name":"Fernandes","age":27, "place_of_birth": "London", "gender":"male", "email":"dumy@g.com"}`
	var P sPerson
	err := json.Unmarshal([]byte(b), &P)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", P)
}
