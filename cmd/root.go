/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"go-fetch/fetch"
	"os"

	"go-fetch/server"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-fetch",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// ./go-fetch fetch https://google.com https://meemeals.com
// go-fetch fetch https://google.com https://meemeals.com www.google.com
var fetchCmd = &cobra.Command{
	Use:   "fetch [urls...]",
	Short: "fetch can be considered a variadic go function, provide any number of strings as input, where strings need to be valid Urls",
	Args:  cobra.MinimumNArgs(1), // Requires at least one variadic string argument
	Run: func(cmd *cobra.Command, args []string) {
		// Access the variadic strings passed as arguments
		fmt.Println("Fetching urls:", args)
		var urls []string

		for _, arg := range args {
			// Perform actions with the variadic strings here
			urls = append(urls, arg)
		}
		res := fetch.GetUrlsSync(urls)

		// TODO: Add sync flag once endpoints have been successfully decoupled
		// var res interface{}
		// if useSync {
		// }
		// if syncd, ok := res.(fetch.Syncd); ok {
		fmt.Printf("Fetching urls was a %s\n", res.Status)
		fmt.Printf("Fetching urls took a time of %fs \n", res.Duration)
		// }
	},
}

// ./go-fetch validate https://google.com https://meemeals.com www.google.com
// go-fetch validate https://google.com https://meemeals.com www.google.com
var validateCmd = &cobra.Command{
	Use:   "validate [urls...]",
	Short: "validate can be considered a variadic go function, provide any number of strings as input.",
	Args:  cobra.MinimumNArgs(1), // Requires at least one variadic string argument
	Run: func(cmd *cobra.Command, args []string) {
		// Access the variadic strings passed as arguments
		fmt.Println("validating urls:", args)
		var urls []string
		var inValidUrls []string
		valid := true
		for _, arg := range args {
			// Perform actions with the variadic strings here
			urls = append(urls, arg)
			resp, err := fetch.GetUrl(arg)
			if err != nil || resp == nil {
				valid = false
				inValidUrls = append(inValidUrls, arg)
			}
		}

		if valid {
			fmt.Println("The urls are valid")
		} else {
			fmt.Println("The urls are not valid")
			fmt.Println("The invalid urls: ", inValidUrls)
		}
	},
}

// ./go-fetch run --port 8081
// go-fetch run --port 8081
func NewRun() *cobra.Command {
	var b struct {
		Port string
	}
	var runCmd = &cobra.Command{
		Use:   "run ...opts",
		Short: "runs the programs rest api to listen for api request on port 8080",
		Args:  cobra.MinimumNArgs(0), // Requires at least one variadic string argument
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args)
			fmt.Println("Port value", b.Port)
			// // TODO add configuration values for server run
			cfg := server.Config{
				Port: fmt.Sprintf(":%s", b.Port),
			}

			server.Run(cfg)
		},
	}

	// var sPort string
	runCmd.Flags().StringVar(
		&b.Port,
		"port",
		"8080",
		"Port to run on",
	)

	return runCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var useSync bool

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-fetch.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.AddCommand(fetchCmd)
	fetchCmd.Flags().BoolVar(&useSync, "useSync", false, "Use Sync library to fetch urls")
	rootCmd.AddCommand(validateCmd)

	rootCmd.AddCommand(NewRun())
}
