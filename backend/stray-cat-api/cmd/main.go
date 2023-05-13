package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"stray-cat-api/db"
	"stray-cat-api/infrastructure/persistence/postgres"
	"stray-cat-api/presentation/handler"
	"stray-cat-api/usecase/interactor"
	"strings"

	"net/http"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// Read environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	// Create database connection string
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%s", dbUser, dbPassword, dbName, dbHost, dbPort)

	// Run database migrations
	db.RunMigrations(connStr)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Initialize handler
	userInfoRepository := postgres.NewUserInfoRepository(db)
	userInteractor := interactor.UserInfoInteractor{
		UserInfoRepository: userInfoRepository,
	}
	userHandler := handler.UserInfoHandler{
		UserInfoInteractor: userInteractor,
	}

	strayCatRepository := postgres.NewStrayCatRepository(db)
	strayCatInteractor := interactor.StrayCatInteractor{
		StrayCatRepository: strayCatRepository,
	}
	strayCatHandler := handler.StrayCatHandler{
		StrayCatInteractor: strayCatInteractor,
	}

	// Initialize Firebase
	// app := initFirebase()

	r := gin.Default()

	// CORS
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type"}

	r.Use(cors.New(config))

	// strayCats := r.Group("/stray-cats")
	// {
	// 	strayCats.POST("/", strayCatHandler.Create)
	// }

	r.GET("/users", userHandler.GetAll)
	r.GET("/users/:userId", userHandler.GetByID)
	r.POST("/users/:userId", userHandler.Create)
	r.POST("/stray-cats", strayCatHandler.Create)
	r.GET("/stray-cats/search", strayCatHandler.GetAll)
	r.GET("/stray-cats/:catId", strayCatHandler.GetByID)

	// api := r.Group("/api")
	// {
	// 	//
	// 	v1 := api.Group("/v1")
	// 	{
	// 		// 認証が必要なエンドポイント
	// 		authUsers := v1.Group("/auth-users")
	// 		authUsers.Use(FirebaseAuthMiddleware(app))
	// 		{
	// 			authUsers.GET("/", userHandler.GetAll)
	// 			authUsers.GET("/:id", userHandler.GetByID)
	// 			authUsers.POST("/", userHandler.Create)
	// 			authUsers.PUT("/:id", userHandler.Update)
	// 			authUsers.DELETE("/:id", userHandler.Delete)
	// 		}
	// 	}
	// }

	r.Run(":8080")
}

// Initialize Firebase
func initFirebase() *firebase.App {
	saKeyPath := "serviceAccountKey.json"
	opt := option.WithCredentialsFile(saKeyPath)

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("Error initializing Firebase: %v", err)
	}

	return app
}

// FirebaseAuthMiddleware is a middleware for Firebase authentication
func FirebaseAuthMiddleware(app *firebase.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		log.Printf(authHeader)
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header not found"})
			c.Abort()
			return
		}

		idToken := strings.TrimSpace(strings.Replace(authHeader, "Bearer", "", 1))

		log.Printf(idToken)

		client, err := app.Auth(context.Background())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error initializing Auth client"})
			c.Abort()
			return
		}

		token, err := client.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid ID token"})
			c.Abort()
			return
		}

		c.Set("uid", token.UID)

		c.Next()
	}
}
