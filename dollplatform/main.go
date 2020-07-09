package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	_ "dollmachine/dollplatform/docs"
	"dollmachine/dollplatform/ff_setup"
	"os"
	"path/filepath"
)

var cfgFile string
var Verbose bool

// @title 娃娃机运营平台 Api文档
// @version 1.0
// @description 娃娃机运营平台 Api文档
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 192.168.0.167:9556
// @BasePath /dev/plf/v1
func main() {

	var RootCmd = &cobra.Command{
		Use:   "DollPlatform Server",
		Short: "DollPlatform Server",
		Long:  "DollPlatform Server",
		Run: func(cmd *cobra.Command, args []string) {
			//日志信息
			loggerMap := viper.GetStringMap("logger")
			filePath, _ := loggerMap["file_path"].(string)
			ff_setup.SetupLogger(filePath)

			//mysql配置
			databaseMap := viper.GetStringMap("mysql")
			hostDb, _ := databaseMap["host"].(string)
			port, _ := databaseMap["port"].(string)
			user, _ := databaseMap["user"].(string)
			password, _ := databaseMap["password"].(string)
			dbname, _ := databaseMap["dbname"].(string)
			charset, _ := databaseMap["charset"].(string)
			poolnum, _ := databaseMap["pollnum"].(int)
			ff_setup.SetupMysql(hostDb, port, user, password, dbname, charset, poolnum)

			//redis配置
			redisMap := viper.GetStringMap("redis")
			hostRedis, _ := redisMap["host"].(string)
			auth, _ := redisMap["auth"].(string)
			ff_setup.SetupRedis(hostRedis, auth)

			//rpcServer
			rpcMap := viper.GetStringMap("rpc")
			hostRpc, _ := rpcMap["host"].(string)
			ff_setup.SetupRpcServer(hostRpc)

			//http服务配置
			httpMap := viper.GetStringMap("http")
			host, _ := httpMap["host"].(string)
			ff_setup.SetupServer(host)
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
		viper.SetConfigName("dev_conf")
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
