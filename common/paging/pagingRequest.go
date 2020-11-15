package paging

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type (
	PagingRequest struct {
		Page  int      `json:"page"`
		Limit int      `json:"limit"`
		Sort  []string `json:"sort"`
	}
)

func NewPagingRequest(ctx echo.Context) (PagingRequest, error) {
	return PagingRequest{
		Page:  getIntParamWithDefault(ctx, "page", 1),
		Limit: getIntParamWithDefault(ctx, "limit", 100),
		Sort:  getSort(ctx, []string{"-created_at"}),
	}, nil
}

func (pr PagingRequest) SortQuery() string {
	var sort = ""
	for _, s := range pr.Sort {
		if len(s) < 1 {
			continue
		}

		if len(sort) != 0 {
			sort += ","
		}

		if "-" != string(s[0]) {
			sort += s
		} else {
			sort += fmt.Sprintf("%s %s", s[1:], "desc")
		}
	}
	return sort

}

func getIntParamWithDefault(ctx echo.Context, key string, defaultValue int) int {
	dataQp := ctx.QueryParam(key)
	if dataQp == "" {
		return defaultValue
	}

	dataInt, err := strconv.Atoi(dataQp)
	if err != nil {
		return defaultValue
	}

	return dataInt
}

func getSort(ctx echo.Context, defaultSort []string) []string {
	dataQp := ctx.QueryParam("sort")
	sorts := strings.Split(dataQp, ",")

	sortsLen := len(sorts)

	if sortsLen == 0 {
		return defaultSort
	}

	for i := 0; i < sortsLen; i++ {
		sorts[i] = strings.TrimSpace(sorts[i])
	}

	return sorts
}
