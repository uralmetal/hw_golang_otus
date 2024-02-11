package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
	"github.com/uralmetal/hw_golang_otus/hw12_13_14_15_calendar/internal/app"
	"github.com/uralmetal/hw_golang_otus/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/uralmetal/hw_golang_otus/hw12_13_14_15_calendar/internal/server/http"
	memorystorage "github.com/uralmetal/hw_golang_otus/hw12_13_14_15_calendar/internal/storage/memory"
)

var configFile string

func init() {
	rootCmd.AddCommand(versionCmd)
	flag.StringVar(&configFile, "config", "configs/calendar.yaml", "Path to configuration file")
}

func init() {

}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Hugo",
	Long:  `All software has versions. This is Hugo's`,
	Run: func(cmd *cobra.Command, args []string) {
		printVersion()
	},
}

func main() {
	flag.Parse()

	//if flag.Arg(0) == "version" {
	//	printVersion()
	//	return
	//}
	config, err := NewConfig(configFile)
	if err != nil {
		fmt.Println("Error handle config:", err)
		os.Exit(1)
	}
	logg := logger.New(config.Logger.Level)
	storage := memorystorage.New()
	calendar := app.New(logg, storage)

	server := internalhttp.NewServer(logg, calendar)

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			logg.Error("failed to stop http server: " + err.Error())
		}
	}()

	logg.Info("calendar is running...")

	if err := server.Start(ctx); err != nil {
		logg.Error("failed to start http server: " + err.Error())
		cancel()
		os.Exit(1) //nolint:gocritic
	}
}
