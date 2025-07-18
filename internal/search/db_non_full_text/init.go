package db_non_full_text

import (
	"codeberg.org/alist/alist/v3/internal/search/searcher"
)

var config = searcher.Config{
	Name:       "database_non_full_text",
	AutoUpdate: true,
}

func init() {
	searcher.RegisterSearcher(config, func() (searcher.Searcher, error) {
		return &DB{}, nil
	})
}
