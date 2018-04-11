package cli

import (
	"github.com/spf13/viper"
	"github.com/spf13/cobra"

	"fmt"
	"os"
	"github.com/mitchellh/go-homedir"
)

type CLIStruct struct {
	rootCmd *cobra.Command
	versionCmd *cobra.Command
	cfgFile string
	flagstr1 string

}


func (cli *CLIStruct) init() {

	//cobra.OnInitialize(cmd.initConfig)mp
	cli.rootCmd = &cobra.Command{
		Use:   "mp3tag",
		Short: "Hugo is a very fast static site generator",
		Long: `mp3tag is for updating the tags in mp3
				`,

		Args: cobra.ArbitraryArgs,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here

			fmt.Println("This is main cmd line utility", args, cli.rootCmd.Flags().Lookup("myflag1").Value.String())
		},
	}

	cli.rootCmd.Flags().StringVar(&cli.flagstr1, "myflag1", "defval1", "how to use")
	cli.rootCmd.PersistentFlags().StringVar(&cli.cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	//rootCmd.PersistentFlags().StringVarP(&projectBase, "projectbase", "b", "", "base project directory eg. github.com/spf13/")
	cli.rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "Author name for copyright attribution")
	//rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "Name of license for the project (can provide `licensetext` in config)")
	cli.rootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")


	cli.versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Print the version number",
		Long:  `All software has versions. This is Hugo's`,

		Args: cobra.ArbitraryArgs,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
		},
	}

	cli.rootCmd.AddCommand(cli.versionCmd)

	/*	viper.BindPFlag("author", cmd.rootCmd.PersistentFlags().Lookup("author"))
		viper.BindPFlag("projectbase", cmd.rootCmd.PersistentFlags().Lookup("projectbase"))
		viper.BindPFlag("useViper", cmd.rootCmd.PersistentFlags().Lookup("viper"))
		viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
		viper.SetDefault("license", "apache")
	*/
}

func (cli *CLIStruct) initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	if cli.cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cli.cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cobra")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}


var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "newApp",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.newApp.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".newApp" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".newApp")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
