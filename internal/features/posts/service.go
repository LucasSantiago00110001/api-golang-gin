package posts

import "context"

type Service interface {
	List(ctx context.Context, page int) (ListResponse, error)
}

type PostService struct {
	repository Repository
}

func NewService(repository Repository) *PostService {
	return &PostService{repository: repository}
}

func (s *PostService) List(ctx context.Context, page int) (ListResponse, error) {
	registros, total, err := s.repository.List(ctx, page)
	if err != nil {
		return ListResponse{}, err
	}
	if registros == nil {
		registros = []Post{}
	}

	return ListResponse{Total: total, Page: page, Registros: registros}, nil
}
