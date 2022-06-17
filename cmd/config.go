package cmd

import (
	"github.com/spf13/cobra"
	"github.com/wkoszek/rcli/config"
)

var (
	serverIP string
	password string
	db       string
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:     "config",
	Short:   "Set ip address and port number for redis server",
	Long:    `Set ip address and port number for redis server`,
	Example: "rcli config --ip 192.168.1.5:6379 --pass 123456",
	Run: func(cmd *cobra.Command, args []string) {
		config.WriteConfig(serverIP, password, db)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.Flags().StringVar(&serverIP, "ip", "", "Redis server ip and port [ip:port]")
	configCmd.Flags().StringVar(&password, "pass", "", "Redis server password")
	configCmd.Flags().StringVar(&db, "db", "0", "Redis server db number")

}
