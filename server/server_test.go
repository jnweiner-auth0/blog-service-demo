package server_test

import (
	blogProto "blog-service/rpc/blog"
	"blog-service/server"
	"blog-service/server/mock"
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
)

/*
for reference:
gomock docs: https://github.com/golang/mock
gomock tutorial: https://gist.github.com/thiagozs/4276432d12c2e5b152ea15b3f8b0012e
*/

var mockCtrl *gomock.Controller
var mockDBClient *mock.MockDBClient
var testServer *server.Server

func beforeEach(t *testing.T) {
	mockCtrl = gomock.NewController(t)
	mockDBClient = mock.NewMockDBClient(mockCtrl)
	testServer = server.NewServer(mockDBClient)
}

func TestCreateBlog(t *testing.T) {
	beforeEach(t)
	// mockCtrl.Finish() will assert the expectation statements
	defer mockCtrl.Finish()

	testReq := &blogProto.CreateBlogRequest{
		Title:   "Test title",
		Content: "Test content",
	}

	t.Run("calls DB with given data and returns expected response", func(t *testing.T) {
		testResp := &blogProto.CreateBlogResponse{
			Id:      "1",
			Title:   "Test title",
			Content: "Test content",
		}

		mockDBClient.EXPECT().CreateBlog(testReq).Return(testResp, nil).Times(1)

		resp, err := testServer.CreateBlog(context.Background(), testReq)
		if err != nil {
			t.Fatalf("unexpected error in DB call: %q", err)
		}
		if want, got := testResp, resp; want != got {
			t.Fatalf("wanted response: %q, got response: %q", want, got)
		}
	})

	t.Run("returns error if DB call fails", func(t *testing.T) {
		dbError := errors.New("Mock DB error")

		mockDBClient.EXPECT().CreateBlog(testReq).Return(nil, dbError).Times(1)

		resp, err := testServer.CreateBlog(context.Background(), testReq)
		if resp != nil {
			t.Fatalf("unexpected resp on failed DB call: %q", resp)
		}
		if want, got := err, dbError; want != got {
			t.Fatalf("wanted error: %q, got error: %q", want, got)
		}
	})
}

func TestDeleteBlog(t *testing.T) {
	beforeEach(t)
	defer mockCtrl.Finish()

	testReq := &blogProto.DeleteBlogRequest{
		Id: "1",
	}

	t.Run("calls DB with given data and returns expected response", func(t *testing.T) {
		testResp := &blogProto.DeleteBlogResponse{
			Id: "1",
		}

		mockDBClient.EXPECT().DeleteBlog(testReq).Return(testResp, nil).Times(1)

		resp, err := testServer.DeleteBlog(context.Background(), testReq)
		if err != nil {
			t.Fatalf("unexpected error in DB call: %q", err)
		}
		if want, got := testResp, resp; want != got {
			t.Fatalf("wanted response: %q, got response: %q", want, got)
		}
	})

	t.Run("returns error if DB call fails", func(t *testing.T) {
		dbError := errors.New("Mock DB error")

		mockDBClient.EXPECT().DeleteBlog(testReq).Return(nil, dbError).Times(1)

		resp, err := testServer.DeleteBlog(context.Background(), testReq)
		if resp != nil {
			t.Fatalf("unexpected resp on failed DB call: %q", resp)
		}
		if want, got := err, dbError; want != got {
			t.Fatalf("wanted error: %q, got error: %q", want, got)
		}
	})
}

