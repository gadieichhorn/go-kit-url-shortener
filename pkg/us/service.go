package us

type RedirectService interface {
	Find(code string) (string, error)
	Store(code string, url string) error
}

type redirectService struct {
	repository *RedirectRepository
}

func NewRedirectService(repository *RedirectRepository) RedirectService {
	return &redirectService{
		repository,
	}
}

func (service *redirectService) Find(code string) (string, error) {
	return "", nil
}

func (service *redirectService) Store(code string, url string) error {
	return nil
}
