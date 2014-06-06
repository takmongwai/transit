/**
解析配置文件
*/
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sort"
)

//转换器数据结构
type Converter struct {
	Method map[string]string            `json:"method"`
	Params map[string]map[string]string `json:"params"`
}

//请求源配置
type Source struct {
	Converter *Converter `json:"converter"`
	Params    []string   `json:"params"`
	Path      []string   `json:"path"`
	Url       []string   `json:"url"`
}

//请求目标
type Target struct {
	Path string `json:"path"`
	Url  string `json:"url"`
}

//单一配置项
type Config struct {
	Id       string   `json:"id"`
	Extends  []string `json:"extends"`
	Priority int      `json:"priority"`
	Source   *Source  `json:"source"`
	Target   *Target  `json:"target"`
}

//监听配置
type Listen struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	Unix string `json:"unix"`
}

//配置文件
type ConfigFile struct {
	Config  []*Config `json:"config"`
	Define  []*Config `json:"define"`
	Default *Config   `json:"default"`
	Listen  *Listen   `json:"listen"`
}

//按优先级进行排序
type sortByPriority []*Config

func (v sortByPriority) Len() int           { return len(v) }
func (v sortByPriority) Swap(i, j int)      { v[i], v[j] = v[j], v[i] }
func (v sortByPriority) Less(i, j int) bool { return v[i].Priority < v[j].Priority }

//读取配置文件
func loadConfigFile(filename string) (cf *ConfigFile, err error) {
	var b []byte
	if b, err = ioutil.ReadFile(filename); err != nil {
		log.Printf("read config file error -- %v", err)
		return
	}

	cf = &ConfigFile{}
	if err = json.Unmarshal(b, cf); err != nil {
		log.Printf("parse config file error -- %v", err)
		return
	}
	sort.Sort(sortByPriority(cf.Config))

	return
}

func main() {
	cfg, err := loadConfigFile("/Users/weidewang/ownCloud/Workspace/go/transit/etc/dev_config.json")
	if err != nil {
		log.Println(err)
	}
	log.Println(cfg)
}
