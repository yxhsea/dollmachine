package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"dollmachine/dollmqtt/ff_setup"
	"os"
	"path/filepath"
	"time"
)

var cfgFile string
var Verbose bool

func main() {

	var RootCmd = &cobra.Command{
		Use:   "DollMachine MqttSubscribe",
		Short: "DollMachine Mqtt Subscribe of DollMachine",
		Long:  "DollMachine Mqtt Subscribe is subscribe part of DollMachine",
		Run: func(cmd *cobra.Command, args []string) {

			databaseMap := viper.GetStringMap("database")
			cacheMap := viper.GetStringMap("cache")
			loggerMap := viper.GetStringMap("logger")
			mqttSubMap := viper.GetStringMap("mqtt_sub")

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

			err = ff_setup.SetupMqttClient(mqttSubMap["broker"].(string), mqttSubMap["user"].(string), mqttSubMap["password"].(string), mqttSubMap["client_id"].(string), mqttSubMap["store"].(string), mqttSubMap["topic"].(string), mqttSubMap["qos"].(int64))
			if err != nil {
				fmt.Println(err.Error())
				return
			}

			t1 := time.NewTimer(time.Second * 10)
			for {
				select {
				case <-t1.C:
					t1.Reset(time.Second * 10)
				}
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
		viper.SetConfigName("pro_mqtt")
		viper.AddConfigPath("./")
		viper.AddConfigPath("./ff_app/ff_mqtt")
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
