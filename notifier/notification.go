package notifier

import (
	"fmt"

	"github.com/gen2brain/beeep"
	"github.com/philip727/gcc-scraper/types"
)

// Creates a notification about it being in stock
func createStockNotification(p types.ProductInformation, s bool) {
    image := fmt.Sprintf("https:%s", p.ImageList[0].Url)

	if s {
		beeep.Notify(p.Name, fmt.Sprintf("Now in stock. ($%s)", p.Price.Display), image)
		return
	}

	beeep.Notify(p.Name, "Is no longer in stock", image)
}
