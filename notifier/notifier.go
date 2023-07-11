package notifier

import (
	"fmt"
	"time"

	"github.com/philip727/gcc-scraper/gcc"
	"github.com/philip727/gcc-scraper/types"
)

func ConstantlyCheckStore(s *types.Settings, c *types.Cache) {
	for {
		var products types.ProductRequest
		// Grabs the products from the store
		err := gcc.RequestProducts(&products)
		if err != nil {
			fmt.Printf("Err when making request to gcc: %s", err)
		}

		// Checks if we have the products selected and if we do, notify that they are in stock
		updateCacheAndNotify(products.Products, s, c)
		time.Sleep(5 * time.Second)
	}
}

func updateCacheAndNotify(pi []types.ProductInformation, s *types.Settings, c *types.Cache) {
	for _, product := range pi {
		if c.GetItemStatus(product.Slug) != product.InStock && s.ItemSelected(product.Name) {
            c.ChangeItemStatus(product.Slug, product.InStock)
            
            createStockNotification(product, product.InStock)
            continue
		}

        // We only need to proceed if the item doesn't exist yet in the cache
        if c.ItemExists(product.Slug) {
            continue
        }

		// We set to false so if the user updates his settings, it will notify about item in stock
        c.ChangeItemStatus(product.Slug, false)
	}
}
