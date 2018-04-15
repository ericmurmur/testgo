package cli

import (
	"github.com/spf13/cobra"
	"fmt"
)




// startCmd represents the start command
var helpCmd = &cobra.Command{
	Use:   "myhelp",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("help called")
	},
}

func init() {
	rootCmd.AddCommand(helpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}