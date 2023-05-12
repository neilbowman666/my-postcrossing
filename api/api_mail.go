package api

import (
	"github.com/gin-gonic/gin"
	"my-postcrossing/api/form"
	"my-postcrossing/api/resp"
	"my-postcrossing/db"
)

// SendMail
// @Summary Send mail
// @Description send mail
// @Tags sent-mail
// @Param HTTP_POST_body body form.SendMailForm true "send mail form"
// @Router /api/v1/sent-mails [POST]
func SendMail(c *gin.Context) {
	f := form.FillJsonForm[form.SendMailForm](c)
	m := f.ToM()
	err := db.DefaultTx.Save(m)
	if err != nil {
		resp.Error(c, 500, "db error", err)
	} else {
		resp.SuccessMessageDataCode(c, "OK", gin.H{}, 201)
		return
	}
}

func ListSentMails(c *gin.Context) {
}
