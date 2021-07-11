package ngenx

import "strings"

type Config struct {
	Servers []Server `json:"servers" yaml:"servers"`
}

type Server struct {
	Listen      int       `json:"listen" yaml:"listen"`
	ServerNames []string  `json:"serverNames" yaml:"serverNames"`
	Proxy       ProxyInfo `json:"proxy" yaml:"proxy"`
}

type ProxyInfo struct {
	Url     string            `json:"url" yaml:"url"`
	Headers map[string]string `json:"headers" yaml:"headers"`
}

func (cfg *Config) Prepare() {
	for idx, _ := range cfg.Servers {
		proxyUrl := cfg.Servers[idx].Proxy.Url
		if !strings.HasSuffix(proxyUrl, "/") {
			cfg.Servers[idx].Proxy.Url += "/"
		}
	}
}
