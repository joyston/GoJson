package main

type Person struct {
	FirstName string
	LastName  string
	Age       int
	email     string
}

//using struct field tags
type sPerson struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	email     string `json:"email"`
}

func main() {
	// ExecuteMarshaling()
	// ExecuteMarshalingWIthTags()
	// ExecuteUnmarshaling()
	ExecuteUnmarshalingWIthTags()
}
