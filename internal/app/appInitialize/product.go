package appInitialize

import "DatabaseLab/internal/app/product"

func init() {
	apps = append(apps, &product.Product{Name: "Product module"})
}
