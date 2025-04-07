package main

import "fmt"

func main() {
	anonStruct := struct {
		Bools map[string]bool
		Slice []string
	}{
		Bools: map[string]bool{
			"key":        true,
			"anotherKey": false,
		},
		Slice: []string{
			"key", "anotherKey", "oneMoreKey",
		},
	}

	fmt.Println(anonStruct)
}
