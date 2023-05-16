package db

import (
	"gorm.io/gorm"
	"math"
	"my-postcrossing/api/resp"
	"my-postcrossing/m"
)

type Transaction[T m.CouldBeVisible] struct {
	Tx *gorm.DB
}

func (tx Transaction[T]) Save(value *T) error {
	return DBConn.Save(value).Error
}

func (tx Transaction[T]) Page(page int, size int, desc bool) resp.PagingPack[T] {
	// limit
	limit := 10
	if size != 0 {
		limit = size
	}
	// offset
	offset := (page - 1) * limit
	// sort
	sort := "id asc"
	if desc {
		sort = "id desc"
	}

	// 查询列表
	var dataList = make([]T, 0)
	tx.Tx.Scopes(
		func(db *gorm.DB) *gorm.DB {
			return db.Limit(limit).Offset(offset).Order(sort)
		},
	).Find(&dataList)
	var totalRows int64 = 0

	tx.Tx.Count(&totalRows)
	totalPages := int(math.Ceil(float64(totalRows) / float64(limit)))

	return resp.PagingPack[T]{
		Page:       page,
		PageSize:   limit,
		TotalPages: totalPages,
		TotalRows:  totalRows,
		List:       dataList,
	}
}

func (tx Transaction[T]) Delete(id uint) error {
	var mm T
	return DBConn.Delete(&mm, id).Error
}
