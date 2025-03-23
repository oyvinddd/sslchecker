package app

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sslchecker/account"
	mw "sslchecker/middleware"
	"time"

	"github.com/julienschmidt/httprouter"
)

const (
	googleSignInPath   = "/api/v1/auth/google"
	googleCallbackPath = "/api/v1/auth/google/cb"
	accountPath        = "/api/v1/account"
)

const (
	readTimeout  = time.Second * 15
	writeTimeout = time.Second * 15
)

const (
	databaseConnectionKey = "DATABASE_CONNECTION"
	googleClientIDKey     = "GOOGLE_CLIENT_ID"
	googleClientSecretKey = "GOOGLE_CLIENT_SECRET"
)

type App struct {
	server http.Server
}

func New(config Configuration) *App {

	dbConn, err := config.EnvironmentVariable(databaseConnectionKey)
	googleClientID, err := config.EnvironmentVariable(googleClientIDKey)
	googleClientSecret, err := config.EnvironmentVariable(googleClientSecretKey)
	if err != nil {
		log.Fatalln(err.Error())
	}

	serviceFactory := NewMockedServiceFactory(context.Background(), dbConn)
	accountService := serviceFactory.AccountService()
	domainService := serviceFactory.DomainService()

	googleAuthHandler := account.NewGoogleAuthHandler(accountService, googleClientID, googleClientSecret)
	accountHandler := account.NewHandler(accountService)

	router := httprouter.New()

	// public routes
	router.GET(googleSignInPath, mw.Logger(googleAuthHandler.SignIn))
	router.GET(googleCallbackPath, mw.Logger(googleAuthHandler.Callback))

	// restricted routes
	router.GET(accountPath, mw.Logger(mw.Auth(accountHandler.GetAccount)))
	router.DELETE(accountPath, mw.Logger(mw.Auth(accountHandler.DeleteAccount)))

	return &App{
		server: http.Server{
			Addr:         config.Address,
			Handler:      router,
			ReadTimeout:  readTimeout,
			WriteTimeout: writeTimeout,
		},
	}
}

func (app *App) Run() error {
	fmt.Printf("Running HTTP server on %s...\n", app.server.Addr)
	return app.server.ListenAndServe()
}
