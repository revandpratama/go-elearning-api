package api

import "github.com/revandpratama/go-elearning-api/service"

type krsAPI struct {
	service service.KRSService
}

func NewKRSAPI(s service.KRSService) *krsAPI{
	return &krsAPI{
		service: s,
	}
}

