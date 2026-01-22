package search

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8"
)

type ProductDocument struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Brand       string  `json:"brand"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
	Rating      float64 `json:"rating"`
	InStock     bool    `json:"in_stock"`
	CreatedAt   string  `json:"created_at"`
}

func IndexProduct(ctx context.Context, es *elasticsearch.Client, p ProductDocument) error {
	body, _ := json.Marshal(p)

	_, err := es.Index(
		"products",
		bytes.NewReader(body),
		es.Index.WithDocumentID(p.ID),
		es.Index.WithContext(ctx),
	)
	return err
}
