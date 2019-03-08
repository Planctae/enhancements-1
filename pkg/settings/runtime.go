package settings

import (
	"fmt"
	"os"
	"sync"

	"github.com/spf13/viper"
)

type Runtime interface {
	ReceiptsLocation() string
	Persist(toLocation string) error
}

const (
	// the directory containing release tracking subdirectories
	ReceiptsDirname   = "releases"
	Filename          = "settings"
	EnvPrefix         = "reltrackr"
	ReceiptsConfigKey = "receipts_location"
	ConfigDirName     = ".reltrackr"
)

const (
	configType = "yaml"
)

func NewRuntime(receiptsLocation string) (Runtime, error) {
	if _, statErr := os.Stat(receiptsLocation); os.IsNotExist(statErr) {
		return nil, fmt.Errorf("refusing to create runtime settings with non existent location: %s", receiptsLocation)
	}

	v := viper.New()
	v.SetConfigType(configType)
	v.SetConfigName(Filename)

	v.Set(ReceiptsConfigKey, receiptsLocation)

	r := &runtime{
		underlying: v,
		locker:     new(sync.RWMutex),
	}

	return r, nil
}

type runtime struct {
	underlying *viper.Viper
	locker     sync.Locker
}

func (r *runtime) ReceiptsLocation() string {
	r.locker.Lock()
	defer r.locker.Unlock()

	return r.underlying.GetString(ReceiptsConfigKey)
}

func (r *runtime) Persist(toLocation string) error {
	r.locker.Lock()
	defer r.locker.Unlock()

	return r.underlying.WriteConfigAs(toLocation)
}
