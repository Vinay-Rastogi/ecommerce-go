package services

import (
	"context"
	"strconv"

	"ecommerce/internal/search"

	"github.com/elastic/go-elasticsearch/v8"
)

type SearchService struct {
	es *elasticsearch.Client
}

func NewSearchService(es *elasticsearch.Client) *SearchService {
	return &SearchService{es: es}
}

type SearchParams struct {
	Query     string
	Brand     string
	Category  string
	InStock   string
	MinPrice  string
	MaxPrice  string
	SortBy    string // price | rating
	SortOrder string // asc | desc
	Page      string
	Limit     string
}

func (s *SearchService) SearchProducts(
	ctx context.Context,
	params SearchParams,
) ([]map[string]interface{}, error) {

	// Pagination defaults
	page := 1
	limit := 10

	if params.Page != "" {
		page, _ = strconv.Atoi(params.Page)
	}
	if params.Limit != "" {
		limit, _ = strconv.Atoi(params.Limit)
	}

	from := (page - 1) * limit

	// Filters
	filters := []map[string]interface{}{}

	if params.Brand != "" {
		filters = append(filters, map[string]interface{}{
			"term": map[string]interface{}{"brand": params.Brand},
		})
	}

	if params.Category != "" {
		filters = append(filters, map[string]interface{}{
			"term": map[string]interface{}{"category": params.Category},
		})
	}

	if params.InStock != "" {
		val, _ := strconv.ParseBool(params.InStock)
		filters = append(filters, map[string]interface{}{
			"term": map[string]interface{}{"in_stock": val},
		})
	}

	if params.MinPrice != "" || params.MaxPrice != "" {
		rangeQuery := map[string]interface{}{}
		if params.MinPrice != "" {
			rangeQuery["gte"], _ = strconv.ParseFloat(params.MinPrice, 64)
		}
		if params.MaxPrice != "" {
			rangeQuery["lte"], _ = strconv.ParseFloat(params.MaxPrice, 64)
		}
		filters = append(filters, map[string]interface{}{
			"range": map[string]interface{}{
				"price": rangeQuery,
			},
		})
	}

	// Sorting
	sort := []map[string]interface{}{}
	if params.SortBy == "price" || params.SortBy == "rating" {
		order := "asc"
		if params.SortOrder == "desc" {
			order = "desc"
		}
		sort = append(sort, map[string]interface{}{
			params.SortBy: map[string]interface{}{
				"order": order,
			},
		})
	}

	// Final ES Query
	query := map[string]interface{}{
		"from": from,
		"size": limit,
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{
						"multi_match": map[string]interface{}{
							"query":  params.Query,
							"fields": []string{"name", "description"},
						},
					},
				},
				"filter": filters,
			},
		},
	}

	if len(sort) > 0 {
		query["sort"] = sort
	}

	return search.SearchProducts(ctx, s.es, query)
}
