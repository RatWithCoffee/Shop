package internal

import (
	"context"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"graphql/internal/custom_models"
	graph2 "graphql/internal/graph"
	session2 "graphql/internal/session"
	"graphql/internal/user"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func GetApp() http.Handler {
	dsn := "host=localhost user=rat dbname=shop password=12345 port=8008 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true, Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		panic(err)
	}
	initDb(db)

	config := graph2.Config{Resolvers: &graph2.Resolver{Db: db}}
	jwtSession := session2.JwtSession{db}
	userRepo := user.UserRepo{Db: db}
	userHandler := user.UserHandler{&jwtSession, &userRepo}
	config.Directives.Authorized = authorized

	mux := http.NewServeMux()
	srv := handler.NewDefaultServer(graph2.NewExecutableSchema(config))
	mux.Handle("/playground", playground.Handler("GraphQL playground", "/query"))
	mux.Handle("/query", srv)

	auth := session2.AuthMiddleware(mux, &jwtSession)
	muxAuth := http.NewServeMux()
	muxAuth.Handle("/", auth)
	muxAuth.HandleFunc("/register", userHandler.RegistrationHandler)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", defaultPort)
	return muxAuth
}

func authorized(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
	_, ok := session2.SessionFromCtx(ctx)
	if !ok {
		return nil, fmt.Errorf("User not authorized")
	}

	return next(ctx)
}

func initDb(db *gorm.DB) {
	c, ioErr := os.ReadFile("./build/_postgresql/init_db.sql")
	if ioErr != nil {
		panic(ioErr)
	}
	sql := string(c)
	err := db.Exec(sql).Error
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&user.User{}, &custom_models.CartItem{})
}
