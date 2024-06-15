package appInitialize

import "DatabaseLab/internal/app/ping"

func init() {
	apps = append(apps, &ping.Ping{Name: "ping module"})
}
