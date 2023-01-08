package types

import "github.com/fajarhadifirmansyah/bedbo/utils"

type PagingReqBasic struct {
	Page     *int
	PageSize *int
	Column   string
	Sort     string
	Search   string
}

func (reqPaging *PagingReqBasic) SetPagingParam(ent interface{}, defaultCol string, defaultSort string) (*int, *string, *int) {

	if *reqPaging.Page == 0 {
		*reqPaging.Page = 1
	}

	switch {
	case *reqPaging.PageSize > 100:
		*reqPaging.PageSize = 100
	case *reqPaging.PageSize <= 0:
		*reqPaging.PageSize = 10
	}

	if len(reqPaging.Column) <= 0 {
		reqPaging.Column = defaultCol
	}

	if len(reqPaging.Sort) <= 0 {
		reqPaging.Sort = defaultSort
	}

	column := utils.GetColumnName(ent, reqPaging.Column)
	order := column + " " + reqPaging.Sort
	reqPaging.Search = "%" + reqPaging.Search + "%"
	offset := (*reqPaging.Page - 1) * *reqPaging.PageSize

	return reqPaging.PageSize, &order, &offset

}
