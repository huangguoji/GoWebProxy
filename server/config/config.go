package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type Config struct {
	WebPort string	`yaml:"web_port"`
	ServerPort string	`yaml:"proxy_port"`
	Proxies []ProxyConfig	`yaml:"proxies"`
}

type ProxyConfig struct {
	Name string	`yaml:"name"`
	Enable bool	`yaml:"enable"`
	Domains []string	`yaml:"domains"`
	Rules	[]Rule	`yaml:"rules"`
}

type Rule struct {
	Pattern string	`yaml:"pattern"`
	Servers []ProxyServer	`yaml:"servers"`
	Local string	`yaml:"local"`
}

type ProxyServer struct {
	Address string	`yaml:"address"`
	Enable bool	`yaml:"enable"`
}

func (c *Config) Load(filePath string) (*Config,error){
	_,err := os.Stat(filePath)
	if err != nil {
		return nil,fmt.Errorf("%s not exists",filePath)
	}
	b,err:= ioutil.ReadFile(filePath)
	if err !=nil {
		return nil,fmt.Errorf("read file error : %v",err)
	}
	err =yaml.Unmarshal(b,c)
	if err !=nil {
		return nil,fmt.Errorf("parse file error : %v",err)
	}
	return c,nil
}

func (c *Config) Save(filePath string) error {
	b,err:= yaml.Marshal(c)
	if err!=nil {
		return fmt.Errorf("yaml Marshal faild : %v",err)
	}
	err = ioutil.WriteFile(filePath,b,0666)
	if err!=nil {
		return fmt.Errorf("write yaml faild : %v",err)
	}
	return err
}