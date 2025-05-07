package item

import "github.com/naphat-sirisubkulchai/shop/modules/models"
type(
	CreateItemReq struct {
		Title    string  `json:"title" validate:"required,max=64"`
		Price    float64 `json:"price" validate:"required"`
		Defect   string   `json:"defect" validate:"required"`
		ImageUrl string  `json:"image_url" validate:"required,max=255"`
	}

	ItemShowCase struct {
		ItemId   string  `json:"item_id"`
		Title    string  `json:"title"`
		Price    float64 `json:"price"`
		Defect   string  `json:"defect"`
		ImageUrl string  `json:"image_url"`
	}

	ItemSearchReq struct {
		Title string `query:"title" validate:"max=64"`
		models.PaginateReq
	}

	ItemUpdateReq struct {
		Title    string  `json:"title" validate:"required,max=64"`
		Price    float64 `json:"price" validate:"required"`
		ImageUrl string  `json:"image_url" validate:"required,max=255"`
		Defect   string  `json:"defect" validate:"required"`
	}

	EnableOrDisableItemReq struct {
		UsageStatus bool `json:"usage_status"`
	}
)