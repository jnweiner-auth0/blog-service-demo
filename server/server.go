//go:generate mockgen -destination mock/server_mock.go -package mock blog-service/server DBClient

package server

import (
	blogProto "blog-service/rpc/blog"
	"context"
)

type DBClient interface {
	Connect() error
	CreateBlog(*blogProto.CreateBlogRequest) (*blogProto.CreateBlogResponse, error)
	GetBlog(*blogProto.GetBlogRequest) (*blogProto.GetBlogResponse, error)
	UpdateBlog(*blogProto.UpdateBlogRequest) (*blogProto.UpdateBlogResponse, error)
	DeleteBlog(*blogProto.DeleteBlogRequest) (*blogProto.DeleteBlogResponse, error)
	ListBlog(*blogProto.ListBlogRequest) (*blogProto.ListBlogResponse, error)
}

type Server struct {
	db DBClient
}

func NewServer(db DBClient) *Server {
	return &Server{
		db: db,
	}
}

// use the BlogService interface generated in service.twirp.go as guideline to stub out the expected functions

func (s *Server) CreateBlog(ctx context.Context, req *blogProto.CreateBlogRequest) (*blogProto.CreateBlogResponse, error) {
	data := &blogProto.CreateBlogRequest{
		Title:   req.GetTitle(),
		Content: req.GetContent(),
	}

	res, err := s.db.CreateBlog(data)
	return res, err
}

func (s *Server) GetBlog(ctx context.Context, req *blogProto.GetBlogRequest) (*blogProto.GetBlogResponse, error) {
	data := &blogProto.GetBlogRequest{
		Id: req.GetId(),
	}

	res, err := s.db.GetBlog(data)
	return res, err
}

func (s *Server) UpdateBlog(ctx context.Context, req *blogProto.UpdateBlogRequest) (*blogProto.UpdateBlogResponse, error) {
	data := &blogProto.UpdateBlogRequest{
		Id:      req.GetId(),
		Title:   req.GetTitle(),
		Content: req.GetContent(),
	}

	res, err := s.db.UpdateBlog(data)
	return res, err
}

func (s *Server) DeleteBlog(ctx context.Context, req *blogProto.DeleteBlogRequest) (*blogProto.DeleteBlogResponse, error) {
	data := &blogProto.DeleteBlogRequest{
		Id: req.GetId(),
	}

	res, err := s.db.DeleteBlog(data)
	return res, err
}

func (s *Server) ListBlog(ctx context.Context, req *blogProto.ListBlogRequest) (*blogProto.ListBlogResponse, error) {
	limit := int64(25)

	if req.GetLimit() > 0 {
		limit = req.GetLimit()
	}

	data := &blogProto.ListBlogRequest{
		Limit: limit,
	}

	res, err := s.db.ListBlog(data)
	return res, err
}
