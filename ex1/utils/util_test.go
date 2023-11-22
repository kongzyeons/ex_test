package utils_test

import (
	"encoding/json"
	"go_ex1/utils"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxDistanceArrayTree(t *testing.T) {
	fileContent, _ := os.ReadFile("../hard.json")

	var input_test [][]int
	err := json.Unmarshal(fileContent, &input_test)
	if err != nil {
		log.Fatal(err)
	}

	tests := []struct {
		nameTest  string
		inputTest [][]int
		expect    int
	}{
		{nameTest: "test 1",
			inputTest: [][]int{
				{59},
				{73, 41},
				{52, 40, 53},
				{26, 53, 6, 34},
			},
			expect: 237,
		},
		{nameTest: "test 2",
			inputTest: [][]int{{1}, {2, 3}, {4, 5, 6}, {7, 8, 9, 10}},
			expect:    20,
		},
		{nameTest: "test 3",
			inputTest: [][]int{
				{1},
				{3, 5},
				{2, 6, 4},
				{8, 7, 2, 1},
				{9, 5, 3, 8, 7},
				{10, 8, 9, 15, 17, 5},
			},
			expect: 39,
		},
		{nameTest: "test 4",
			inputTest: input_test,
			expect:    7273,
		},
	}

	for i := range tests {
		t.Run(tests[i].nameTest, func(t *testing.T) {
			result := utils.MaxDistanceArrayTree(tests[i].inputTest)
			assert.Equal(t, tests[i].expect, result)
		})
	}

}
