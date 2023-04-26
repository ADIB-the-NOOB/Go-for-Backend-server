package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string   `json:"coursename"`
	Age      int      `json:"age"`
	Email    string   `json:"email"`
	Password string   `json:"-"`
	Language []string `json:"tags,omitempty"`
}

func main() {
	users := []User{
		{"Adib", 12, "mdadib550@gmail.com", "hello", []string{"Python", "Java"}},
		{"Noob", 23, "the@gmail.com", "hello", []string{"Python", "Java"}},
		{"Booo", 23, "kamla@gmail.com", "hello", []string{}},
	}
	finalJson, _ := json.MarshalIndent(users, "", "\t")
	jsonString := string(finalJson)
	fmt.Println(jsonString)
}
