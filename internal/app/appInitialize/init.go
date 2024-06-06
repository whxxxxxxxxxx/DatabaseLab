package appInitialize

import "IOTProject/internal/app"

var (
	apps = make([]app.Module, 0)
)

func GetApps() []app.Module {
	return apps
}
