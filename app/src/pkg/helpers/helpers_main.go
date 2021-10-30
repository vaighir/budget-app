package helpers

import (
	"github.com/vaighir/go-diet/app/pkg/config"
)

var app *config.AppConfig

func InitializeHelpers(a *config.AppConfig) {
	app = a
}
