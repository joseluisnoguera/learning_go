package greetings

import "fmt"

func HelloToName(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
