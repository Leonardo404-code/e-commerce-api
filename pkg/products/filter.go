package products

import (
	"fmt"
	"strconv"
	"strings"
)

type (
	Filter struct {
		ID         []string
		Pagination *Pagination
	}

	Pagination struct {
		Page       int
		MaxPerPage int
	}
)

func ParseFilters(queries map[string]string) (*Filter, error) {
	filter := &Filter{
		ID: []string{},
		Pagination: &Pagination{
			Page:       1,
			MaxPerPage: 1,
		},
	}

	if id, ok := queries["id"]; ok {
		filter.ID = strings.Split(id, ",")
	}

	if page, ok := queries["page"]; !ok {
		pageConverted, err := strconv.Atoi(page)
		if err != nil {
			return nil, err
		}

		if pageConverted == 0 {
			return nil, fmt.Errorf("invalid page")
		}

		filter.Pagination.Page = pageConverted
	}

	if maxPerPage, ok := queries["maxPerPage"]; ok {
		maxPerPageConverted, err := strconv.Atoi(maxPerPage)
		if err != nil {
			return nil, err
		}

		if maxPerPageConverted == 0 {
			return nil, fmt.Errorf("invalid max per page")
		}

		filter.Pagination.MaxPerPage = maxPerPageConverted
	}

	return filter, nil
}
