// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/castillofranciscodaniel/golang-beers/app"
	"github.com/castillofranciscodaniel/golang-beers/domain"
	"github.com/castillofranciscodaniel/golang-beers/provider"
)

// Injectors from wire.go:

func InitializeServer() app.ContainerServiceImp {
	healthHandler := app.NewHealthHandler()
	dbManagerPostgres := provider.NewDbManagerImpl()
	beerRepositoryDb := domain.NewBeersRepositoryDb(dbManagerPostgres)
	currencyClientDefault := provider.NewCurrencyClientDefault()
	beerService := domain.NewBeersServiceDefault(beerRepositoryDb, currencyClientDefault)
	beerHandler := app.NewBeersHandler(beerService)
	containerServiceImp := app.NewContainerServiceImp(healthHandler, beerHandler)
	return containerServiceImp
}