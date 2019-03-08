package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var (
	configPath          = os.Getenv("HOME")
	repositoryCachePath = configPath + "/.bissucket.repositoriescache.json"
	issueCachePath      = configPath + "/.bissucket.issuecache.json"
)

const (
	configFileName = ".bissucket.config"
	configFileType = "json"
)

func setConfigPath() {
	viper.SetConfigName(configFileName)
	viper.AddConfigPath(configPath)
	viper.AddConfigPath(".")
}

func CheckConfig() error {

	setConfigPath()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil

}

func GetConfigValueByKey(key string) (configValue string) {
	configValue = viper.GetString(key)
	return
}

func SetConfigKeyAndValue(key string, value string) error {

	viper.Set(key, value)

	err := writeConfigFile()
	if err != nil {
		return err
	}

	return nil
}

func CreateConfigFile(userName string, pass string) error {

	viper.Set("bitbucketUserName", userName)
	viper.Set("bitbucketPassword", pass)

	viper.Set("repositoryCachePath", repositoryCachePath)
	viper.Set("issueCachePath", issueCachePath)

	err := writeConfigFile()
	if err != nil {
		return err
	}

	return nil

}

func writeConfigFile() error {

	configJson, err := json.MarshalIndent(viper.AllSettings(), "", "    ")
	if err != nil {
		return fmt.Errorf("JsonMarshalError: %s", err)
	}

	err = ioutil.WriteFile(filepath.Join(configPath, configFileName+"."+configFileType), configJson, os.ModePerm)
	if err != nil {
		return fmt.Errorf("WriteFileError: %s", err)
	}

	return nil
}
