package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/wkoszek/rcli/config"
)

// hdelCmd represents the hdel command
var hdelCmd = &cobra.Command{
	Use:     "hdel",
	Short:   "Deletes one or more hash fields",
	Long:    `Deletes one or more hash fields`,
	Example: "rcli hdel user name",
	PreRun: func(cmd *cobra.Command, args []string) {
		config.CheckConfExit()
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := rdb.HDel(ctx, args[0], args[1:]...).Err()
		if err != nil {
			log.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(hdelCmd)

}
