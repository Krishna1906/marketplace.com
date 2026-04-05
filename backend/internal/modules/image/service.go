package image

type Service struct {
	Repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{Repo: repo}
}

func (s *Service) GetImagesByProductID(productID int) ([]Image, error) {
	return s.Repo.GetByProductID(productID)
}