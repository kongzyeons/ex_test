package services_test

import (
	"go_ex3/models"
	"go_ex3/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

var Url = "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"

func TestGetGroupBeefErr(t *testing.T) {
	tests := []struct {
		nameTest string
		beefSrv  services.BeefSrv
		expect   error
	}{
		{nameTest: "test success", beefSrv: services.NewBeefSrv(), expect: nil},
		{nameTest: "test err", beefSrv: services.NewBeefSrvMock("test err"), expect: models.ErrHttpRequest},
	}
	for i := range tests {
		t.Run(tests[i].nameTest, func(t *testing.T) {
			beefSrv := tests[i].beefSrv
			_, err := beefSrv.GetGroupBeef()
			assert.ErrorIs(t, err, tests[i].expect)
		})
	}
}

func TestGetGroupBeefResult(t *testing.T) {
	tests := []struct {
		nameTest string
		text     string
		expect   models.BeefRes
	}{
		{nameTest: "test result 1",
			text: "Fatback t-bone t-bone, pastrami  ..   t-bone.  pork, meatloaf jowl enim.  Bresaola t-bone.",
			expect: models.BeefRes{
				Beef: map[string]int{
					"t-bone":   4,
					"fatback":  1,
					"pastrami": 1,
					"pork":     1,
					"meatloaf": 1,
					"jowl":     1,
					"enim":     1,
					"bresaola": 1,
				},
			},
		},
	}
	for i := range tests {
		t.Run(tests[i].nameTest, func(t *testing.T) {
			beefSrv := services.NewBeefSrvMock("")
			result, _ := beefSrv.GetGroupBeefMock(tests[i].text)
			assert.Equal(t, tests[i].expect, result)
		})
	}

}
