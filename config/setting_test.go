package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"reflect"
	"testing"
)

func TestConfigInit(t *testing.T) {
	err := Init(".", "config", "yaml")
	if err != nil {
		t.Error(err.Error())
	}
	pwd, _ := os.Getwd()
	fmt.Println(pwd)
	fmt.Println(viper.AllSettings())
	if Cfg.Name == "" || Cfg.Level == "" {
		t.Error("配置文件，反序列化失败...")
	} else {
		v := reflect.ValueOf(Cfg).Elem()
		t := v.Type()
		println(t)
		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("字段名: %s, 字段类型: %s, 字段值: %v\n",
				t.Field(i).Name,
				t.Field(i).Type,
				v.Field(i).Interface())
		}
	}
}
