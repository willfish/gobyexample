package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	var marshalled []byte

	marshalled, _ = json.Marshal(true)
	fmt.Println(string(marshalled))

	marshalled, _ = json.Marshal(1)
	fmt.Println(string(marshalled))

	marshalled, _ = json.Marshal(2.34)
	fmt.Println(string(marshalled))

	marshalled, _ = json.Marshal("gopher")
	fmt.Println(string(marshalled))

	marshalled, _ = json.Marshal([]string{"apple", "peach", "pear"})
	fmt.Println(string(marshalled))

	marshalled, _ = json.Marshal(map[string]int{"apple": 5, "lettuce": 7})
	fmt.Println(string(marshalled))

	fixture := &response1{Page: 1, Fruits: []string{"apple", "peach", "pear"}}
	marshalled, _ = json.Marshal(fixture)
	fmt.Println(string(marshalled))

	fixture2 := &response2{Page: 1, Fruits: []string{"apple", "peach", "pear"}}
	marshalled, _ = json.Marshal(fixture2)
	fmt.Println(string(marshalled))

	rawBytes := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var data map[string]any

	fmt.Println(string(rawBytes))

	if err := json.Unmarshal(rawBytes, &data); err != nil {
		panic(err)
	}
	fmt.Println(data)

	testNum := 64.5

	fmt.Printf("value type: %.2f, value is %s\n", testNum, reflect.TypeOf(testNum))

	num := data["num"].(float64)

	fmt.Println(num)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	enc.Encode(d)

}
