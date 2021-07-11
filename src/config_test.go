package ngenx

import "testing"

func TestConfig_Prepare(t *testing.T) {
	cfg := Config{
		Servers: []Server{
			{
				Listen: 80,
				ServerNames: []string{"subdomain.localhost"},
				Proxy: ProxyInfo{
					Url: "http://server.com",
				},
			},
		},
	}

	cfg.Prepare()
	if cfg.Servers[0].Proxy.Url != "http://server.com/" {
		t.Errorf("proxy_pass url not suffixed with a trailing '/'")
	}
}