package main

import (
	"fmt"
)

func main() {
	greetMap := make(map[string]string)
	greetMap["Ronnie"] = "hello"
	greetMap["Jenny"] = "hey"

	newGreeting := make(map[string]string)
	newGreeting["Joe"] = "yo"
	newGreeting["Peter"] = "sup"

	fmt.Println(greetMap)
	fmt.Println(newGreeting)

	for k, v := range newGreeting {
		greetMap[k] = v
	}

	fmt.Println(greetMap)
}
