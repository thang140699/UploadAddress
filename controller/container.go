package main

import (
	database "WeddingUtilities/database/mongo"
	"WeddingUtilities/database/repository"
	"WeddingUtilities/utilities/provider/jwt"
	"WeddingUtilities/utilities/provider/mongo"
)

type Provider struct {
	*mongo.MongoProvider
	*jwt.JWTService
}

type Container struct {
	*Provider

	Config            Config
	AddressRepository repository.AddressRepository
}

func NewContainer(config Config) (*Container, error) {
	container := new(Container)
	err := container.InitContainer(config)
	if err != nil {
		return nil, err
	}

	container.Config = config

	return container, nil
}

func (container *Container) InitContainer(config Config) error {
	// Load providers into container
	err := container.LoadProviders(config)
	if err != nil {
		return err
	}

	// Load repositories
	container.LoadRepositoryImplementations(config)

	return nil
}

func (container *Container) LoadProviders(config Config) error {

	// redisProvider := redis.NewRedisProviderFromURL(config.RedisURL)
	mongoProvider := mongo.NewMongoProviderFromURL(config.MongoURL)

	container.Provider = &Provider{
		MongoProvider: mongoProvider,
		// RedisProvider: redisProvider,
		JWTService: jwt.NewJWT(config.JwtKey),
	}
	return nil
}
func (container *Container) LoadRepositoryImplementations(config Config) {
	container.AddressRepository = database.NewAddressMongoRepository(container.MongoProvider)
}
