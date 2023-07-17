package service

import (
	"github.com/aravion1/Scrennshoter/structs"
	"github.com/chromedp/chromedp"
)

type ImageGenerator interface {
	GetImage(p structs.Params) ([]byte, error)
	GetFullPageImage(p structs.Params, res *[]byte) chromedp.Tasks
}

type Service struct {
	ImageGenerator
}

func NewService() *Service {
	return &Service{ImageGenerator: NewImageGenerator()}
}
