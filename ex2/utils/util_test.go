package utils_test

import (
	"go_ex/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeMinNum(t *testing.T) {
	tests := []struct {
		nameTest  string
		inputTest string
		expect    string
	}{
		{nameTest: "test L=R", inputTest: "L=R", expect: "1001"},
		{nameTest: "test R=L", inputTest: "R=L", expect: "0110"},
		{nameTest: "test RRR", inputTest: "RRR", expect: "0123"},
		{nameTest: "test ===", inputTest: "===", expect: "0000"},
		{nameTest: "test LLL", inputTest: "LLL", expect: "3210"},
		{nameTest: "test LLRR", inputTest: "LLRR", expect: "21012"},
		{nameTest: "test RRLL", inputTest: "RRLL", expect: "01210"},
		{nameTest: "test LLRR=", inputTest: "LLRR=", expect: "210122"},
		{nameTest: "test ==RLL", inputTest: "==RLL", expect: "000210"},
		{nameTest: "test =LLRR", inputTest: "=LLRR", expect: "221012"},
		{nameTest: "test RLRLL=L", inputTest: "RLRLL=L", expect: "01032110"},
		{nameTest: "test RRRLLLL=L", inputTest: "RRRLLLL=L", expect: "0125432110"},
	}
	for i := range tests {
		t.Run(tests[i].nameTest, func(t *testing.T) {
			result := utils.DecodeMinNum(tests[i].inputTest)
			assert.Equal(t, tests[i].expect, result)
		})
	}

}
