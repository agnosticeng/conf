package conf

import "github.com/mitchellh/mapstructure"

type OptionsBuilderFunc func(*Options) *Options

type Options struct {
	ConfigFilePath    string
	EnvPrefix         string
	MapstructureHooks []mapstructure.DecodeHookFunc
}

func DefaultOptions() *Options {
	return &Options{
		ConfigFilePath: "",
		EnvPrefix:      "CONF",
	}
}

func WithConfigFilePath(path string) OptionsBuilderFunc {
	return func(opts *Options) *Options {
		opts.ConfigFilePath = path
		return opts
	}
}

func WithEnvPrefix(prefix string) OptionsBuilderFunc {
	return func(opts *Options) *Options {
		opts.EnvPrefix = prefix
		return opts
	}
}

func WithMapstructureHooks(hooks ...mapstructure.DecodeHookFunc) OptionsBuilderFunc {
	return func(opts *Options) *Options {
		opts.MapstructureHooks = hooks
		return opts
	}
}
