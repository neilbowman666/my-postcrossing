package api

import (
	"github.com/gin-gonic/gin"
	"my-postcrossing/api/form"
	"my-postcrossing/api/resp"
	"my-postcrossing/db"
	"my-postcrossing/m"
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
