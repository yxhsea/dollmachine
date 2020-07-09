package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"dollmachine/liveserver/ff_setup"
	"os"
	"path/filepath"
)

var cfgFile string
var Verbose bool

func main() {
	var RootCmd = &cobra.Command{
		Use:   "LiverServer",
		Short: "LiverServer",
		Long:  "LiverServer",
		Run: func(cmd *cobra.Command, args []string) {

			fmt.Println("This is LiveServer...")

			loggerMap := viper.GetStringMap("logger")
			mailMap := viper.GetStringMap("mail")
			restMap := viper.GetStringMap("rest")
			liveMap := viper.GetStringMap("live")

			var err error
			err = ff_setup.SetupLogger(loggerMap, mailMap)
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			err = ff_setup.SetupLive(liveMap["read_buffer"].(int64), liveMap["read_buffer"].(int64))
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			err = ff_setup.SetupRouter(restMap["host"].(string), restMap["mode"].(string))
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		},
	}

	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file")
	RootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func initConfig() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}else{
		viper.SetConfigName("pro_live")
		viper.AddConfigPath("./")
		viper.AddConfigPath("./ff_app/ff_live")
		viper.AddConfigPath(dir)
		viper.AutomaticEnv()
	}

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		fmt.Println(err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
}
