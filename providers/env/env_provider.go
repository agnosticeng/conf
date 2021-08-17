package env

import (
	"errors"
	"os"
	"sort"
	"strings"

	"github.com/agnosticeng/dynamap"
	"github.com/iancoleman/strcase"
)

type EnvProvider struct {
	prefix string
}

func NewEnvProvider(prefix string) *EnvProvider {
	return &EnvProvider{
		prefix: prefix,
	}
}

func (this *EnvProvider) ReadBytes() ([]byte, error) {
	return nil, errors.New("env provider does not support this method")
}

func (this *EnvProvider) Read() (map[string]interface{}, error) {
	var (
		res  interface{}
		err  error
		envs = make([]string, len(os.Environ()))
	)

	copy(envs, os.Environ())
	sort.Strings(envs)

	for _, kv := range envs {
		segments := strings.SplitN(kv, "=", 2)
		path := strings.Split(segments[0], "__")

		if len(this.prefix) > 0 {
			if path[0] != this.prefix {
				continue
			} else {
				path = path[1:]
			}
		}

		for i := 0; i < len(path); i++ {
			path[i] = strings.ToLower(path[i])
			path[i] = strcase.ToCamel(path[i])
		}

		res, err = dynamap.SSet(res, segments[1], path...)

		if err != nil {
			return nil, err
		}
	}

	return res.(map[string]interface{}), nil
}

func (this *EnvProvider) Watch(cb func(event interface{}, err error)) error {
	return errors.New("env provider does not support this method")
}
