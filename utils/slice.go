package utils

import "fmt"

func PrintNItems[T any](items []T, max int) {
	fmt.Println("\nItems:")
	for _, itemId := range items[:5] {
		fmt.Println(itemId)
	}
	if len(items) > max {
		fmt.Println("... total items", len(items))
	}
}
