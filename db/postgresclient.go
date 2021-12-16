package db

import (
	blogProto "blog-service/rpc/blog"
	"database/sql"
	"fmt"
	"strconv"

	"github.com/twitchtv/twirp"

	_ "github.com/lib/pq" // importing so drivers are registered with database/sql package, _ means we will not directly reference this package in code
)

/*
for reference:
https://pkg.go.dev/database/sql
https://www.calhoun.io/using-postgresql-with-go/
*/

type PostgresClient struct{}

var SqlDB *sql.DB

func NewPostgresClient() PostgresClient {
	return PostgresClient{}
}

func (p PostgresClient) Connect() error {
	fmt.Println("Connecting to Postgres")

	const (
		host     = "localhost"
		port     = 5432
		user     = "root"
		password = "password"
		dbname   = "root"
	)

	dbConnectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", dbConnectionString) // does not create connect to db, just validates arguments
	if err != nil {
		return err
	}

	err = db.Ping() // verifies connection to db, establishes connection if necessary
	if err != nil {
		return err
	}

	SqlDB = db

	fmt.Println("Successfully connected to Postgres")
	return nil
}

func (p PostgresClient) CreateBlog(data *blogProto.CreateBlogRequest) (*blogProto.CreateBlogResponse, error) {
	sqlStatement := "INSERT INTO blogs (title, content) VALUES ($1, $2) RETURNING id"
	id := 0
	err := SqlDB.QueryRow(sqlStatement, data.Title, data.Content).Scan(&id)
	if err != nil {
		return nil, twirp.NewError(twirp.InvalidArgument, fmt.Sprintf("There was an error creating a blog: %v", err))
	}

	return &blogProto.CreateBlogResponse{
		Id:      strconv.Itoa(id),
		Title:   data.Title,
		Content: data.Content,
	}, nil
}

func (p PostgresClient) GetBlog(data *blogProto.GetBlogRequest) (*blogProto.GetBlogResponse, error) {
	return &blogProto.GetBlogResponse{}, nil
}

func (p PostgresClient) UpdateBlog(data *blogProto.UpdateBlogRequest) (*blogProto.UpdateBlogResponse, error) {
	return &blogProto.UpdateBlogResponse{}, nil
}

func (p PostgresClient) DeleteBlog(data *blogProto.DeleteBlogRequest) (*blogProto.DeleteBlogResponse, error) {
	return &blogProto.DeleteBlogResponse{}, nil
}

func (p PostgresClient) ListBlog(data *blogProto.ListBlogRequest) (*blogProto.ListBlogResponse, error) {
	return &blogProto.ListBlogResponse{}, nil
}
