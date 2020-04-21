package us

type RedirectRepository interface {
	Find(code string) (string, error)
	Store(code string, url string) error
}

type redirectRepository struct {
	repository RedirectRepository
}

func NewRedirectRepository(repository RedirectRepository) {
	// return &RedirectRepository

}
