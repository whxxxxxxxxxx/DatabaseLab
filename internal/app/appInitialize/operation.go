package appInitialize

import "DatabaseLab/internal/app/operation"

func init() {
	apps = append(apps, &operation.Operation{Name: "Operation module"})
}
