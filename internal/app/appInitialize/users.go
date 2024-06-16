package appInitialize

import "DatabaseLab/internal/app/users"

func init() {
	apps = append(apps, &users.Users{Name: "Users module"})
}
