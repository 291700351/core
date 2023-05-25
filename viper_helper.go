package core

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func NewViperHelper(file string) *ViperHelper {
	_, err := os.Stat(file)
	if nil != err {
		if !os.IsNotExist(err) {
			panic(err)
		}
	}

	ext := filepath.Ext(file)
	if len(ext) == 0 {
		panic(errors.New("unknown file extension, please specify .yaml、.json file"))
	}

	v := viper.New()
	v.SetConfigFile(file)
	_ = v.ReadInConfig()

	return &ViperHelper{
		v: v,
	}
}

type ViperHelper struct {
	v *viper.Viper
}

func (v *ViperHelper) Set(key string, value any) {
	v.v.Set(key, value)
}

func (v *ViperHelper) WriteConfig() {
	v.v.WriteConfig()
}

func (v *ViperHelper) Get(key string) any {
	return v.v.Get(key)
}
func (v *ViperHelper) GetBool(key string) bool {
	return v.v.GetBool(key)
}
func (v *ViperHelper) GetString(key string) string {
	return v.v.GetString(key)
}
func (v *ViperHelper) GetInt(key string) int {
	return v.v.GetInt(key)
}
