package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/wkoszek/rcli/config"
)

// hgetCmd represents the hget command
var hgetCmd = &cobra.Command{
	Use:     "hget",
	Short:   "Gets the value of a hash field stored at the specified key",
	Long:    `Gets the value of a hash field stored at the specified key`,
	Example: "rcli hget user name",
	PreRun: func(cmd *cobra.Command, args []string) {
		config.CheckConfExit()
	},
	Run: func(cmd *cobra.Command, args []string) {
		val, err := rdb.HGet(ctx, args[0], args[1]).Result()
		if err != nil {
			log.Println(err)
		}
		fmt.Println(val)
	},
}

func init() {
	rootCmd.AddCommand(hgetCmd)

}
