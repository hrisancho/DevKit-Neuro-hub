package config

import (
	"log"
	"os"
	"reflect"
	"time"

	"go.uber.org/zap/zapcore"
	"smart-doors-server/internal/validator"

	"github.com/spf13/viper"
)

const (
	ProtocolHTTP  = "http"
	ProtocolHTTPS = "https"
)

const (
	DefaultTimeFormat        = "02.01.2006 15:04:05"
	FilterTimeFormat         = "01/02/2006"
	DefaultCompanyUUID       = "d6b1890b-cbb3-4cfd-aba0-e99f814413db"
	DefaultDemoUserLifeTime  = time.Duration(time.Minute * 5)
	FaceRecognizerTimeout    = time.Duration(time.Second * 7)
	ControllerRequestTimeout = time.Duration(time.Second * 15)

	DefaultUserGroupGuest      uint32 = 1
	DefaultUserGroupUser       uint32 = 2
	DefaultUserGroupSecurity   uint32 = 3
	DefaultUserGroupService    uint32 = 4
	DefaultUserGroupSuperAdmin uint32 = 5
	DefaultUserGroupAdmin      uint32 = 6

	ControllerOtaFirmwareProtocol     = ProtocolHTTPS   // ProtocolHTTP, ProtocolHTTPS
	ControllerOtaFirmwareMinSizeBytes = 2 * 1024 * 1024 // Min 2 MB
	ControllerOtaFirmwareMaxSizeBytes = 4 * 1024 * 1024 // Max 4 MB
)

var (
	DebugLevel = zapcore.DebugLevel
)

type ConfigRedis struct {
	// Addr - адрес Redis сервера
	Addr     string `mapstructure:"REDIS_ADDR" validate:"required"`
	Username string `mapstructure:"REDIS_USERNAME"`
	Password string `mapstructure:"REDIS_PASSWORD"`
	// DB - номер БД Redis
	DB int `mapstructure:"REDIS_DB"`
}

type ConfigInflux struct {
	Bucket       string `mapstructure:"INFLUX_BUCKET" validate:"required"`
	Organization string `mapstructure:"INFLUX_ORGANIZATION" validate:"required"`
	Token        string `mapstructure:"INFLUX_TOKEN" validate:"required"`
	URL          string `mapstructure:"INFLUX_URL" validate:"required"`
}

// Config используется для хранения конфигурации сервера.
type Config struct {
	// DBSource - DSN БД
	DBSource string `mapstructure:"DB_SOURCE" validate:"required"`
	// ServerAddr - адрес HTTP сервера (по умолчанию: 127.0.0.1:8080)
	ServerAddr string `mapstructure:"SERVER_ADDR"`
	// SiteURL - URL адрес сайта (обязательное значение)
	SiteURL string `mapstructure:"SITE_URL" validate:"required"`
	// TimeFormat - формат времени (по умолчанию: 02.01.2006 15:04:05)
	TimeFormat string `mapstructure:"TIME_FORMAT"`
	// TokenTG - токен для telegram bot
	TokenTG string `mapstructure:"TOKEN_TG" validate:"required"`

	// SMTP* - данные для smtp
	SMPTHost     string `mapstructure:"SMTP_HOST" validate:"required"`
	SMPTPort     int    `mapstructure:"SMTP_PORT" validate:"required"`
	SMPTLogin    string `mapstructure:"SMTP_LOGIN" validate:"required"`
	SMPTPassword string `mapstructure:"SMTP_PASSWORD" validate:"required"`

	// Secret - секрет (длина 32)
	Secret string `mapstructure:"SECRET"`
	// AllowOriginsCors - CORS разрешенные источники (по умолчанию: *)
	AllowOriginsCors string `mapstructure:"ALLOW_ORIGINS_CORS"`
	// TokenLifetime - время жизни bearer (JWT) токена (по умолчанию: 168h)
	TokenLifetime time.Duration `mapstructure:"TOKEN_LIFETIME"`

	// MqttBrokerAddr - MQTT broker full addr (пример: tcp://192.168.1.108:1883)
	MqttBrokerAddr string `mapstructure:"MQTT_BROKER_ADDR" validate:"required"`
	MqttUsername   string `mapstructure:"MQTT_USERNAME"`
	MqttPassword   string `mapstructure:"MQTT_PASSWORD"`
	// MqttKeepAlive - MQTT client keep alive
	MqttKeepAlive time.Duration `mapstructure:"MQTT_KEEP_ALIVE"`
	// MqttPingTimeout - MQTT client ping timeout
	MqttPingTimeout time.Duration `mapstructure:"MQTT_PING_TIMEOUT"`
	// MqttQOS - MQTT client QoS
	MqttQOS byte `mapstructure:"MQTT_QOS"`

	// DefaultSuperAdminEmail - логин первого супер администратора
	DefaultSuperAdminEmail string `mapstructure:"DEFAULT_SUPER_ADMIN_EMAIL" validate:"required"`
	// DefaultSuperAdminPassword - пароль первого супер администратора
	DefaultSuperAdminPassword string `mapstructure:"DEFAULT_SUPER_ADMIN_PASSWORD" validate:"required"`

	// DefaultCompanyName - название первой компании
	DefaultCompanyName string `mapstructure:"DEFAULT_COMPANY_NAME" validate:"required"`

	// DefaultDemoUserLifeTime - время жизни demo user в период которого ему приходят уведомления
	DefaultDemoUserLifeTime time.Duration `mapstructure:"DEFAULT_DEMO_USER_LIFE_TIME"`

	ConfigRedis
	ConfigInflux
}

func initDefaultConfig() (v *viper.Viper) {
	v = viper.New()

	v.SetDefault("SERVER_ADDR", "0.0.0.1:8080")
	v.SetDefault("TIME_FORMAT", DefaultTimeFormat)
	v.SetDefault("SECRET", "&7JHHOA8*I5un5iOt7Kr2MpXGfGl7a#O")
	v.SetDefault("ALLOW_ORIGINS_CORS", "*")
	v.SetDefault("TOKEN_LIFETIME", time.Hour*24*7)
	v.SetDefault("DEFAULT_DEMO_USER_LIFE_TIME", DefaultDemoUserLifeTime)

	v.SetDefault("MQTT_KEEP_ALIVE", 60*time.Second)
	v.SetDefault("MQTT_PING_TIMEOUT", 10*time.Second)
	v.SetDefault("MQTT_QOS", 2)

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

	if _, err = os.Stat("./config/main.env"); err == nil {
		config, err = loadConfigFile(v, "./config")
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
