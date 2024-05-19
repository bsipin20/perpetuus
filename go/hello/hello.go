package main

import "fmt"

const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "


func Hello(name string) string {
	if name == "" {
      return "Hello, World!"
  }
  if language == "spanish" {
    return spanishHelloPrefix + name

  }
	return "Hello, " + name + "!"
}

func main() {
	fmt.Println(Hello("World"))
}
