package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kelseyhightower/envconfig"
	"github.com/samuell20/FruitTracker/internal/events"
	create_product "github.com/samuell20/FruitTracker/internal/operations/create/product"
	create_user "github.com/samuell20/FruitTracker/internal/operations/create/user"
	delete_product "github.com/samuell20/FruitTracker/internal/operations/delete/product"
	delete_user "github.com/samuell20/FruitTracker/internal/operations/delete/user"
	product_query "github.com/samuell20/FruitTracker/internal/operations/get/product"
	user_query "github.com/samuell20/FruitTracker/internal/operations/get/user"
	update_product "github.com/samuell20/FruitTracker/internal/operations/update/product"
	update_user "github.com/samuell20/FruitTracker/internal/operations/update/user"

	"github.com/samuell20/FruitTracker/internal/platform/bus/inmemory"
	"github.com/samuell20/FruitTracker/internal/platform/server"
	product_repository "github.com/samuell20/FruitTracker/internal/platform/storage/mysql/product"
	user_repository "github.com/samuell20/FruitTracker/internal/platform/storage/mysql/user"
)

func Run() error {
	var cfg config
	err := envconfig.Process("WEB", &cfg)
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

	//REPOSITORIES
	productRepository := product_repository.NewProductRepository(db, cfg.DbTimeout)
	userRepository := user_repository.NewUserRepository(db, cfg.DbTimeout)

	//QUERY SERVICES
	queryServices := map[string]interface{}{}
	queryServices["products"] = product_query.NewProductQuery(productRepository)
	queryServices["users"] = user_query.NewUserQuery(userRepository)

	//COMMAND SERVICES
	createProductService := create_product.NewProductService(productRepository, eventBus)
	deleteProductService := delete_product.NewDeleteProductService(productRepository, eventBus)
	updateProductService := update_product.NewUpdateProductService(productRepository, eventBus)
	createUserService := create_user.NewUserService(userRepository, eventBus)
	deleteUserService := delete_user.NewDeleteUserService(userRepository, eventBus)
	updateUserService := update_user.NewUpdateUserService(userRepository, eventBus)

	//HANDLERS
	createProductCommandHandler := create_product.NewProductCommandHandler(createProductService)
	deleteProductCommandHandler := delete_product.NewDeleteProductCommandHandler(deleteProductService)
	updateProductCommandHandler := update_product.NewUpdateProductCommandHandler(updateProductService)
	createUserCommandHandler := create_user.NewUserCommandHandler(createUserService)
	deleteUserCommandCommand := delete_user.NewDeleteUserCommandHandler(deleteUserService)
	updateUserCommandHandler := update_user.NewUpdateUserCommandHandler(updateUserService)

	//COMMAND REGISTRATION
	commandBus.Register(create_product.ProductCommandType, createProductCommandHandler)
	commandBus.Register(delete_product.DeleteProductCommandType, deleteProductCommandHandler)
	commandBus.Register(update_product.UpdateProductCommandType, updateProductCommandHandler)

	commandBus.Register(create_user.UserCommandType, createUserCommandHandler)
	commandBus.Register(delete_user.DeleteUserCommandType, deleteUserCommandCommand)
	commandBus.Register(update_user.UpdateUserCommandType, updateUserCommandHandler)
	//EVENT SUBSCRIPTION
	eventBus.Subscribe(events.ProductCreatedEventType, create_product.NewTestOnProductCreated())

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
