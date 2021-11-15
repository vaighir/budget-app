package helpers

import (
	"github.com/vaighir/budget-app/app/pkg/config"
)

var app *config.AppConfig

func InitializeHelpers(a *config.AppConfig) {
	app = a
}
