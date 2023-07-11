package types

type ProductRequest struct {
	Products []ProductInformation
}

type ProductInformation struct {
	Id      int
	Name    string
	Slug    string
	InStock bool `json:"in_stock"`
	Price   PriceInformation
    ImageList []ImageItem `json:"image_list"`
}

type PriceInformation struct {
	Display  string
	Numeric  float64
	Currency string
}

type ImageItem struct {
	Url string
}

type Cache struct {
	StockedItems map[string]bool
}

// Checks if an item is in stock
func (c *Cache) GetItemStatus(n string) bool {
	val, ok := c.StockedItems[n]

	return ok && val
}

// Changes an item stock status, if the item isn't in the cache then it is made
func (c *Cache) ChangeItemStatus(n string, s bool) {
	c.StockedItems[n] = s
}

func (c *Cache) ItemExists(n string) bool {
	_, ok := c.StockedItems[n]
	return ok
}
