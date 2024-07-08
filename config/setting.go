package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

var Cfg = new(AppConfig)

type AppConfig struct {
	*WebConfig    `mapstructure:"web"`
	*MyAuthConfig `mapstructure:"auth"`
	*LogConfig    `mapstructure:"log"`
	*MySQLConfig  `mapstructure:"mysql"`
	*RedisConfig  `mapstructure:"redis"`
}

type WebConfig struct {
	Name      string `mapstructure:"name"`
	Mode      string `mapstructure:"mode"`
	Version   string `mapstructure:"version"`
	Port      int    `mapstructure:"port"`
	StartTime string `mapstructure:"start_time"`
	MachineID int64  `mapstructure:"machine_id"`
}

type MyAuthConfig struct {
	JwtExpire int `mapstructure:"jwt_expire"`
}

type MySQLConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DB       string `mapstructure:"dbname"`
	MaxOpen  int    `mapstructure:"max_open"`
	MaxIdle  int    `mapstructure:"max_idle"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"dbname"`
	PoolSize int    `mapstructure:"pool_size"`
	MinIdle  int    `mapstructure:"min_idle"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

func Init(ConfigPath string, ConfigName string, ConfigType string) (err error) {
	viper.SetConfigName(ConfigName)
	viper.SetConfigType(ConfigType)
	viper.AddConfigPath(ConfigPath)
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Fatal error config file: %s\n", err)
		return err
	}
	// 把读取到的序列信息反序列化
	if err := viper.Unmarshal(Cfg); err != nil {
		fmt.Printf("Fatal error config file: %s\n", err)
		return err
	}
	// 监控配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Printf("Config file changed: %s\n", e.Name)
		if err := viper.Unmarshal(Cfg); err != nil {
			fmt.Printf("Fatal error config file: %s\n", err)
		}
	})
	return nil
}
