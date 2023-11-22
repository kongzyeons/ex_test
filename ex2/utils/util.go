package utils

import (
	"fmt"
	"strconv"
)

func DecodeMinNum(input string) (result string) {
	type Data struct {
		code rune
		num  int
	}
	var groupData []Data
	var data Data
	var currentNum int

	// Create dummy
	switch input[0] {
	case 76: // L
		currentNum = 1
		data = Data{code: 76, num: currentNum}
	case 61: // =
		currentNum = 0
		data = Data{code: 61, num: currentNum}
	case 82: // R
		currentNum = 0
		data = Data{code: 82, num: currentNum}
	}
	groupData = append(groupData, data)

	for i := 0; i < len(input); i++ {
		if input[i] != 76 && input[i] != 61 && input[i] != 82 {
			fmt.Println("input invalid")
			return
		}
		switch input[i] {
		case 76: // L
			currentNum = currentNum - 1
			data = Data{code: 76, num: currentNum}
		case 61: // =
			data = Data{code: 61, num: currentNum}
		case 82: // R
			currentNum = currentNum + 1
			data = Data{code: 82, num: currentNum}
		}

		groupData = append(groupData, data)
		if currentNum < 0 {
			currentNum = 0 - currentNum
			checkIdx := i + 1
			groupData[checkIdx].num = groupData[checkIdx].num + currentNum
			for {
				if checkIdx == 0 {
					break
				}
				if groupData[checkIdx].code == 76 && groupData[checkIdx].num >= groupData[checkIdx-1].num {
					groupData[checkIdx-1].num = groupData[checkIdx-1].num + currentNum
				} else if groupData[checkIdx].code == 61 && groupData[checkIdx].num != groupData[checkIdx-1].num {
					groupData[checkIdx-1].num = groupData[checkIdx-1].num + currentNum
				} else if groupData[checkIdx].code == 82 && groupData[checkIdx].num <= groupData[checkIdx-1].num {
					groupData[checkIdx-1].num = groupData[checkIdx-1].num + currentNum
				} else {
					break
				}
				checkIdx = checkIdx - 1
			}
			currentNum = 0
		}
	}
	for _, v := range groupData {
		result = result + strconv.Itoa(v.num)
	}
	return result
}
