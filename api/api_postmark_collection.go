package api

import (
	"github.com/gin-gonic/gin"
	"my-postcrossing/api/form"
	"my-postcrossing/api/resp"
	"my-postcrossing/db"
	"my-postcrossing/m"
	"strconv"
)

// AddPostmarkCollection
// @Summary Add Postmark Collection
// @Description Add postmark collection
// @Tags postmark collection
// @Param HTTP_POST_body body form.AddPostmarkCollectionForm true "Add postmark collection form"
// @Success 201 {object} resp.Pack[resp.None]
// @Router /api/v1/postmark-collections [POST]
func AddPostmarkCollection(c *gin.Context) {
	f := form.FillJsonForm[form.AddPostmarkCollectionForm](c)
	postmarkCollection := f.ToM()
	err := db.Transaction[m.PostmarkCollection]{Tx: db.DBConn}.Save(postmarkCollection)
	if err != nil {
		resp.Error(c, 500, "db error", err)
	} else {
		resp.SuccessMessageDataCode(c, "OK", gin.H{}, 201)
		return
	}
}

// ListPostmarkCollections
// @Summary List Postmark Collections
// @Description List postmark collections
// @Tags postmark collection
// @Param page query int true "page number, start from 1."
// @Param page_size query int true "page size"
// @Success 200 {object} resp.Pack[resp.PagingPack[m.PostmarkCollectionVo]]
// @Router /api/v1/postmark-collections [GET]
func ListPostmarkCollections(c *gin.Context) {
	page, err := strconv.Atoi(c.Request.URL.Query().Get("page"))
	pageSize, err := strconv.Atoi(c.Request.URL.Query().Get("page_size"))
	if err != nil {
		resp.Error(c, 400, "param \"page\" or \"page_size\" is not valid, please use an integer number", err)
		return
	}

	tx := db.Transaction[m.PostmarkCollection]{Tx: db.DBConn}
	pageResult := tx.Page(page, pageSize, true)

	resp.SuccessMessageDataCode(c, "OK", resp.ParsePager(pageResult), 200)
	return
}

// DeletePostmarkCollection
// @Summary Delete Postmark Collection
// @Description Delete postmark collection
// @Tags postmark collection
// @Param id path int true "postmark collection ID"
// @Success 200 {object} resp.Pack[resp.None]
// @Router /api/v1/postmark-collections/{id} [DELETE]
func DeletePostmarkCollection(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		resp.Error(c, 400, "path param `id` should be integer", err)
		return
	}

	tx := db.Transaction[m.PostmarkCollection]{Tx: db.DBConn}
	err = tx.Delete(uint(id))
	if err != nil {
		resp.Error(c, 500, "system error, contact administrator", err)
		return
	}
	resp.SuccessMessage(c, "OK")
	return
}

// DeletePostmarkCollectionsInBatch
// @Summary Delete PostmarkCollections In Batch
// @Description Delete postmark collections in batch
// @Tags postmark collection
// @Param HTTP_POST_body body []uint true "batch deletion ID list"
// @Success 200 {object} resp.Pack[resp.None]
// @Router /api/v1/postmark-collections/any/batch-deletions [POST]
func DeletePostmarkCollectionsInBatch(c *gin.Context) {
	f := form.FillJsonForm[[]uint](c)
	tx := db.Transaction[m.PostmarkCollection]{Tx: db.DBConn}
	err := tx.Delete(*f...)
	if err != nil {
		resp.Error(c, 500, "system error, contact administrator", err)
		return
	}
	resp.SuccessMessage(c, "OK")
	return
}
