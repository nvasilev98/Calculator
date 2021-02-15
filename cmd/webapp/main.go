package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	"github.com/nvasilev98/calculator/cmd/webapp/env"
	"github.com/nvasilev98/calculator/cmd/webapp/internal/calculator/controller"
	"github.com/nvasilev98/calculator/cmd/webapp/internal/calculator/presenter"
	"github.com/nvasilev98/calculator/pkg/calculation"
	"github.com/nvasilev98/calculator/pkg/database/connection"
	"github.com/nvasilev98/calculator/pkg/database/dataaccess"
	"github.com/nvasilev98/calculator/pkg/database/postgres"
	"github.com/pkg/errors"
)

func main() {

	dbConfig, err := connection.Load()
	if err != nil {
		panic(errors.Wrap(err, "failed to get db variables"))
	}
	db, err := connection.ConnectDB(dbConfig)

	if err != nil {
		panic(errors.Wrap(err, "failed to connect to database"))
	}

	defer db.Close()

	cfg, err := env.Load()
	if err != nil {
		panic(errors.Wrap(err, "failed to get variables"))
	}

	router := gin.Default()

	var operationfactory = calculation.NewOperationFactory()
	var reader = calculation.NewParser(operationfactory)
	var converter = calculation.NewReverseNotation(operationfactory)

	calculationDAO, err := postgres.NewCalculation(db)
	if err != nil {
		panic(errors.Wrap(err, "failed to create calculation dao"))
	}

	calculator := calculation.NewCalculator(operationfactory, reader, converter)
	controller := controller.NewController(calculator, dataaccess.NewDataAccess(db, calculationDAO))
	presenter := presenter.NewPresenter(controller)

	calculatorAuth := router.Group("/", gin.BasicAuth(gin.Accounts{
		cfg.Username: cfg.Password,
	}))

	calculatorAuth.POST("/calculate", presenter.CalculateExpression)
	calculatorAuth.POST("/create", presenter.InsertExpression)
	calculatorAuth.GET("/get/:id", presenter.GetExpression)
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(errors.Wrap(err, "listen and serve failed"))
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		panic(errors.Wrap(err, "failed to shutdown server"))
	}

	log.Println("Server exiting")
}
