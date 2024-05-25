package helper

import "github.com/revandpratama/go-elearning-api/dto"

type ResponseWithoutData struct {
	Status  string
	Message string
}

type ResponseWithData struct {
	Status  string
	Message string
	Data    any
}

func Response(param dto.ResponseParams) any {
	var response any
	var status string

	if param.StatusCode >= 200 && param.StatusCode < 400 {
		status = "success"
	} else {
		status = "failed"
	}

	if param.Data != nil {
		response = &ResponseWithData{
			Status:  status,
			Message: param.Message,
			Data:    param.Data,
		}
	} else {
		response = &ResponseWithoutData{
			Status:  status,
			Message: param.Message,
		}
	}

	return response
}
