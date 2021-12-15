package server

import (
	blogProto "blog-service/rpc/blog"
	"context"
)

type Server struct{}

func (*Server) CreateBlog(ctx context.Context, req *blogProto.CreateBlogRequest) (*blogProto.CreateBlogResponse, error) {
	return &blogProto.CreateBlogResponse{}, nil
}

func (*Server) GetBlog(ctx context.Context, req *blogProto.GetBlogRequest) (*blogProto.GetBlogResponse, error) {
	return &blogProto.GetBlogResponse{}, nil
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
