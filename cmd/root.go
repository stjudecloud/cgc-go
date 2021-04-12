package cmd

import (
	"github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/stjudecloud/cgc-go/apps"
	"github.com/stjudecloud/cgc-go/rate"
)

var cfgFile string
var verbose bool

func init() {
	log.SetLevel(log.PanicLevel)

	//Initalize Viper
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cgc.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose logging")

	cobra.OnInitialize(initConfig)

	//Add all the endpoints to root command
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(actionCmd)
	rootCmd.AddCommand(apps.Cmd)
	rootCmd.AddCommand(billingCmd)
	rootCmd.AddCommand(filesCmd)
	rootCmd.AddCommand(markersCmd)
	rootCmd.AddCommand(projectsCmd)
	rootCmd.AddCommand(rate.Cmd)
	rootCmd.AddCommand(storageCmd)
	rootCmd.AddCommand(tasksCmd)
	rootCmd.AddCommand(uploadCmd)
	rootCmd.AddCommand(userCmd)
	rootCmd.AddCommand(usersCmd)

	viper.SetConfigType("yaml")
	viper.SetDefault("rooturl", "https://cgc-api.sbgenomics.com/v2/")
}

var rootCmd = &cobra.Command{
	Use:   "cgc",
	Short: "CLI for the Cancer Genomics Cloud",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
	}
}

func initConfig() {
	if verbose {
		log.SetLevel(log.DebugLevel)
	}

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.AddConfigPath(".")
		viper.SetConfigName(".cgc")
	}

	viper.SetEnvPrefix("cgc")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Errorln("Failed to parse config file")
	}
}
