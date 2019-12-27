package main

import "fmt"

func main() {
	colors := map[string]string{
		"Red":    "red",
		"Yellow": "yellow",
	}

	for _, value := range colors {
		fmt.Println("1:", value)
	}

	removeColor(colors)

	for _, value := range colors {
		fmt.Println("2:", value)
	}
}

func removeColor(colors map[string]string) {
	delete(colors, "Red")
}
