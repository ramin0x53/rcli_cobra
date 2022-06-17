package cmd

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis/v9"
	"github.com/spf13/cobra"
	"github.com/wkoszek/rcli/config"
)

var (
	configData = config.ReadConfig()
	rdb        = redis.NewClient(&redis.Options{
		Addr:     configData.Ip,
		Password: configData.Password,
		DB:       IntConv(configData.Db),
	})
	ctx = context.Background()
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rcli",
	Short: "Easy to use Redis CLI that can be used from the command line mode",
}

func IntConv(text string) int {
	s, err := strconv.Atoi(text)
	if err != nil {
		log.Println(err)
	}
	return s
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
