package db

import (
	blogProto "blog-service/rpc/blog"
	"context"
	"fmt"
	"time"

	"github.com/twitchtv/twirp"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
for reference:
https://pkg.go.dev/go.mongodb.org/mongo-driver@v1.7.2/mongo
https://docs.mongodb.com/drivers/go/current/quick-start/
https://docs.mongodb.com/drivers/go/current/fundamentals/bson/
https://docs.mongodb.com/drivers/go/current/fundamentals/crud/
*/

type MongoClient struct{}

var Collection *mongo.Collection

type BlogItem struct {
	Id      primitive.ObjectID `bson:"_id"`
	Title   string             `bson:"title"`
	Content string             `bson:"content"`
}

func NewMongoClient() MongoClient {
	return MongoClient{}
}

func (m MongoClient) Connect() error {
	fmt.Println("Connecting to MongoDB")

	// returned cancel function will cancel the created ctx and all associated resources, so ensures cleanup once db operations complete
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// mongo.Connect will create a new client and enable access to the MongoDB instance running on 27107
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return err
	}

	Collection = client.Database("mydb").Collection("blog")

	fmt.Println("Successfully connected to MongoDB")
	return nil
}

func (m MongoClient) CreateBlog(data *blogProto.CreateBlogRequest) (*blogProto.CreateBlogResponse, error) {
	res, err := Collection.InsertOne(context.Background(), data)
	if err != nil {
		return nil, twirp.NewError(twirp.InvalidArgument, fmt.Sprintf("There was an error creating a blog: %v", err))
	}

	// type assertion that res.InsertedID is of type primitive.ObjectID
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, twirp.NewError(twirp.InvalidArgument, fmt.Sprintf("Cannot convert to oid: %v", ok))
	}
	return &blogProto.CreateBlogResponse{
		Id:      oid.Hex(),
		Title:   data.Title,
		Content: data.Content,
	}, nil
}

func (m MongoClient) GetBlog(data *blogProto.GetBlogRequest) (*blogProto.GetBlogResponse, error) {
	return &blogProto.GetBlogResponse{}, nil
}

func (m MongoClient) UpdateBlog(data *blogProto.UpdateBlogRequest) (*blogProto.UpdateBlogResponse, error) {
	return &blogProto.UpdateBlogResponse{}, nil
}

func (m MongoClient) DeleteBlog(data *blogProto.DeleteBlogRequest) (*blogProto.DeleteBlogResponse, error) {
	return &blogProto.DeleteBlogResponse{}, nil
}

func (m MongoClient) ListBlog(data *blogProto.ListBlogRequest) (*blogProto.ListBlogResponse, error) {
	return &blogProto.ListBlogResponse{}, nil
}
