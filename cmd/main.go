package main

import (
	"log"
	"os"

	"github.com/ShadamHarizky/dating-app/infrastructure/auth"
	"github.com/ShadamHarizky/dating-app/infrastructure/database"
	"github.com/ShadamHarizky/dating-app/infrastructure/redis"
	"github.com/ShadamHarizky/dating-app/interfaces"
	"github.com/ShadamHarizky/dating-app/interfaces/middleware"
	"github.com/ShadamHarizky/dating-app/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	//To load our environmental variables.
	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}
}

func main() {

	dbdriver := os.Getenv("DB_DRIVER")
	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASSWORD")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	//redis details
	redis_host := os.Getenv("REDIS_HOST")
	redis_port := os.Getenv("REDIS_PORT")
	redis_password := os.Getenv("REDIS_PASSWORD")

	repositories, err := database.NewRepositories(dbdriver, user, password, port, host, dbname)
	if err != nil {
		panic(err)
	}
	defer repositories.Close()
	repositories.Automigrate()
	redisDb, err := redis.NewRedisDB(redis_host, redis_port, redis_password)
	if err != nil {
		log.Fatal(err)
	}

	swipeServices, err := service.NewSwipeService(*redisDb, repositories.User)
	if err != nil {
		panic(err)
	}

	userServices, err := service.NewUserService(repositories.User, *swipeServices)
	if err != nil {
		panic(err)
	}

	authorization, err := auth.NewAuth(redisDb.Client)
	if err != nil {
		log.Fatal(err)
	}

	tk := auth.NewToken()

	users := interfaces.NewUsers(userServices, authorization, tk)
	swipe := interfaces.NewSwipes(swipeServices, authorization, tk)
	authenticate := interfaces.NewAuthenticate(userServices, authorization, tk)

	r := gin.Default()
	r.Use(middleware.CORSMiddleware()) //For CORS

	// this api is for user register
	r.POST("/register", users.SaveUser)

	// this api is for get all dating profiles exclude profile when the user already swipe and if the user is not premium and reach the limit per day the api will be response empty.
	r.GET("/profile", users.GetUsers)

	// this api is when the user want to check detail dating profile and if the user is not premium then reach the limit per day the api cannot used.
	r.GET("/profile/:profile_id", users.GetUser)

	// this api is for the user want to got the premium feature, one of the feature is there is no limit daily, free how much you want to swipe
	r.POST("/profile/purchase/premium", users.PurchasePremium)

	// this api is for login application, all the api require jwt token and that is to get jwt token
	r.POST("/login", authenticate.Login)

	// this api is for logout app and delete tokens
	r.POST("/logout", authenticate.Logout)

	// this api is a function that uses the refresh_token to generate new pairs of refresh and access tokens.
	r.POST("/refresh", authenticate.Refresh)

	r.POST("/swipe", swipe.Swipes)

	//Starting the application
	app_port := os.Getenv("PORT") //using heroku host
	if app_port == "" {
		app_port = "8888" //localhost
	}
	log.Fatal(r.Run(":" + app_port))
}
