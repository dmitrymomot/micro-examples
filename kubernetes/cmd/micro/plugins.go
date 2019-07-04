package main

import (
	"github.com/micro/go-plugins/micro/cors"
	rpc "github.com/micro/go-plugins/micro/disable_rpc"
	"github.com/micro/go-plugins/micro/index"
	"github.com/micro/go-plugins/micro/stats_auth"
	"github.com/micro/micro/api"
	"github.com/micro/micro/plugin"
	"github.com/micro/micro/web"
)

func init() {
	api.Register(rpc.NewPlugin())
	api.Register(cors.NewPlugin())
	api.Register(index.NewPlugin())
	web.Register(index.NewPlugin())
	plugin.Register(stats_auth.NewPlugin())
}
