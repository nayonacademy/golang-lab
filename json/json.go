package main

import "fmt"

type Dictionary map[string]interface{}

func main() {
	data := []Dictionary{
		{
			"metrics": []Dictionary{
				{"tag_name": "output_current", "id": 3},
				{"tag_name": "input_voltage", "id": 2},
			},
			"port":       161,
			"timeout":    1,
			"sleep_time": 5,
		},
	}

	fmt.Println(data)
}