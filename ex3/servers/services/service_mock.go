package services

import (
	"go_ex3/models"
	"go_ex3/utils"
)

func NewBeefSrvMock(url string) beefSrv {
	return beefSrv{url}
}

func (obj beefSrv) GetGroupBeefMock(text string) (result models.BeefRes, err error) {
	result.Beef = utils.GroupTextWord(text)
	return result, nil
}
