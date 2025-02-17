package app

import (
	"anchorage/account"
	"anchorage/event"
	mw "anchorage/middleware"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

const (
	googleSignInPath   = "/api/v1/auth/google"
	googleCallbackPath = "/api/v1/auth/google/callback"
	accountPath        = "/api/v1/account"
	eventsPath         = "/api/v1/events"
	eventPath          = "/api/v1/events/:id"
	eventTicketsPath   = "/api/v1/events/:id/tickets"
	eventTicketPath    = "/api/v1/events/:id/tickets/:tid"
	ticketPath         = "/api/v1/tickets/:id"
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
	eventService := serviceFactory.EventService()

	googleAuthHandler := account.NewGoogleAuthHandler(accountService, googleClientID, googleClientSecret)
	accountHandler := account.NewHandler(accountService)
	eventHandler := event.NewHandler(eventService)

	router := httprouter.New()

	// public routes
	router.GET(googleSignInPath, mw.Logger(googleAuthHandler.SignIn))
	router.GET(googleCallbackPath, mw.Logger(googleAuthHandler.Callback))
	router.GET(eventPath, mw.Logger(eventHandler.GetEvent))
	router.POST(eventTicketsPath, mw.Logger(eventHandler.CreateTicket))

	// restricted routes
	router.GET(accountPath, mw.Logger(mw.Auth(accountHandler.GetAccount)))
	router.DELETE(accountPath, mw.Logger(mw.Auth(accountHandler.DeleteAccount)))
	router.POST(eventsPath, mw.Logger(mw.Auth(eventHandler.CreateEvent)))
	router.GET(eventsPath, mw.Logger(mw.Auth(eventHandler.ListEvents)))
	router.DELETE(eventPath, mw.Logger(mw.Auth(eventHandler.DeleteEvent)))
	router.GET(eventTicketsPath, mw.Logger(mw.Auth(eventHandler.ListTickets)))
	router.PUT(ticketPath, mw.Logger(mw.Auth(eventHandler.ValidateTicket)))

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
