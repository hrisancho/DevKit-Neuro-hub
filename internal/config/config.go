package config

import (
	"log"
	"os"
	"reflect"
	"time"

	"DevKit-Neuro-hub/internal/validator"
	"go.uber.org/zap/zapcore"

	"github.com/spf13/viper"
)

const (
	DefaultTimeFormat = "02.01.2006 15:04:05"
	FilterTimeFormat  = "01/02/2006"
)

var (
	DebugLevel = zapcore.DebugLevel
)

// Config используется для хранения конфигурации сервера.
type Config struct {
	// Addr   string `mapstructure:"PUBLIC_ADDR" validate:"required"`
	Addr   string `mapstructure:"PUBLIC_ADDR"`
	UseCrc bool   `mapstructure:"USE_CRC"`
}

func initDefaultConfig() (v *viper.Viper) {
	v = viper.New()

	v.SetDefault("PUBLIC_ADDR", "127.0.0.1:8889")
	v.SetDefault("PUBLIC_ADDR", "127.0.0.1:8889")
	return
}

func loadConfigFile(v *viper.Viper, path string) (config Config, err error) {
	v.AddConfigPath(path)
	v.SetConfigName("main")
	v.SetConfigType("env")

	v.AutomaticEnv()

	err = v.ReadInConfig()
	if err != nil {
		return
	}

	configReflectType := reflect.ValueOf(&config).Elem()
	configFieldsCount := configReflectType.NumField()

	err = v.Unmarshal(&config)
	if err != nil {
		return
	}
	time.Now().Unix()

	for fieldInd := 0; fieldInd < configFieldsCount; fieldInd++ {
		configField := configReflectType.Field(fieldInd)

		if configField.Kind() != reflect.Struct {
			continue
		}

		err = v.Unmarshal(configField.Addr().Interface())
		if err != nil {
			return
		}
	}

	return
}

func loadConfigEnv(v *viper.Viper) (config Config, err error) {
	envNameList := envNameListByConfig(reflect.TypeOf(config))
	for _, envName := range envNameList {
		err = v.BindEnv(envName, envName)
		if err != nil {
			return
		}
	}

	err = v.Unmarshal(&config)
	return
}

func envNameListByConfig(configType reflect.Type) (envNameList []string) {
	configFieldsCount := configType.NumField()
	envNameList = make([]string, 0, configFieldsCount)

	for fieldInd := 0; fieldInd < configFieldsCount; fieldInd++ {
		configField := configType.Field(fieldInd)

		if configField.Type.Kind() == reflect.Struct {
			envNameList = append(envNameList, envNameListByConfig(configField.Type)...)
		}

		envNameList = append(envNameList, configField.Tag.Get("mapstructure"))
	}
	return
}

func LoadConfig(appValidator *validator.AppValidator) (config Config, err error) {
	v := initDefaultConfig()

	if _, err = os.Stat("./main.env"); err == nil {
		config, err = loadConfigFile(v, "./")
	} else {
		log.Println("Loading config from env...")
		config, err = loadConfigEnv(v)
	}

	if err = appValidator.Struct(&config); err != nil {
		err = appValidator.ErrorTranslated(err)
		return
	}

	return config, err
}
