package utils

import (
	"strings"

	"gorm.io/gorm"
)

type (
	PageFilter struct {
		Page     int  `json:"page"`
		Limit    int  `json:"limit"`
		AllPages bool `json:"allPages"`
	}

	SortFilter struct {
		SortField string `json:"sortField"`
		SortOrder string `json:"sortOrder"`
	}
)

func AddPagination(db *gorm.DB, pf PageFilter, sf SortFilter) *gorm.DB {
	// Handle sorting
	if sf.SortField != "" {
		order := "asc"
		if strings.ToLower(sf.SortOrder) == "desc" {
			order = "desc"
		}
		db = db.Order(sf.SortField + " " + order)
	}

	// Handle pagination
	if !pf.AllPages {
		if pf.Limit <= 0 {
			pf.Limit = 10
		}
		if pf.Page <= 0 {
			pf.Page = 1
		}

		offset := (pf.Page - 1) * pf.Limit
		db = db.Offset(offset).Limit(pf.Limit)
	}

	return db
}
