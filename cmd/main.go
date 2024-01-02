package main

import (
	"context"

	"github.com/maadiii/helli/config"
	"github.com/maadiii/helli/internal/application/user/account"
	"github.com/maadiii/helli/internal/application/user/user"
	accountdata "github.com/maadiii/helli/internal/data_access/account"
	userdata "github.com/maadiii/helli/internal/data_access/user"
	accountdomain "github.com/maadiii/helli/internal/domain/account"
	userdomain "github.com/maadiii/helli/internal/domain/user"
	"github.com/maadiii/helli/pkg/event"
	"github.com/maadiii/hertzwrapper/server"
)

func main() {
	prepare()

	server.Run(
		config.Application().DevMode,
		server.WithAddress(config.Application().Address),
	)
}

func prepare() {
	ctx := context.Background()
	udata, adata := dataAccess()
	event := enentHandler(ctx)

	initUser(udata)
	initAccount(adata, udata, event)
}

func initUser(data userdata.DataAccess) {
	domain := userdomain.New(data)
	user.Init(domain)
}

func initAccount(data accountdata.DataAccess, udata userdata.DataAccess, event event.Raiser) {
	domain := accountdomain.New(data, udata, event)
	account.Init(domain)
}

func dataAccess() (userdata.DataAccess, accountdata.DataAccess) {
	return userdata.New(), accountdata.New()
}

func enentHandler(ctx context.Context) event.Raiser {
	return event.NewRaiser(
		ctx,
		config.Application().EventWorkerCount,
	)
}
