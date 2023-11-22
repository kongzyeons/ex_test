package main

import (
	"fmt"
	"go_ex/utils"
)

func main() {
	// input := "LLRR="
	input_1 := "LLRR="
	input_2 := "==RLL"
	input_3 := "=LLRR"

	fmt.Printf("input = %v output = %v\n", input_1, utils.DecodeMinNum(input_1))
	fmt.Printf("input = %v output = %v\n", input_2, utils.DecodeMinNum(input_2))
	fmt.Printf("input = %v output = %v\n", input_3, utils.DecodeMinNum(input_3))

}
