package utils

import (
	"errors"
)

type PaginationRes struct {
	Total       int  `json:"total"`
	TotalPage   *int `json:"total_page"`
	CurrentPage int  `json:"current_page"`
	NextPage    *int `json:"next_page"`
}

func GetPagination(numberItems int, limit int, currPage int) (totalPage *int, nextPage *int, err error) {
	if numberItems <= 0 || limit <= 0 || currPage <= 0 {
		return nil, nil, errors.New("params must be positive integers")
	}

	totalPage = new(int)
	nextPage = new(int)

	*totalPage = numberItems / limit
	if numberItems%limit != 0 {
		*totalPage++
	}

	if currPage < *totalPage {
		*nextPage = currPage + 1
	} else {
		nextPage = nil
	}

	return totalPage, nextPage, nil
}
