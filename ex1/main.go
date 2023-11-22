package main

import (
	"encoding/json"
	"fmt"
	"go_ex1/utils"
	"log"
	"os"
)

func main() {

	input_1 := [][]int{
		{59},
		{73, 41},
		{52, 40, 53},
		{26, 53, 6, 34},
	}
	fmt.Println(utils.MaxDistanceArrayTree(input_1))

	fileContent, _ := os.ReadFile("hard.json")

	var input_2 [][]int
	err := json.Unmarshal(fileContent, &input_2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(utils.MaxDistanceArrayTree(input_2))

}
