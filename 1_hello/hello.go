package main

import "fmt"

// func main() {
// 	fmt.Println("Hello, World")
// }

// const greetPrefix = "Hello, "
const spanishGreetPrefix = "Hola, "
const englishGreetPrefix = "Hello, "
const frenchGreetPrefix = "Test, "

// Hello ...
func Hello(name string, lang string) string {
	if name == "" {
		name = "world"
	}
	prefix := englishGreetPrefix

	switch lang {
	case "Spanish":
		prefix = spanishGreetPrefix
	case "French":
		prefix = frenchGreetPrefix
	}
	return prefix + name
	// if lang == "Spanish" {
	// 	return spanishGreetPrefix + name
	// }
	// return englishGreetPrefix + name

}

func main() {
	fmt.Println(Hello("Hiro", "spanish"))
}
