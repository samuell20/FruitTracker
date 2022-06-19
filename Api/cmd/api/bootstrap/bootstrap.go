package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"
	"github.com/samuell20/FruitTracker/internal/events"
	"github.com/samuell20/FruitTracker/internal/operations/create/product"
	product_query "github.com/samuell20/FruitTracker/internal/operations/get/product"
	user_query "github.com/samuell20/FruitTracker/internal/operations/get/user"
	"github.com/samuell20/FruitTracker/internal/platform/bus/inmemory"
	"github.com/samuell20/FruitTracker/internal/platform/server"
	product_repository "github.com/samuell20/FruitTracker/internal/platform/storage/mysql/product"
	user_repository "github.com/samuell20/FruitTracker/internal/platform/storage/mysql/user"
)

func Run() error {
	var cfg config
	err := envconfig.Process("MOOC", &cfg)
	if err != nil {
		return err
	}

	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.DbUser, cfg.DbPass, cfg.DbHost, cfg.DbPort, cfg.DbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	var (
		commandBus = inmemory.NewCommandBus()
		eventBus   = inmemory.NewEventBus()
	)
	queryServices := map[string]interface{}{}
	productRepository := product_repository.NewProductRepository(db, cfg.DbTimeout)

	queryServices["products"] = product_query.NewProductQuery(productRepository)
	createProductService := product.NewProductService(productRepository, eventBus)

	userRepository := user_repository.NewUserRepository(db, cfg.DbTimeout)
	queryServices["users"] = user_query.NewUserQuery(userRepository)

	createProductCommandHandler := product.NewProductCommandHandler(createProductService)
	commandBus.Register(product.ProductCommandType, createProductCommandHandler)

	eventBus.Subscribe(events.ProductCreatedEventType, product.NewTestOnProductCreated())

	ctx, srv := server.New(context.Background(), cfg.Host, cfg.Port, cfg.ShutdownTimeout, commandBus, queryServices)
	return srv.Run(ctx)
}

type config struct {
	// Server configuration
	Host            string        `default:"0.0.0.0"`
	Port            uint          `default:"9000"`
	ShutdownTimeout time.Duration `default:"10s"`
	// Database configuration
	DbUser    string        `default:"fruittracker"`
	DbPass    string        `default:"fruittracker"`
	DbHost    string        `default:"db"`
	DbPort    uint          `default:"3306"`
	DbName    string        `default:"fruittracker"`
	DbTimeout time.Duration `default:"5s"`
}
