package conf

import "gitlab.ghn.vn/common-projects/go-sdk/sdk"

var app *sdk.App

func SetApp(a *sdk.App) {
	app = a
}

func GetApp() *sdk.App {
	if app == nil {
		panic("app = nil")
	}

	return app
}
