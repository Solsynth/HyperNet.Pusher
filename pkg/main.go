package main

import (
	"fmt"
	pkg "git.solsynth.dev/hypernet/pusher/pkg/internal"
	"git.solsynth.dev/hypernet/pusher/pkg/internal/gap"
	"git.solsynth.dev/hypernet/pusher/pkg/internal/grpc"
	"git.solsynth.dev/hypernet/pusher/pkg/internal/provider"
	"git.solsynth.dev/hypernet/pusher/pkg/internal/scheduler"
	"github.com/fatih/color"
	"os"
	"os/signal"
	"syscall"

	"github.com/robfig/cron/v3"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
}

func main() {
	// Booting screen
	fmt.Println(color.YellowString(` ____            _
|  _ \ _   _ ___| |__   ___ _ __
| |_) | | | / __| '_ \ / _ \ '__|
|  __/| |_| \__ \ | | |  __/ |
|_|    \__,_|___/_| |_|\___|_|`))
	fmt.Printf("%s v%s\n", color.New(color.FgHiYellow).Add(color.Bold).Sprintf("Hypernet.Pusher"), pkg.AppVersion)
	fmt.Printf("The notification / email delivery service in Hypernet\n")
	color.HiBlack("=====================================================\n")

	// Configure settings
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.SetConfigName("settings")
	viper.SetConfigType("toml")

	// Load settings
	if err := viper.ReadInConfig(); err != nil {
		log.Panic().Err(err).Msg("An error occurred when loading settings.")
	}

	// Connect to nexus
	if err := gap.InitializeToNexus(); err != nil {
		log.Fatal().Err(err).Msg("An error occurred when connecting to nexus...")
	} else {
		log.Info().Msg("Connected to nexus successfully!")
	}

	// Initialize pusher conn
	fcmCredentials := viper.GetString("provider.fcm.credentials")
	if len(fcmCredentials) > 0 {
		if err := provider.InitFCM(fcmCredentials); err != nil {
			log.Fatal().Err(err).Msg("An error occurred when initializing FCM connection...")
		} else {
			log.Info().Msg("Pusher conn with FCM is initialized!")
		}
	} else {
		log.Warn().Msg("Pusher conn with FCM was not configured...")
	}
	apnCredentials := viper.GetString("provider.apns.credentials")
	if len(apnCredentials) > 0 {
		key := viper.GetString("provider.apns.key")
		team := viper.GetString("provider.apns.team")
		topic := viper.GetString("provider.apns.topic")
		if err := provider.InitAPN(apnCredentials, key, team, topic); err != nil {
			log.Fatal().Err(err).Msg("An error occurred when initializing APN connection...")
		} else {
			log.Info().Msg("Pusher conn with APN is initialized!")
		}
	} else {
		log.Warn().Msg("Pusher conn with APN was not configured...")
	}

	// Subscribe to MQ
	if err := scheduler.SubscribeToQueue(); err != nil {
		log.Error().Err(err).Msg("Unable to subscribe to MQ via nexus, cannot get push requests from MQ...")
	} else {
		log.Info().Msg("Subscribed to MQ!")
	}

	// Grpc Server
	go grpc.NewServer().Listen()

	// Configure timed tasks
	quartz := cron.New(cron.WithLogger(cron.VerbosePrintfLogger(&log.Logger)))
	quartz.Start()

	// Messages
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	quartz.Stop()
}
