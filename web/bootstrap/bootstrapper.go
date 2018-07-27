package bootstrap

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/core/router"
	"github.com/kataras/iris/hero"
	"github.com/kataras/iris/middleware/recover"
	"github.com/speedwheel/immigrationrealestate/datasource/mongo"
)

type Configurator func(*Bootstrapper)

type Bootstrapper struct {
	*iris.Application
	WWW router.Party
}

func New(cfgs ...Configurator) *Bootstrapper {
	b := &Bootstrapper{
		Application: iris.New(),
	}
	b.WWW = b.Subdomain("www")

	b.Logger().SetLevel("debug")
	b.Use(recover.New())

	db := mongo.New()
	hero.Register(db)

	return b
}

func (b *Bootstrapper) Configure(cfgs ...Configurator) {
	for _, cfg := range cfgs {
		cfg(b)
	}
}

func (b *Bootstrapper) Bootstrap() *Bootstrapper {
	return b
}
