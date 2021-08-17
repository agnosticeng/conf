package main

import (
	"encoding/json"
	"fmt"
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
	Impl  string
	Items []subConfigItem
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
				{
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
	os.Setenv("APP__SUB__ITEMS__0__PORT", "443")
	os.Setenv("APP__SUB__ITEMS__1__HOST", "loco.local")

	cfg := defaultConfig()

	if err := conf.Load(&cfg, "examples/basic/conf.yaml", "APP"); err != nil {
		panic(err)
	}

	fmt.Println(pp(cfg))
}
