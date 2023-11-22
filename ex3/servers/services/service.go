package services

import (
	"go_ex3/models"
	"go_ex3/utils"
)

type BeefSrv interface {
	GetGroupBeef() (result models.BeefRes, err error)
}

type beefSrv struct {
	url string
}

func NewBeefSrv() BeefSrv {
	url := "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"
	return beefSrv{url}
}

func (obj beefSrv) GetGroupBeef() (result models.BeefRes, err error) {
	text, err := utils.CallApiGetBody(obj.url)
	if err != nil {
		return result, err
	}
	result.Beef = utils.GroupTextWord(text)
	return result, nil
}
