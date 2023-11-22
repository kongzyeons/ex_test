package controller

import (
	"go_ex3/models"
	"go_ex3/services"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type BeefRest interface {
	GetGroupBeef(c *fiber.Ctx) error
}

type beefRest struct {
	beefSrv services.BeefSrv
}

func NewBeefRest(beefSrv services.BeefSrv) BeefRest {
	return beefRest{beefSrv}
}

func (obj beefRest) GetGroupBeef(c *fiber.Ctx) error {
	result, err := obj.beefSrv.GetGroupBeef()
	if err != nil {
		log.Println("error", err)
		return c.Status(http.StatusInternalServerError).JSON(models.Err_response(err))
	}
	log.Println("success", c.Status(http.StatusOK))
	return c.Status(http.StatusOK).JSON(result)
}
