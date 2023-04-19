package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Business struct {
	DataVolumeSize int64
	OsVolumeSnapid int64
}

type AppConfig struct {
	AppName         string
	HttpPort        int
	RunMode         string
	AutoReader      bool
	CopyRequestBody bool
	EnbaleDocs      bool
}

var (
	App AppConfig
)

func LoadIniFile() (err error) {
	viper.SetConfigName("app")
	viper.AddConfigPath("./conf/")
	viper.SetConfigType("ini")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("viper ReadInConfig err: ", err)
		return err
	}

	App.AppName = viper.GetString("default.appname")
	App.HttpPort = viper.GetInt("default.httpport1")

	fmt.Printf("app : %+v\n", App)
	return
}

func LoadYaml() (err error) {
	viper.SetConfigName("app2.yaml")
	viper.AddConfigPath("./conf/")
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("viper ReadInConfig err: ", err)
		return err
	}

	Name := viper.GetString("name")
	Jacket := viper.GetString("clothing.jacket")
	Hobbies := viper.GetStringSlice("hobbies")
	fmt.Printf("name : %+v jacket : %+v, hobbies: %+v\n", Name, Jacket, Hobbies)
	return
}

func main() {
	if err := LoadYaml(); err != nil {
		fmt.Printf("err :%+v", err)
	}
}
