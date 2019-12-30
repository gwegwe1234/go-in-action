package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	c := make(map[string]interface{})
	c["name"] = "Gopher"
	c["title"] = "programmer"
	c["contact"] = map[string]interface{}{
		"home": "123456",
		"cell": "56789000",
	}

	//맵을 JSON으로 마샬링
	data, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		log.Println("Error", err)
		return
	}

	fmt.Println(string(data))

}
