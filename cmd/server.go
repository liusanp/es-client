package cmd

import (
	"es-client/commons"
	"es-client/models"
	"es-client/router"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var (
	config models.Config
)

// ServerCmd represents the server command
var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the server",
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func run() {
	var err error
	config, err = commons.InitESClient()
	if err != nil {
		log.Fatal(err)
	}
	r := router.Router()

	address := fmt.Sprintf(":%d", config.App.Port)
	log.Fatal(r.Run(address))
}

func init() {
	RootCmd.AddCommand(ServerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
