package api

import (
	"github.com/gin-gonic/gin"
	"my-postcrossing/api/form"
	"my-postcrossing/api/resp"
	"my-postcrossing/db"
	"my-postcrossing/m"
	"strconv"
)

// SendMail
// @Summary Send mail
// @Description send mail
// @Tags sent-mail
// @Param HTTP_POST_body body form.SendMailForm true "send mail form"
// @Success 201 {object} resp.Pack[resp.None]
// @Router /api/v1/sent-mails [POST]
func SendMail(c *gin.Context) {
	f := form.FillJsonForm[form.SendMailForm](c)
	mail := f.ToM()
	err := db.Transaction[m.SentMail]{Tx: db.DBConn}.Save(mail)
	if err != nil {
		resp.Error(c, 500, "db error", err)
	} else {
		resp.SuccessMessageDataCode(c, "OK", gin.H{}, 201)
		return
	}
}

// ListSentMails
// @Summary List Sent Mails
// @Description List sent mails
// @Tags sent-mail
// @Param page query int true "page number, start from 1."
// @Param page_size query int true "page size"
// @Success 200 {object} resp.Pack[resp.PagingPack[m.SentMailVo]]
// @Router /api/v1/sent-mails [GET]
func ListSentMails(c *gin.Context) {
	page, err := strconv.Atoi(c.Request.URL.Query().Get("page"))
	pageSize, err := strconv.Atoi(c.Request.URL.Query().Get("page_size"))
	if err != nil {
		resp.Error(c, 400, "param \"page\" or \"page_size\" is not valid, please use an integer number", err)
		return
	}

	tx := db.Transaction[m.SentMail]{Tx: db.DBConn}
	pageResult := tx.Page(page, pageSize, true)

	resp.SuccessMessageDataCode(c, "OK", resp.ParsePager(pageResult), 200)
	return
}
