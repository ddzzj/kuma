package config

import (
	"encoding/json"
	"reflect"
)

func ConfigForDisplay(cfg Config) (Config, error) {
	// copy config so we don't override values, because nested structs in config are pointers
	newCfg, err := copyConfig(cfg)
	if err != nil {
		return nil, err
	}
	newCfg.Sanitize()
	return newCfg, nil
}

func copyConfig(cfg Config) (Config, error) {
	cfgBytes, err := json.Marshal(cfg)
	if err != nil {
		return nil, err
	}

	newCfg := reflect.New(reflect.TypeOf(cfg).Elem()).Interface().(Config)
	if err := json.Unmarshal(cfgBytes, newCfg); err != nil {
		return nil, err
	}
	return newCfg, nil
}
