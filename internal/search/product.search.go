package search

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
)

func SearchProducts(
	ctx context.Context,
	es *elasticsearch.Client,
	query map[string]interface{},
) ([]map[string]interface{}, error) {

	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		return nil, err
	}

	res, err := es.Search(
		es.Search.WithContext(ctx),
		es.Search.WithIndex("products"),
		es.Search.WithBody(&buf),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.IsError() {
		return nil, fmt.Errorf("elasticsearch error: %s", res.String())
	}

	var response map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}

	// 🔒 SAFE EXTRACTION (NO PANIC)
	hitsObj, ok := response["hits"].(map[string]interface{})
	if !ok {
		return []map[string]interface{}{}, nil
	}

	rawHits, ok := hitsObj["hits"].([]interface{})
	if !ok {
		return []map[string]interface{}{}, nil
	}

	results := make([]map[string]interface{}, 0)

	for _, h := range rawHits {
		hit, ok := h.(map[string]interface{})
		if !ok {
			continue
		}

		source, ok := hit["_source"].(map[string]interface{})
		if !ok {
			continue
		}

		results = append(results, source)
	}

	return results, nil
}
