package conf

type OptionsBuilderFunc func(*Options) *Options

type Options struct {
	ConfigFilePath string
	EnvPrefix      string
}

func DefaultOptions() *Options {
	return &Options{
		ConfigFilePath: "config.yaml",
		EnvPrefix:      "CONF",
	}
}

func WithConfigFilePath(path string) OptionsBuilderFunc {
	return func(opts *Options) *Options {
		opts.ConfigFilePath = path
		return opts
	}
}

func WithEnPrefix(prefix string) OptionsBuilderFunc {
	return func(opts *Options) *Options {
		opts.EnvPrefix = prefix
		return opts
	}
}
