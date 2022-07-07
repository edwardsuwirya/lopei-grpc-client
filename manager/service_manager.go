package manager

type ServiceManager interface {
	CustomerCheckBalance(lopeiId int32) (float32, error)
	CustomerDoPayment(lopeiId int32, amount float32) error
}

type serviceManager struct {
	repoManager RepositoryManager
}

func (s *serviceManager) CustomerCheckBalance(lopeiId int32) (float32, error) {
	return s.repoManager.CustomerRepo().CheckBalance(lopeiId)
}

func (s *serviceManager) CustomerDoPayment(lopeiId int32, amount float32) error {
	return s.repoManager.CustomerRepo().DoPayment(lopeiId, amount)
}

func NewServiceManager(repoManager RepositoryManager) ServiceManager {
	return &serviceManager{
		repoManager: repoManager,
	}
}
