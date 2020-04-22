package shortener

type RedirectRepository interface {
	Find(code string) (string, error)
	Store(code string, url string) error
}

type redirectRepository struct {
}

func (r *redirectRepository) Find(code string) (string, error) {
	panic("implement me")
}

func (r *redirectRepository) Store(code string, url string) error {
	panic("implement me")
}

func NewRedirectRepository() RedirectRepository {
	return &redirectRepository{}
}
