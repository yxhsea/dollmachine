package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"dollmachine/dollbarrage/ff_setup"
)

var cfgFile string
var Verbose bool

func main() {

	var RootCmd = &cobra.Command{
		Use:   "DollBarrage Server",
		Short: "DollBarrage Server",
		Long:  "DollBarrage Server",
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

			//WebSocket配置
			WebSocket := viper.GetStringMap("web_socket")
			readBuffer, _ := WebSocket["read_buffer"].(int64)
			writeBuffer, _ := WebSocket["write_buffer"].(int64)
			ff_setup.SetupWebSocket(readBuffer, writeBuffer)

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
		viper.SetConfigName("barrage")
		viper.AddConfigPath("./")
		viper.AddConfigPath("./ff_app/ff_barrage")
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
