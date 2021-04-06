package cmd

import (
	"github.com/Eldius/k3s-dashboard-go/server"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts dashboard server",
	Long: `Starts dashboard server.
For example:

k3s-dashboard-go -p <port-number>

`,
	Run: func(cmd *cobra.Command, args []string) {
		server.Start(startPort)
	},
}

var (
	startPort int
)

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().IntVarP(&startPort, "port", "p", 8080, "-p 8080")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
