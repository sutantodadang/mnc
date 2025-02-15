package main

import (
	"context"
	"mnc/infrastructure"
	"mnc/internal/app/kafka/consumer"
	"mnc/internal/app/payment"
	"mnc/internal/app/topup"
	"mnc/internal/app/transaction"
	"mnc/internal/app/transfer"
	"mnc/internal/app/user"
	"mnc/internal/http/middlewares"
	"mnc/internal/http/route"
	"mnc/internal/repositories"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	log.Logger = zerolog.New(os.Stdout).With().Caller().Timestamp().Logger()
}

func main() {

	app := gin.Default()

	db := infrastructure.ConnectPG()

	defer db.Close()

	kafka, err := infrastructure.InitKafka()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initialize kafka")
	}

	defer kafka.Producer.Close()
	defer kafka.Consumer.Close()

	app.Use(middlewares.Trace())
	app.Use(middlewares.RequestLoggerMiddleware(), middlewares.ResponseLoggerMiddleware())

	setupContainer(app, db, kafka)

	server := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: app,
	}

	go func() {
		log.Info().Msgf("Starting server... on port %s", os.Getenv("PORT"))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("Server failed")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Info().Msg("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("Server forced to shutdown")
	}

	log.Info().Msg("Server exiting")

}

func setupContainer(app *gin.Engine, db *pgxpool.Pool, cfgKafka *infrastructure.KafkaConfig) {

	repo := repositories.New(db)

	middleware := middlewares.NewMiddleware(repo)

	// service
	userService := user.NewUserService(repo)
	topUpService := topup.NewTopUpService(repo)
	paymentService := payment.NewPaymentService(repo)
	transferService := transfer.NewTransferService(repo, cfgKafka)
	transactionService := transaction.NewTransactionService(repo)
	kafkaConsumer := consumer.NewKafkaConsumer(repo, transferService)

	// handler
	topUpHandler := topup.NewTopUpHandler(topUpService)
	userHandler := user.NewUserHandler(userService)
	paymentHandler := payment.NewPaymentHandler(paymentService)
	transferHandler := transfer.NewTransferHandler(transferService)
	transactionHandler := transaction.NewTransactionHandler(transactionService)

	// route
	route.RegisterUserRoute(app, userHandler, middleware)
	route.RegisterTopUpRoute(app, topUpHandler, middleware)
	route.RegisterPaymentRoute(app, paymentHandler, middleware)
	route.RegisterTransferRoute(app, transferHandler, middleware)
	route.RegisterTransactionRoute(app, transactionHandler, middleware)

	go cfgKafka.Consumer.Consume(context.Background(), []string{os.Getenv("KAFKA_TOPIC")}, kafkaConsumer)

}
