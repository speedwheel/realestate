package routes

import (
	"github.com/kataras/iris/hero"
	"github.com/speedwheel/immigrationrealestate/web/bootstrap"
)

func Configure(b *bootstrap.Bootstrapper) {
	b.WWW.Get("/", hero.Handler(GetHomeHandler))
}
