package delivery

import (
	"enigmacamp.com/lopei_grpc_cnlt/config"
	"enigmacamp.com/lopei_grpc_cnlt/manager"
)

type cli struct {
	serviceManager manager.ServiceManager
}

func (c *cli) Run() {
	c.serviceManager.CustomerCheckBalance(int32(1))
}
func Cli() *cli {
	c := config.NewConfig()
	infraManager := manager.NewInfraManager(c)
	repoManager := manager.NewRepositoryManager(infraManager)
	serviceManager := manager.NewServiceManager(repoManager)
	return &cli{
		serviceManager: serviceManager,
	}
}
