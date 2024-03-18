package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/uralmetal/hw_golang_otus/hw12_13_14_15_calendar/internal/app"
	configHandler "github.com/uralmetal/hw_golang_otus/hw12_13_14_15_calendar/internal/config"
	"github.com/uralmetal/hw_golang_otus/hw12_13_14_15_calendar/internal/logger"
	internalhttp "github.com/uralmetal/hw_golang_otus/hw12_13_14_15_calendar/internal/server/http"
	memorystorage "github.com/uralmetal/hw_golang_otus/hw12_13_14_15_calendar/internal/storage/memory"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var configFile string

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.PersistentFlags().StringVar(&configFile, "config", "configs/calendar.yaml", "config file (default is configs/calendar.yaml)")
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run application",
	Long:  `All software has versions. This is application`,
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

type Config struct {
	Logger configHandler.LoggerConf `yaml:"logger"`
	// TODO
}

func NewConfig(path string) (Config, error) {
	var config Config
	err := configHandler.ParseConfig(path, &config)
	return config, err
}

func run() {
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
