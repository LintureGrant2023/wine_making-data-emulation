package api

import (
	"backend/api/control_api"
	"backend/api/dashboard_api"
	"backend/api/globalctl_api"
	"backend/api/localctl_api"
	"backend/api/settings_api"
)

type ApiGroup struct {
	SettingsApi  settings_api.SettingsApi
	DashboardApi dashboard_api.DashboardApi
	GlobalctlApi globalctl_api.GlobalctlApi
	LocalctlApi  localctl_api.LocalctlApi
	ControlApi   control_api.ControlApi
}

var ApiGroupApp = new(ApiGroup)
