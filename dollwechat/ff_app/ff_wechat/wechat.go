package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"dollmachine/dollwechat/ff_setup"
	"os"
	"path/filepath"
)

var cfgFile string
var Verbose bool

func main() {

	var RootCmd = &cobra.Command{
		Use:   "Wechat",
		Short: "Wechat",
		Long:  "Wechat",
		Run: func(cmd *cobra.Command, args []string) {

			databaseMap := viper.GetStringMap("database")
			cacheMap := viper.GetStringMap("cache")
			loggerMap := viper.GetStringMap("logger")
			wechatMap := viper.GetStringMap("wechat")
			restMap := viper.GetStringMap("rest")

			var err error
			err = ff_setup.SetupLogger(loggerMap["file_path"].(string))
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			err = ff_setup.SetupMysql(databaseMap["user"].(string), databaseMap["password"].(string), databaseMap["host"].(string), databaseMap["port"].(string), databaseMap["dbname"].(string), databaseMap["charset"].(string), int(databaseMap["poolnum"].(int64)))
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			err = ff_setup.SetupRedis(cacheMap["host"].(string), cacheMap["auth"].(string), int(cacheMap["poolnum"].(int64)))
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			err = ff_setup.SetupWxMpClient(wechatMap["wx_mp_appid"].(string), wechatMap["wx_mp_appsecret"].(string),wechatMap["wx_ori_id"].(string),wechatMap["wx_token"].(string),wechatMap["wx_encoded_aes_key"].(string))
			if err != nil {
				fmt.Println(err.Error())
			}

			err = ff_setup.SetupServer(restMap["host"].(string), restMap["mode"].(string))
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
	} else {
		viper.SetConfigName("pro_wechat")
		viper.AddConfigPath("./")
		viper.AddConfigPath("./ff_app/ff_wechat")
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
