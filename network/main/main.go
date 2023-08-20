package main

import (
	"context"
	"manager/network/adapter"
	"manager/shared/resource"
	"os"
	"os/signal"
)

const (
	svcName    = "manager-network"
	svcVersion = "0.0.1"
)

// func ConfigCheckAccess(sr *resource.ServerResource) port.ICheckAccessService {
// 	return service.NewCheckUserPermissionService(adapter_access.NewUserRepository(sr.Db), sr)
// }

func main() {

	_ = svcName
	_ = svcVersion

	//===============================
	//Signal Interruption: Configure
	//===============================
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-c
		cancel()
	}()
	//============================
	//Service: Configure and Start
	//============================
	sr := resource.NewServerResource("env.toml")

	//Global Middlewares
	// sr.SetServiceCheckAccess(ConfigCheckAccess)
	// sr.UseGlobalMiddleware(middleware.CheckAccess(sr))

	//Register Handlers
	sr.AddHandler(adapter.NewCompanyHandlerRest(sr))

	sr.Run(ctx)
}