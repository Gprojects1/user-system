package middleware

import (
	"net/http"
	"user-system/pkg/errors"

	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
	var status int
	errorType := errors.GetType(err)
	switch errorType {
	case errors.BadRequest:
		status = http.StatusBadRequest
	case errors.NotFound:
		status = http.StatusNotFound
	case errors.NilUserId:
		status = http.StatusBadRequest
	case errors.NilRole:
		status = http.StatusBadRequest
	case errors.NotSaved:
		status = http.StatusInternalServerError
	case errors.WrongType:
		status = http.StatusBadRequest
	case errors.EmptyData:
		status = http.StatusBadRequest
	default:
		status = http.StatusInternalServerError
	}
	c.Writer.WriteHeader(status)

	response := gin.H{"error": err.Error(), "message": errorType.Message()}

	errorContext := errors.GetErrorContext(err)
	if errorContext != nil {
		response["context"] = errorContext
	}
	c.JSON(status, response)
}
