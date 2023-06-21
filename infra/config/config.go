package config

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"sync"
)

type Project struct {
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	PortPub int    `yaml:"port_pub"`
	PortPri int    `yaml:"port_pri"`
}

type MySQL struct {
	Username        string `yaml:"username"`
	Password        string `yaml:"password"`
	DbHost          string `yaml:"db_host"`
	DbPort          int    `yaml:"db_port"`
	DbName          string `yaml:"db_name"`
	Charset         string `yaml:"charset"`
	Timeout         string `yaml:"timeout"`
	TimeoutRead     string `yaml:"timeout_read"`
	TimeoutWrite    string `yaml:"timeout_write"`
	MaxOpenConns    int    `yaml:"max_open_conns"`
	MaxIdleConns    int    `yaml:"max_idle_conns"`
	ConnMaxLifetime int    `yaml:"conn_max_lifetime"`
}

type Redis struct {
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	Password   string `yaml:"password"`
	DB         int    `yaml:"db"`
	MaxRetries int    `yaml:"max_retries"`
	PoolSize   int    `yaml:"pool_size"`
}

type Config struct {
	Project Project `yaml:"project"`
	MySQL   MySQL   `yaml:"mysql"`
	Redis   Redis   `yaml:"redis"`
}

var (
	configOnce sync.Once
	configImpl *Config
)

func NewConfig() *Config {
	configOnce.Do(func() {

		// 生产环境
		configFilePath := "config.yaml"

		conf := &Config{}
		err := conf.loadConfig(configFilePath)
		if err != nil {
			log.Fatalf("load %v failed: %v", configFilePath, err)
			return
		}
	})

	return configImpl
}

// loadConfig 加载配置
func (conf *Config) loadConfig(path string) (err error) {

	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("load %v failed: %v", path, err)
	}

	err = yaml.Unmarshal(file, &configImpl)
	if err != nil {
		log.Fatalf("unmarshal yaml file failed: %v", err)
	}

	return
}
