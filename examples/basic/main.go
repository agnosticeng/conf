package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/agnosticeng/conf"
)

type config struct {
	Amount  int
	Tag     string
	Name    string
	Timeout time.Duration
	Sub     subConfig
}

type subConfig struct {
	Impl   string
	Target *url.URL
	Items  []subConfigItem
}

type subConfigItem struct {
	Host              string
	Port              int
	ConnectionTimeout time.Duration
	RequestTimeout    time.Duration
}

func defaultConfig() config {
	return config{
		Amount:  1,
		Tag:     "lolo",
		Timeout: time.Second,
		Sub: subConfig{
			Impl: "coco",
			Items: []subConfigItem{
				subConfigItem{
					Host:              "test.local",
					Port:              80,
					ConnectionTimeout: time.Second,
					RequestTimeout:    time.Second,
				},
			},
		},
	}
}

func pp(i interface{}) string {
	b, _ := json.MarshalIndent(i, "", "    ")
	return string(b)
}

func main() {
	os.Setenv("APP__NAME", "KOANF")
	os.Setenv("APP__SUB__TARGET", "s3://my-bucket/coco.json")
	os.Setenv("APP__SUB__ITEMS__0__PORT", "443")
	os.Setenv("APP__SUB__ITEMS__1__HOST", "loco.local")
	os.Setenv("APP__SUB__ITEMS__1__REQUEST_TIMEOUT", "17s")

	cfg := defaultConfig()

	if err := conf.Load(&cfg, conf.WithConfigFilePath("examples/basic/conf.yaml"), conf.WithEnvPrefix("APP")); err != nil {
		panic(err)
	}

	fmt.Println(pp(cfg))
}
