package greetings

import "fmt"

//Hello returns a greeting for the named person.
func Hello (name string,name2 string) string {

	message := fmt.Sprintf("Hi, %v %v . Welcome!", name,name2)
	return message
} 