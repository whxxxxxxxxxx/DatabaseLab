package appInitialize

import "DatabaseLab/internal/app/common"

func init() {
	apps = append(apps, &common.Common{Name: "Common module"})
}
