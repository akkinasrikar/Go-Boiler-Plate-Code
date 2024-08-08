package helper

import (
	"errors"
	"net/url"

	"example/boiler-plate/constants"
	"example/boiler-plate/models"
)

var ErrorParamMissingOrInvalid = func(msg string, param string) *models.AppError {
	return &models.AppError{
		Code:    422,
		Message: errors.New(msg),
		Type:    "param_missing_or_invalid",
		Param:   param,
	}
}

var ErrorInternalSystemError = func(msg string) *models.AppError {
	return &models.AppError{
		Code:    int64(constants.ErrorType.INTERNAL_SYSTEM_ERROR.Code),
		Message: errors.New(msg),
		Type:    constants.ErrorType.INTERNAL_SYSTEM_ERROR.Name,
	}
}

var ErrorDownStreamError = func() *models.AppError {
	return &models.AppError{
		Code:    int64(constants.ErrorType.DOWNSTREAM_ERROR.Code),
		Message: errors.New("downstream error"),
		Type:    constants.ErrorType.DOWNSTREAM_ERROR.Name,
	}
}

func SetInternalError(errMsg string) models.AppErrorResponse {
	zwErrBody := models.AppErrorBody{
		Message: errMsg,
		Type:    constants.ErrorType.INTERNAL_SYSTEM_ERROR.Name,
		Code:    int64(constants.ErrorType.INTERNAL_SYSTEM_ERROR.Code),
	}
	return models.AppErrorResponse{
		ErrorType: zwErrBody,
	}
}

func GetValidationEcomError(e url.Values) models.AppError {
	var ecomErr models.AppError
	for key, value := range e {
		ecomErr = *ErrorParamMissingOrInvalid(value[0], key)
		break
	}
	return ecomErr
}

func ErrorUnknownParam(param string) *models.AppError {
	return &models.AppError{
		Code:    422,
		Message: errors.New("unknown param"),
		Type:    "unknown_param",
		Param:   param,
	}
}
