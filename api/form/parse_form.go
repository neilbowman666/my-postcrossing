package form

import (
	"github.com/gin-gonic/gin"
	"my-postcrossing/api/resp"
	"net/http"
)

func FillJsonForm[FORM interface{}](c *gin.Context) *FORM {
	var form FORM
	err := c.BindJSON(&form)

	if err != nil {
		resp.Error(c, http.StatusBadRequest, "JSON format not allowed!", err)
		return nil
	}

	return &form
}
