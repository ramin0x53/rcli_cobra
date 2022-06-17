package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type ServerInfo struct {
	Ip       string
	Password string
	Db       string
}

func GetConfigFilePath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	fileName := homeDir + "/.configrcli.json"
	return fileName
}

//Save server ip and password in a config file
func WriteConfig(ipHost string, pass string, db string) {
	data := ServerInfo{ipHost, pass, db}
	dataJson, err := json.Marshal(data)

	if err != nil {
		log.Println(err)
	}

	configPath := GetConfigFilePath()

	file, err := os.Create(configPath)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()

	err = ioutil.WriteFile(configPath, dataJson, 0644)
	if err != nil {
		log.Println(err)
	}
}

//Read configs from the config file
func ReadConfig() ServerInfo {

	if !CheckConfig() {
		return ServerInfo{"", "", "0"}
	}

	configPath := GetConfigFilePath()

	jsonData, err := ioutil.ReadFile(configPath)

	if err != nil {
		log.Println(err)
	}

	var data ServerInfo

	err = json.Unmarshal(jsonData, &data)

	if err != nil {
		log.Println(err)
	}

	return data
}

//Check if config file exists
func CheckConfig() bool {
	configPath := GetConfigFilePath()

	if _, err := os.Stat(configPath); err != nil {
		return false
	}
	return true
}

func CheckConfExit() {
	exist := CheckConfig()

	if !exist {
		fmt.Println("First, enter the server IP and Password with the following command:\n     rcli config")
		os.Exit(0)
	}
}
