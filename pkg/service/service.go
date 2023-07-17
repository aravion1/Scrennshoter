package service

import (
	"github.com/aravion1/Scrennshoter/structs"
)

type ImageGenerator interface {
	GetImage(p structs.Params) ([]byte, error)
	GetElementImageByUrl(p structs.Params) ([]byte, error)
}

type Service struct {
	ImageGenerator
}

func NewService() *Service {
	return &Service{ImageGenerator: NewImageGenerator()}
}