func TestUpdateBlog(t *testing.T) {
	beforeEach(t)
	defer mockCtrl.Finish()

	testReq := &blogProto.UpdateBlogRequest{
		Id:      "1",
		Title:   "Test title",
		Content: "Test content",
	}

	t.Run("calls DB with given data and returns expected response", func(t *testing.T) {
		testResp := &blogProto.UpdateBlogResponse{
			Id:      "1",
			Title:   "Test title",
			Content: "Test content",
		}

		mockDBClient.EXPECT().UpdateBlog(testReq).Return(testResp, nil).Times(1)

		resp, err := testServer.UpdateBlog(context.Background(), testReq)
		if err != nil {
			t.Fatalf("unexpected error in DB call: %q", err)
		}
		if want, got := testResp, resp; want != got {
			t.Fatalf("wanted response: %q, got response: %q", want, got)
		}
	})

	t.Run("returns error if DB call fails", func(t *testing.T) {
		dbError := errors.New("Mock DB error")

		mockDBClient.EXPECT().UpdateBlog(testReq).Return(nil, dbError).Times(1)

		resp, err := testServer.UpdateBlog(context.Background(), testReq)
		if resp != nil {
			t.Fatalf("unexpected resp on failed DB call: %q", resp)
		}
		if want, got := err, dbError; want != got {
			t.Fatalf("wanted error: %q, got error: %q", want, got)
		}
	})
}

func TestGetBlog(t *testing.T) {
	beforeEach(t)
	defer mockCtrl.Finish()

	testReq := &blogProto.GetBlogRequest{
		Id: "1",
	}

	t.Run("calls DB with given data and returns expected response", func(t *testing.T) {
		testResp := &blogProto.GetBlogResponse{
			Id:      "1",
			Title:   "Test title",
			Content: "Test content",
		}

		mockDBClient.EXPECT().GetBlog(testReq).Return(testResp, nil).Times(1)

		resp, err := testServer.GetBlog(context.Background(), testReq)
		if err != nil {
			t.Fatalf("unexpected error in DB call: %q", err)
		}
		if want, got := testResp, resp; want != got {
			t.Fatalf("wanted response: %q, got response: %q", want, got)
		}
	})

	t.Run("returns error if DB call fails", func(t *testing.T) {
		dbError := errors.New("Mock DB error")

		mockDBClient.EXPECT().GetBlog(testReq).Return(nil, dbError).Times(1)

		resp, err := testServer.GetBlog(context.Background(), testReq)
		if resp != nil {
			t.Fatalf("unexpected resp on failed DB call: %q", resp)
		}
		if want, got := err, dbError; want != got {
			t.Fatalf("wanted error: %q, got error: %q", want, got)
		}
	})
}

func TestListBlog(t *testing.T) {
	beforeEach(t)
	defer mockCtrl.Finish()

	testReq := &blogProto.ListBlogRequest{
		Limit: 10,
	}

	testReqNoLimit := &blogProto.ListBlogRequest{}

	t.Run("calls DB with given Limit (if Limit provided) and returns expected response", func(t *testing.T) {
		testResp := &blogProto.ListBlogResponse{
			Blogs: []*blogProto.CreateBlogResponse{},
		}

		mockDBClient.EXPECT().ListBlog(testReq).Return(testResp, nil).Times(1)

		resp, err := testServer.ListBlog(context.Background(), testReq)
		if err != nil {
			t.Fatalf("unexpected error in DB call: %q", err)
		}
		if want, got := testResp, resp; want != got {
			t.Fatalf("wanted response: %q, got response: %q", want, got)
		}
	})

	t.Run("calls DB with Limit: 25 (if no Limit provided) and returns expected response", func(t *testing.T) {
		testResp := &blogProto.ListBlogResponse{
			Blogs: []*blogProto.CreateBlogResponse{},
		}

		mockDBClient.EXPECT().ListBlog(&blogProto.ListBlogRequest{Limit: 25}).Return(testResp, nil).Times(1)

		resp, err := testServer.ListBlog(context.Background(), testReqNoLimit)
		if err != nil {
			t.Fatalf("unexpected error in DB call: %q", err)
		}
		if want, got := testResp, resp; want != got {
			t.Fatalf("wanted response: %q, got response: %q", want, got)
		}
	})

	t.Run("returns error if DB call fails", func(t *testing.T) {
		dbError := errors.New("Mock DB error")

		mockDBClient.EXPECT().ListBlog(testReq).Return(nil, dbError).Times(1)

		resp, err := testServer.ListBlog(context.Background(), testReq)
		if resp != nil {
			t.Fatalf("unexpected resp on failed DB call: %q", resp)
		}
		if want, got := err, dbError; want != got {
			t.Fatalf("wanted error: %q, got error: %q", want, got)
		}
	})
}
