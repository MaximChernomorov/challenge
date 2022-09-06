package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/MaximChernomorov/challenge-test/internal/api"
	"github.com/MaximChernomorov/challenge-test/internal/repository"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var apiCmd = &cobra.Command{
	Use: "api",
	Run: func(cmd *cobra.Command, args []string) {
		startAPI()
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)
}

func startAPI() {
	e := echo.New()
	e.Logger.SetLevel(log.ERROR)
	e.Use(middleware.Logger())
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, // 1 KB
		LogLevel:  log.ERROR,
	}))

	repo, err := repository.NewPostgresRepo(viper.GetString("psqURL"))
	if err != nil {
		e.Logger.Fatal(err)
	}
	defer func() {
		err = repo.Close()
		if err != nil {
			e.Logger.Fatal(err)
		}
	}()
	APIInstance := &api.API{}
	APIInstance.SetRepo(repo)

	apiGroup := e.Group("/api")
	apiGroup.POST("/ip/locate", APIInstance.LocateIP)

	docsGroup := e.Group("/docs", middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == viper.Sub("docsAuth").GetString("user") && password == viper.Sub("docsAuth").GetString("pass") {
			return true, nil
		}
		return false, nil
	}))
	docsGroup.Static("/", "api/docs")

	go func() {
		e.Logger.Fatal(e.Start(viper.GetString("httpAddr")))
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM)
	signal.Notify(quit, syscall.SIGINT)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
