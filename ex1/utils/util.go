package utils

func MaxDistanceArrayTree(input [][]int) int {

	l := len(input)
	for i := 0; i < l-1; i++ {
		array_1 := input[len(input)-1]
		array_2 := input[len(input)-2]

		for j := range array_2 {
			if array_1[j] > array_1[j+1] {
				array_2[j] = array_2[j] + array_1[j]
			} else {
				array_2[j] = array_2[j] + array_1[j+1]
			}
		}
		input = input[:len(input)-1]
		input[len(input)-1] = array_2
	}
	return input[0][0]
}
