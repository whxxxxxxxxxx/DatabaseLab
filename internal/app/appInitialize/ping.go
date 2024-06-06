package appInitialize

import "IOTProject/internal/app/ping"

func init() {
	apps = append(apps, &ping.Ping{Name: "ping module"})
}
