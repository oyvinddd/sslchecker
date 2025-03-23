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

func (s mockedService)Register(domainURL string) error {

	newDomain, err := New(domainURL)
	if err != nil {
		return err
	}

	s.domains = append(s.domains, *newDomain)

	return nil
}


