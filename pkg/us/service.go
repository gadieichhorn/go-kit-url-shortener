package us

type RedirectService interface {
	Find(code string) (string, error)
	Store(url string) (string, error)
}

type redirectService struct {
	repository RedirectRepository
}

func NewRedirectService(repository RedirectRepository) RedirectService {
	return &redirectService{
		repository,
	}
}

func (service *redirectService) Find(code string) (url string, err error) {
	return "", nil
}

func (service *redirectService) Store(url string) (code string, err error) {
	return "", nil
}
