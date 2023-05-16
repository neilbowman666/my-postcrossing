package resp

import "my-postcrossing/m"

type PagingPack[T any] struct {
	Page       int   `json:"page" example:"10"`
	PageSize   int   `json:"page_size" example:"1"`
	TotalPages int   `json:"total_pages" example:"3"`
	TotalRows  int64 `json:"total_rows" example:"24"`
	List       []T   `json:"list"`
}

func ParsePager[FROM m.CouldBeVisible](originalPager PagingPack[FROM]) PagingPack[m.Visible] {
	toData := make([]m.Visible, 0)

	for _, one := range originalPager.List {
		toData = append(toData, one.ToVo())
	}

	return PagingPack[m.Visible]{
		Page:       originalPager.Page,
		PageSize:   originalPager.PageSize,
		TotalPages: originalPager.TotalPages,
		TotalRows:  originalPager.TotalRows,
		List:       toData,
	}
}
