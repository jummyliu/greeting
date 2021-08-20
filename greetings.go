package greetings

import "fmt"

// Hello returns a greetings for the named person
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v, welcome!", name)
	return message
}
