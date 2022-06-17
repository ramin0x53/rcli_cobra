package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/wkoszek/rcli/config"
)

// hmsetCmd represents the hmset command
var hmsetCmd = &cobra.Command{
	Use:     "hmset",
	Short:   "Sets multiple hash fields to multiple values",
	Long:    `Sets multiple hash fields to multiple values`,
	Example: "rcli hmset user name John",
	PreRun: func(cmd *cobra.Command, args []string) {
		config.CheckConfExit()
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := rdb.HMSet(ctx, args[0], args[1:]).Err()
		if err != nil {
			log.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(hmsetCmd)

}
