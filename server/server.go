package server

import (
	config "blog-service/config"
	blogProto "blog-service/rpc/blog"
	"context"
)

type Server struct{}

func (*Server) CreateBlog(ctx context.Context, req *blogProto.CreateBlogRequest) (*blogProto.CreateBlogResponse, error) {
	data := &blogProto.CreateBlogRequest{
		Title:   req.GetTitle(),
		Content: req.GetContent(),
	}

	res, err := config.DB.CreateBlog(data)
	return res, err
}

func (*Server) GetBlog(ctx context.Context, req *blogProto.GetBlogRequest) (*blogProto.GetBlogResponse, error) {
	data := &blogProto.GetBlogRequest{
		Id: req.GetId(),
	}

	res, err := config.DB.GetBlog(data)
	return res, err
}

func (*Server) UpdateBlog(ctx context.Context, req *blogProto.UpdateBlogRequest) (*blogProto.UpdateBlogResponse, error) {
	return &blogProto.UpdateBlogResponse{}, nil
}

func (*Server) DeleteBlog(ctx context.Context, req *blogProto.DeleteBlogRequest) (*blogProto.DeleteBlogResponse, error) {
	return &blogProto.DeleteBlogResponse{}, nil
}

func (*Server) ListBlog(ctx context.Context, req *blogProto.ListBlogRequest) (*blogProto.ListBlogResponse, error) {
	return &blogProto.ListBlogResponse{}, nil
}
