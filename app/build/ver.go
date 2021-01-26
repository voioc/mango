/*
 * @Author: jianxuesong
 * @Date: 2019-06-19 09:37:45
 * @LastEditTime: 2020-03-31 10:18:45
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: /Lemon/app/version/version.go
 */
package build

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var (
	// Version should be updated by hand at each release
	RunEnv      string
	ProjectPath string
	// ConfigFile  string
	Version   string
	GitCommit string
	BuildTime string
	GoVersion string
)

var configFile *string

func init() {
	if RunEnv == "" {
		RunEnv = "debug"
	}

	path, _ := filepath.Abs(filepath.Dir(""))
	// config := path[0:strings.LastIndex(path, "/")] + "/config/config_debug.json"
	config := path + "/config/config_debug.json"

	versionFlag := flag.Bool("v", false, "print the version")
	configFile = flag.String("c", config, "配置文件路径")
	flag.Parse()

	if *versionFlag {
		fmt.Printf("App Version: %s \n", Version)
		fmt.Printf("Git Commit: %s \n", GitCommit)
		fmt.Printf("Build Time: %s \n", BuildTime)
		fmt.Printf("Go Version: %s \n", GoVersion)
		os.Exit(0)
	}

	// viper.SetConfigFile(*ConfigFile)
	// err := viper.ReadInConfig() // Find and read the config file
	// if err != nil {             // Handle errors reading the config file
	// 	panic(fmt.Errorf("Fatal error config file: %s \n", err))
	// }

	// if RunEnv != viper.GetString("configName") {
	// 	panic(fmt.Errorf("Fatal error RunEnv: %s \n", RunEnv))
	// }
}

func GetConfigFile() *string {
	return configFile
}

func SetConfigFile(file string) {
	configFile = &file
}
