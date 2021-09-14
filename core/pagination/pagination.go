package pagination

import (
	"errors"
	"github.com/bandungrhapsody/rhaprouter"
	"strconv"
	"strings"
)

type Pagination struct {
	ctx                *rhaprouter.Context
	fields             []string
	paginationRequired bool
}

type Payload map[string]string

/*
	Error Types
*/
var ErrInvalidPaginationFormat = errors.New("invalid Pagination format")

const (
	/*
		Query Params
	*/
	pageQueryParamKey = "_page"
	limitQueryParamKey = "_limit"
	disablePaginationQueryParamKey = "_disable_pagination"

	/*
		Keys
	*/
	pageKey = "page"
	limitKey = "limit"
	startIndexKey = "start_index"
	endIndexKey = "end_index"
)

func New(ctx *rhaprouter.Context, isRequired bool) *Pagination {
	return &Pagination{ctx: ctx, paginationRequired: isRequired}
}

func (pg *Pagination) Fields(fields ...string) *Pagination {
	pg.fields = fields
	return pg
}

func (pg *Pagination) Payload() (results Payload, err error) {
	results = make(map[string]string)
	page := pg.ctx.Query(pageQueryParamKey)
	limit := pg.ctx.Query(limitQueryParamKey)

	if !pg.paginationRequired {
		disablePagination := pg.ctx.Query(disablePaginationQueryParamKey)

		if disablePagination == "true" {
			pg.setFields(results)
			return
		}
	}

	pageNum, limitNum, isValid := validatePageAndLimit(page, limit)
	if !isValid {
		err = ErrInvalidPaginationFormat
		return
	}

	startIndex := (pageNum - 1) * limitNum
	endIndex := pageNum * limitNum

	results[pageKey] = page
	results[limitKey] = limit
	results[startIndexKey] = strconv.Itoa(startIndex)
	results[endIndexKey] = strconv.Itoa(endIndex)

	pg.setFields(results)
	return
}

func validatePageAndLimit(page string, limit string) (resultPage int, resultLimit int, isValid bool) {
	resultPage, err := strconv.Atoi(page)
	if err != nil {
		isValid = false
		return
	}

	resultLimit, err = strconv.Atoi(limit)
	if err != nil {
		isValid = false
		return
	}

	isValid = (resultPage > 0) && (resultLimit > 0)
	return
}

func (pg *Pagination) setFields(payload map[string]string) {
	for _, field := range pg.fields {
		payload[field] = strings.Trim(pg.ctx.Query(field), " ")
	}
}
