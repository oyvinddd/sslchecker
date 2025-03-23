package app

import (
	"anchorage/account"
		"anchorage/event"
			"context"
				"github.com/jackc/pgx/v5/pgxpool"
					"log"
					)

type (
	ServiceFactory interface {
				// AccountService return an instance of the account service
						AccountService() account.Service

								// EventService return an instance of the event service
										EventService() event.Service
											}

												mockServiceFactory struct {
															db *pgxpool.Pool
																}
																)

func NewMockedServiceFactory(ctx context.Context, connectionString string) ServiceFactory {
		db, err := pgxpool.New(ctx, connectionString)
			if err != nil {
						log.Fatalln(err.Error())
							}
								return &mockServiceFactory{db: db}
}

func (sf mockServiceFactory) AccountService() account.Service {
		return account.NewService(account.NewRepository(sf.db))
}

func (sf mockServiceFactory) DomainService() domain.Service {
		return domain.NewService(event.NewRepository(sf.db))
}
