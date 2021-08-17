package conf

import (
	"errors"
	"fmt"
	"path"

	"github.com/agnosticeng/conf/providers/env"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
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

func Load(i interface{}, opts ...string) error {
	var prefix, path string

	if len(opts) > 0 {
		prefix = opts[0]
	}

	if len(opts) == 2 {
		path = opts[1]
	}

	if len(opts) > 2 {
		return errors.New("invalid amount of options")
	}

	k := koanf.New(".")

	if len(path) > 0 {
		parser := lookupParser(path)

		if parser == nil {
			return fmt.Errorf("cannot find a parser for %s", path)
		}

		if err := k.Load(file.Provider(path), parser); err != nil {
			return err
		}
	}

	if err := k.Load(env.NewEnvProvider(prefix), nil); err != nil {
		return err
	}

	if err := k.Unmarshal("", &i); err != nil {
		return err
	}

	return nil
}
