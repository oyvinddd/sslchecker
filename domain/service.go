package domain

type (
	Service interface {

    		Register(Domain) error
	}

	mockedService struct {

		domains []Domain
	}
)

func NewMockedService() Service {
	return &mockedService{
		domains: make([]Domain, 0)
	}
}

func (mockedService)Register(domain Domain) error {
	return nil
}

