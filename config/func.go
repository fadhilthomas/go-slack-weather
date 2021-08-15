package config

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
)

func Set(key string, value string) {
	base[key] = value
}

func Get(key string) string {
	r := os.Getenv(key)
	if r != "" {
		return r
	}

	log.Error().Str("file", "main").Msg(fmt.Sprintf("%v is empty", key))

	if configValue, ok := base[key]; ok {
		return configValue
	}

	return ""
}

func mergeConfig(configs ...map[string]string) map[string]string {
	result := map[string]string{}

	for _, configMap := range configs {
		for key, configValue := range configMap {
			if _, ok := result[key]; ok {
				panic(fmt.Sprintf(`duplicate config key "%s" detected`, key))
			}

			result[key] = configValue
		}
	}

	return result
}
