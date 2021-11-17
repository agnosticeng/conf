package conf

import (
	"fmt"
	"path"

	"github.com/agnosticeng/conf/providers/env"
	mapstructure_hooks "github.com/agnosticeng/mapstructure-hooks"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/mitchellh/mapstructure"
)

func lookupParser(filepath string) koanf.Parser {
	switch path.Ext(filepath) {
	case ".json":
		return json.Parser()
	case ".yaml", ".yml":
		return yaml.Parser()
	default:
		return nil
	}
}

func Load(i interface{}, optsBuilders ...OptionsBuilderFunc) error {
	k := koanf.New(".")

	var opts = DefaultOptions()

	for _, optsBuilder := range optsBuilders {
		opts = optsBuilder(opts)
	}

	if len(opts.ConfigFilePath) > 0 {
		parser := lookupParser(opts.ConfigFilePath)

		if parser == nil {
			return fmt.Errorf("cannot find a parser for %s", opts.ConfigFilePath)
		}

		if err := k.Load(file.Provider(opts.ConfigFilePath), parser); err != nil {
			return err
		}
	}

	if err := k.Load(env.NewEnvProvider(opts.EnvPrefix), nil); err != nil {
		return err
	}

	var hooks []mapstructure.DecodeHookFunc

	hooks = append(hooks, mapstructure.StringToTimeDurationHookFunc())
	hooks = append(hooks, mapstructure_hooks.All()...)

	if len(opts.MapstructureHooks) > 0 {
		hooks = append(hooks, opts.MapstructureHooks...)
	}

	var mdc mapstructure.DecoderConfig

	mdc.Metadata = nil
	mdc.Result = &i
	mdc.WeaklyTypedInput = true
	mdc.DecodeHook = mapstructure.ComposeDecodeHookFunc(hooks...)

	if err := k.UnmarshalWithConf("", &i, koanf.UnmarshalConf{DecoderConfig: &mdc}); err != nil {
		return err
	}

	return nil
}
