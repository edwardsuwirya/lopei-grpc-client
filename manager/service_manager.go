package manager

import (
	"enigmacamp.com/lopei_grpc_cnlt/usecase"
)

type ServiceManager interface {
	CheckBalanceService() usecase.CheckBalanceUseCase
}

type serviceManager struct {
	repoManager RepositoryManager
}

func (s *serviceManager) CheckBalanceService() usecase.CheckBalanceUseCase {
	return usecase.NewCheckBalanceService(s.repoManager.CustomerRepo())
}

func NewServiceManager(repoManager RepositoryManager) ServiceManager {
	return &serviceManager{
		repoManager: repoManager,
	}
}
