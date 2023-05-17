package api

import (
	"github.com/gin-gonic/gin"
	"my-postcrossing/api/form"
	"my-postcrossing/api/resp"
	"my-postcrossing/db"
	"my-postcrossing/m"
	"strconv"
)

// AddContact
// @Summary Add contact
// @Description send mail
// @Tags contact
// @Param HTTP_POST_body body form.AddContactForm true "add contact form"
// @Success 201 {object} resp.Pack[resp.None]
// @Router /api/v1/contacts [POST]
func AddContact(c *gin.Context) {
	f := form.FillJsonForm[form.AddContactForm](c)
	contact := f.ToM()
	err := db.Transaction[m.Contact]{Tx: db.DBConn}.Save(contact)
	if err != nil {
		resp.Error(c, 500, "db error", err)
	} else {
		resp.SuccessMessageDataCode(c, "OK", gin.H{}, 201)
		return
	}
}

// ListContacts
// @Summary List Contacts
// @Description List contacts
// @Tags contact
// @Param page query int true "page number, start from 1."
// @Param page_size query int true "page size"
// @Success 200 {object} resp.Pack[resp.PagingPack[m.ContactVo]]
// @Router /api/v1/contacts [GET]
func ListContacts(c *gin.Context) {
	page, err := strconv.Atoi(c.Request.URL.Query().Get("page"))
	pageSize, err := strconv.Atoi(c.Request.URL.Query().Get("page_size"))
	if err != nil {
		resp.Error(c, 400, "param \"page\" or \"page_size\" is not valid, please use an integer number", err)
		return
	}

	tx := db.Transaction[m.Contact]{Tx: db.DBConn}
	pageResult := tx.Page(page, pageSize, true)

	resp.SuccessMessageDataCode(c, "OK", resp.ParsePager(pageResult), 200)
	return
}
