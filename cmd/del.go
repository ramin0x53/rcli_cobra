package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/wkoszek/rcli/config"
)

// delCmd represents the del command
var delCmd = &cobra.Command{
	Use:     "del",
	Short:   "This command deletes the key, if it exists",
	Long:    `This command deletes the key, if it exists`,
	Example: "rcli del name",
	PreRun: func(cmd *cobra.Command, args []string) {
		config.CheckConfExit()
	},
	Run: func(cmd *cobra.Command, args []string) {
		err := rdb.Del(ctx, args...).Err()
		if err != nil {
			log.Println(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(delCmd)

}
