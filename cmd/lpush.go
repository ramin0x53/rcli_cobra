package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/wkoszek/rcli/config"
)

// lpushCmd represents the lpush command
var lpushCmd = &cobra.Command{
	Use:     "lpush",
	Short:   "Prepends a value to a list",
	Long:    `Prepends a value to a list`,
	Example: "rcli lpush names John",
	PreRun: func(cmd *cobra.Command, args []string) {
		config.CheckConfExit()
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := rdb.LPush(ctx, args[0], args[1]).Err()
		if err != nil {
			log.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(lpushCmd)

}
