package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/wkoszek/rcli/config"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:                "get",
	Short:              "Gets the value of a key",
	Long:               `Gets the value of a key`,
	Example:            "rcli get name",
	DisableFlagParsing: false,
	PreRun: func(cmd *cobra.Command, args []string) {
		config.CheckConfExit()
	},
	Run: func(cmd *cobra.Command, args []string) {
		val, err := rdb.Get(ctx, args[0]).Result()
		if err != nil {
			log.Println(err)
		}
		fmt.Println(val)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

}
