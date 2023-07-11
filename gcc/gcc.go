package gcc

import (
	"encoding/json"
	"net/http"

	"github.com/philip727/gcc-scraper/types"
)

// Makes a request to the products
func RequestProducts(r *types.ProductRequest) error {
	resp, err := http.Get("https://ec588ea8-5a73-46be-993a-a443b2fd4fb8.mysimplestore.com/api/v2/products?per_page=999")
	if err != nil {
        return err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
        return err
	}

    return nil
}
