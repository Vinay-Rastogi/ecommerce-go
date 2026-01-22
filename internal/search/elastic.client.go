package search

import (
	"github.com/elastic/go-elasticsearch/v8"
)

func NewElasticClient() (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}
	return elasticsearch.NewClient(cfg)
}
