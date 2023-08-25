package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
)

const baseURL = "https://www.trade-tariff.service.gov.uk/api/v2/commodities/"
const concurrencyLimit = 1500

type Response struct {
	Data struct {
		ID         string `json:"id"`
		Attributes struct {
			FormattedDescription    string `json:"formatted_description"`
			GoodsNomenclatureItemID string `json:"goods_nomenclature_item_id"`
		} `json:"attributes"`
	} `json:"data"`
}

type SafeDescriptions struct {
	mu           sync.Mutex
	Descriptions map[string]map[string]string
}

func (sd *SafeDescriptions) AddDescription(id string, desc map[string]string) {
	sd.mu.Lock()
	defer sd.mu.Unlock()
	sd.Descriptions[id] = desc
}

func fetchDescription(id string, sd *SafeDescriptions) {
	resp, err := http.Get(baseURL + id)
	if err != nil {
		fmt.Println("Error fetching ID", id, ":", err)
		return
	}
	defer resp.Body.Close()

	var r Response
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		fmt.Println("Error decoding response for ID", id, ":", err)
		return
	}

	description := map[string]string{
		"formatted_description":      r.Data.Attributes.FormattedDescription,
		"goods_nomenclature_item_id": r.Data.Attributes.GoodsNomenclatureItemID,
		"goods_nomenclature_sid":     r.Data.ID,
	}

	sd.AddDescription(id, description)
}

func loadIDsFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var ids []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ids = append(ids, scanner.Text())
	}

	return ids, scanner.Err()
}

func main() {
	ids, err := loadIDsFromFile("goods_nomenclature_item_ids.txt")
	if err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	sd := &SafeDescriptions{
		Descriptions: make(map[string]map[string]string),
	}
	semaphore := make(chan struct{}, concurrencyLimit)

	for _, id := range ids {
		semaphore <- struct{}{}
		wg.Add(1)
		go func(id string) {
			defer wg.Done()
			fetchDescription(id, sd)
			<-semaphore
		}(id)
	}

	wg.Wait()

	// Convert the map to pretty-printed JSON
	prettyJSON, err := json.MarshalIndent(sd.Descriptions, "", "  ")
	if err != nil {
		fmt.Println("Failed to generate JSON:", err)
		return
	}
	fmt.Println(string(prettyJSON))
}
