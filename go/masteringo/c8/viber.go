package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	fflag "github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

func main() {
	viper.BindEnv("GOMAXPROCS")
	val := viper.Get("GOMAXPROCS")
	fmt.Println("GOMAXPROCS:", val)
	viper.Set("GOMAXPROCS", 10)
	val = viper.Get("GOMAXPROCS")
	fmt.Println("GOMAXPROCS:", val)

	i := flag.Int("i", 100, "i param")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	conf := fflag.String("conf", "myConf", "setting the config file")
	fflag.Parse()
	fmt.Println("conf yaml:", *conf)

	*i--
	fmt.Println(*i)

	viper.BindPFlags(pflag.CommandLine)
	*i = viper.GetInt("i")
	fmt.Println(*i)

	unmarshaledData := struct {
		Name  string            `json:"name"`
		Age   int               `json:"age"`
		Job   string            `json:"job"`
		Phone string            `json:"phone"`
		Data  map[string]string `json:"data"`
		Arr   []string          `json:"arr"`
	}{"ali", 19, "developer", "09155093114", map[string]string{"key1": "val1", "key2": "val2", "key3": "val3", "key4": "val4"}, []string{"item1", "item2", "item3"}}

	data, err := json.Marshal(unmarshaledData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
	os.WriteFile("v.json", data, fs.ModePerm)
	// os.Remove("v.json")

	viper.SetConfigType("json")
	viper.SetConfigFile("v.json")
	fmt.Println("conf:", viper.ConfigFileUsed())
	viper.ReadInConfig()

	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		fmt.Println("change operation:", in.Op)
		fmt.Println("string:", in.String())
	})
	viper.WatchConfig()

	if viper.IsSet("name") {
		fmt.Println("json.name:", viper.Get("name"))
	}
	if viper.IsSet("age") {
		fmt.Println("json.age:", viper.Get("age"))
	}
	if viper.IsSet("job") {
		fmt.Println("json.job:", viper.Get("job"))
	}
	if viper.IsSet("phone") {
		fmt.Println("json.age:", viper.Get("phone"))
	}
	if viper.IsSet("data.key2") {

		fmt.Println("json.data.key2:", viper.Get("data.key2"))
	}

	data, _ = yaml.Marshal(unmarshaledData)
	fmt.Println(string(data))
	err = os.WriteFile("/tmp/"+*conf, data, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	viper.SetConfigType("yaml")
	if _, err := os.Stat(*conf); err == nil {
		fmt.Println("Using user specified configuration file!")
		viper.SetConfigFile(*conf)
	} else {
		viper.SetConfigName(*conf)
		viper.AddConfigPath("/tmp")
		viper.AddConfigPath("$HOME")
		viper.AddConfigPath("/etc")
		viper.AddConfigPath(".")
	}
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("conf:", viper.ConfigFileUsed())

	if viper.IsSet("arr") {
		fmt.Println("yaml.arr:", viper.Get("arr"))
	}
	select {}
}
