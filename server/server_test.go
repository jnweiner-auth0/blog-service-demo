package server_test

import (
	blogProto "blog-service/rpc/blog"
	"blog-service/server"
	"blog-service/server/mock"
	"context"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestServer(t *testing.T) {

	// best way to set up a before each?
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDBClient := mock.NewMockDBClient(mockCtrl)

	testServer := server.NewServer(mockDBClient)

	t.Run("CreateBlog() calls DBClient.CreateBlog() with expected data", func(t *testing.T) {
		testBlog := &blogProto.CreateBlogRequest{
			Title:   "Test title",
			Content: "Test content",
		}
		testResult := &blogProto.CreateBlogResponse{
			Id:      "1",
			Title:   "Test title",
			Content: "Test content",
		}
		mockDBClient.EXPECT().CreateBlog(testBlog).Return(testResult, nil).Times(1)

		testServer.CreateBlog(context.Background(), testBlog)
	})

	t.Run("DeleteBlog() calls DBClient.DeleteBlog() with expected data", func(t *testing.T) {
		testBlog := &blogProto.DeleteBlogRequest{
			Id: "1",
		}
		testResult := &blogProto.DeleteBlogResponse{
			Id: "1",
		}
		mockDBClient.EXPECT().DeleteBlog(testBlog).Return(testResult, nil).Times(1)
		testServer.DeleteBlog(context.TODO(), testBlog)
	})

	t.Run("UpdateBlog() calls DBClient.UpdateBlog() with expected data", func(t *testing.T) {
		testBlog := &blogProto.UpdateBlogRequest{
			Id:      "1",
			Title:   "Test title",
			Content: "Test content",
		}
		testResult := &blogProto.UpdateBlogResponse{
			Id:      "1",
			Title:   "Test title",
			Content: "Test content",
		}
		mockDBClient.EXPECT().UpdateBlog(testBlog).Return(testResult, nil).Times(1)
		testServer.UpdateBlog(context.Background(), testBlog)
	})

	t.Run("GetBlog() calls DBClient.GetBlog() with expected data", func(t *testing.T) {
		testBlog := &blogProto.GetBlogRequest{
			Id: "1",
		}
		testResult := &blogProto.GetBlogResponse{
			Id:      "1",
			Title:   "Test title",
			Content: "Test content",
		}
		mockDBClient.EXPECT().GetBlog(testBlog).Return(testResult, nil).Times(1)
		testServer.GetBlog(context.Background(), testBlog)
	})

	t.Run("ListBlog() calls DBClient.ListBlog() with given Limit if Limit provided", func(t *testing.T) {
		testBlogWithLimit := &blogProto.ListBlogRequest{
			Limit: 10,
		}
		testResult := &blogProto.ListBlogResponse{
			Blogs: []*blogProto.CreateBlogResponse{},
		}
		mockDBClient.EXPECT().ListBlog(testBlogWithLimit).Return(testResult, nil).Times(1)
		testServer.ListBlog(context.Background(), testBlogWithLimit)
	})

	t.Run("ListBlog() calls DBClient.ListBlog() with Limit: 25 if no Limit provided", func(t *testing.T) {
		testBlogWithoutLimit := &blogProto.ListBlogRequest{}
		testResult := &blogProto.ListBlogResponse{
			Blogs: []*blogProto.CreateBlogResponse{},
		}
		mockDBClient.EXPECT().ListBlog(&blogProto.ListBlogRequest{Limit: 25}).Return(testResult, nil).Times(1)
		testServer.ListBlog(context.Background(), testBlogWithoutLimit)
	})

}
