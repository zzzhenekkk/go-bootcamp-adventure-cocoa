package main

import (
	"errors"
	"fmt"
)

func getElement(arr []int, idx int) (int, error) {
	if len(arr) < idx+1 || idx < 0 {
		return 0, errors.New("index out of range")
	}

	current := arr[0]
	for i := 0; i < idx; i++ {
		arr = arr[1:]
		current = arr[0]
	}
	return current, nil
}

func main() {
	arr := []int{10, 20, 30, 40, 50}
	element, err := getElement(arr, 5)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Element:", element) // Output: Element: 40
	}
}
