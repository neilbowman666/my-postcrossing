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
// @Tags mail
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
// @Tags mail
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

// DeleteSentMails
// @Summary Delete Sent Mails
// @Description Delete sent mails
// @Tags mail
// @Param id path int true "sent mail ID"
// @Success 200 {object} resp.Pack[resp.None]
// @Router /api/v1/sent-mails/{id} [DELETE]
func DeleteSentMails(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.Error(c, 400, "path param endpoint_id of /api/v1/sent-mails/{id} should be integer", err)
		return
	}

	tx := db.Transaction[m.SentMail]{Tx: db.DBConn}
	err = tx.Delete(uint(id))
	if err != nil {
		resp.Error(c, 500, "system error, contact administrator", err)
		return
	}
	resp.SuccessMessage(c, "OK")
	return
}
