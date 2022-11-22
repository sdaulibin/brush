package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var GlobalConfig *Config

type Config struct {
	DebugMode         bool           `yaml:"DebugMode"`
	NeedPublishConfig bool           `yaml:"NeedPublishConfig"`
	ServerPort        int            `yaml:"ServerPort"`
	WriteDB           DBConfig       `yaml:"WriteDB"`
	ReadDB            DBConfig       `yaml:"ReadDB"`
	JwtToken          JwtTokenConfig `yaml:"JwtToken"`
}

type DBConfig struct {
	Name         string `yaml:"Name"`
	Host         string `yaml:"Host"`
	Port         string `yaml:"Port"`
	User         string `yaml:"User"`
	Password     string `yaml:"Password"`
	Timeout      string `yaml:"Timeout"`
	ReadTimeout  string `yaml:"ReadTimeout"`
	WriteTimeout string `yaml:"WriteTimeout"`
	MaxIdleConns int    `yaml:"MaxIdleConns"`
	MaxOpenConns int    `yaml:"MaxOpenConns"`
}

type JwtTokenConfig struct {
	JwtTokenSignKey        string `yaml:"JwtTokenSignKey"`
	JwtTokenCreatedExpires int64  `yaml:"JwtTokenCreatedExpires"`
	JwtTokenRefreshExpires int64  `yaml:"JwtTokenRefreshExpires"`
	BindContextKeyName     string `yaml:"BindContextKeyName"`
}

func MustInit() {
	GlobalConfig = loadConfig()
}

func loadConfig() *Config {
	filepath := composeConfigFileName("/Users/binginx/gocode/src/brush/conf/config.yml", os.Getenv("SpecifiedConfig"))
	log.Printf("config filepath:%s", filepath)

	f, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	var config Config
	if err = yaml.Unmarshal(f, &config); err != nil {
		panic(err)
	}
	return &config
}

func composeConfigFileName(basePath string, suffix string) string {
	var filepath = basePath

	if suffix != "" {
		filepath = strings.Join([]string{filepath, suffix}, ".")
	}

	return filepath
}
