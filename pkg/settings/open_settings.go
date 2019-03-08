package settings

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/spf13/viper"
)

// Open will do one of:
//
// 1. read an existing YAML configuration file into a viper instance
// 2. read in environment variables into a viper instance if no configuration
//    path is given
//
// In both cases Open will check that the expected configuration values exist or
// will return an error.
func Open(existingPath string) (Runtime, error) {
	panic("not tested")

	switch {
	case existingPath != "":
		filename := filepath.Base(existingPath)
		extension := filepath.Ext(existingPath)
		containingDir := filepath.Dir(existingPath)

		if extension == "" {
			return nil, fmt.Errorf("existing path: %s does not include `.yaml` extension", existingPath)
		}

		// TODO decide whether it makes sense to read in env variables here as well
		v := viper.New()
		v.SetConfigType(configType)
		v.SetConfigName(filename)
		v.AddConfigPath(containingDir)

		err := v.ReadInConfig()
		if err != nil {
			return nil, err
		}

		receiptsLocation := v.GetString(ReceiptsConfigKey)
		if receiptsLocation == "" {
			return nil, fmt.Errorf("found no `receipts_location` in existing settings file: %s", existingPath)
		}

		if _, statErr := os.Stat(receiptsLocation); os.IsNotExist(statErr) {
			return nil, fmt.Errorf("refusing to create settings with non existent `receipts_location`: %s", receiptsLocation)
		}

		r := &runtime{
			underlying: v,
			locker:     new(sync.RWMutex),
		}

		return r, nil

	default:
		v := viper.New()
		v.SetConfigType(configType)
		v.SetConfigName(filename)
		v.SetEnvPrefix(EnvPrefix)
		v.AutomaticEnv()

		receiptsLocation := v.GetString(ReceiptsConfigKey)
		if receiptsLocation == "" {
			return nil, fmt.Errorf("found no `receipts_location` in environment")
		}

		if _, statErr := os.Stat(receiptsLocation); os.IsNotExist(statErr) {
			return nil, fmt.Errorf("refusing to create settings with non existent `receipts_location`: %s", receiptsLocation)
		}

		r := &runtime{
			underlying: v,
			locker:     new(sync.RWMutex),
		}

		return r, nil
	}
}
