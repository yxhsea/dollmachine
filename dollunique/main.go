package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"dollmachine/dollunique/ff_setup"
	"os"
	"path/filepath"
)

var cfgFile string
var Verbose bool

func main() {

	var RootCmd = &cobra.Command{
		Use:   "DollUnique Rpc Server",
		Short: "DollUnique Rpc Server",
		Long:  "DollUnique Rpc Server",
		Run: func(cmd *cobra.Command, args []string) {
			//日志信息
			loggerMap := viper.GetStringMap("logger")
			filePath, _ := loggerMap["file_path"].(string)
			ff_setup.SetupLogger(filePath)

			//redis配置
			redisMap := viper.GetStringMap("redis")
			hostRedis, _ := redisMap["host"].(string)
			auth, _ := redisMap["auth"].(string)
			ff_setup.SetupRedis(hostRedis, auth)

			//rpc服务配置
			rpcMap := viper.GetStringMap("rpc")
			hostRpc, _ := rpcMap["host"].(string)
			ff_setup.SetupServer(hostRpc)
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
	} else {
		viper.SetConfigName("conf")
		viper.AddConfigPath("./")
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
