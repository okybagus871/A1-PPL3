package main

import (
	"database/sql"
	"BackEnd/user"

	"github.com/go-kit/kit/log"
	_ "github.com/lib/pq"
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log/level"
)

var db *sql.DB
var err error

func main(){
	// dsn := "host=localhost user=postgres password=howl94 dbname=user port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// if err != nil {
	// 	log.Fatal("DB error")
	// }

	// db.AutoMigrate(&service.UserProfile{})
	dsn := "host=localhost user=postgres password=howl94 dbname=user port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	var httpAddr = flag.String("http", ":8080", "http listen address")

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "usermgmt",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	var db *sql.DB
	{
		var err error

		db, err = sql.Open("postgres", dsn)
		if err != nil {
			level.Error(logger).Log("exit", err)
			os.Exit(-1)
		}

	}

	flag.Parse()
	ctx := context.Background()
	var srv user.UserService
	{
		repository := user.NewRepo(db, logger)

		srv = user.NewService(repository, logger)
	}

	errs := make(chan error)
	endpoints := user.MakeEndpoints(srv)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	go func() {
		fmt.Println("listening on port", *httpAddr)
		handler := user.NewHTTPServer(ctx, endpoints)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	level.Error(logger).Log("exit", <-errs)
}