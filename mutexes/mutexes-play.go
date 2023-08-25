package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

type Data struct {
	Items []string
}

var sharedData Data
var mu sync.Mutex

func fetchData(itemID int) {
	// Simulate fetching data
	item := fmt.Sprintf("Item %d", itemID)

	mu.Lock()
	sharedData.Items = append(sharedData.Items, item)
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup

	// Launch multiple goroutines to fetch data
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fetchData(id)
		}(i)
	}

	wg.Wait()

	// Convert to JSON and print
	result, _ := json.Marshal(sharedData)
	fmt.Println(string(result))
}
