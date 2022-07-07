package delivery

import (
	"enigmacamp.com/lopei_grpc_cnlt/config"
	"enigmacamp.com/lopei_grpc_cnlt/manager"
	"fmt"
)

type cli struct {
	useCaseManager manager.UseCaseManager
}

func (c *cli) Run() {
	balance, err := c.useCaseManager.CheckBalanceUseCase().GetBalance(int32(1))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(balance)
}
func Cli() *cli {
	c := config.NewConfig()
	infraManager := manager.NewInfraManager(c)
	repoManager := manager.NewRepositoryManager(infraManager)
	useCaseManager := manager.NewUseCaseManager(repoManager)
	return &cli{
		useCaseManager: useCaseManager,
	}
}
