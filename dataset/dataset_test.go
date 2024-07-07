package dataset

import (
	"bluebell/config"
	"bluebell/dataset/mysql"
	"bluebell/dataset/redis"
	"testing"
)

func TestInitMysql(t *testing.T) {
	err := config.Init("../config", "config", "yaml")
	if err != nil {
		t.Error("config init fail", err)
	}
	err = mysql.InitMysql()
	if err != nil {
		t.Error("init mysql fail:", err)
	}
}

func TestInitRedis(t *testing.T) {
	err := config.Init("../config", "config", "yaml")
	if err != nil {
		t.Error("config init fail", err)
	}
	err = redis.InitRedis()
	if err != nil {
		t.Error("init mysql fail:", err)
	}
}
