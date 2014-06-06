/**
解析配置文件
*/
package main

//转换器数据结构
type Converter struct {
  Method map[string]string
  Params map[string]map[string]string
}

//请求源配置
type Source struct {
  Converter *Converter
  Params    map[string]string
  Path      []string
  Url       []string
}

//请求目标
type Target struct {
  Path []string
  Url  []string
}

//单一配置项
type Config struct {
  Id       string
  Extends  []string
  Priority int
  Source   *Source
  Target   *Target
}

//监听配置
type Listen struct {
  Host string
  Port int
  Unix string
}

//配置文件
type ConfigFile struct {
  Config []*Config
  Define []*Config
  Listen *Listen
}
