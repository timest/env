package env

import (
	"testing"
	"os"
	"time"
	"github.com/stretchr/testify/assert"
)
type config struct {
	App     string
	Port    int      `default:"8000"`
	IsDebug  bool     `env:"DEBUG"`
	Hosts   []string `slice_sep:","`
	Timeout time.Duration

	Redis struct {
		Version string `sep:""` // no sep between `CONFIG` and `REDIS`
		Host    string
		Port    int
	}

	MySQL struct {
		Version string `default:"5.7"`
		Host    string
		Port    int
	}
}
func TestGeneralEnv(t *testing.T) {
	os.Setenv("CONFIG_APP", "ENV APP")
	// os.Setenv("CONFIG_PORT", "default") // default value
	os.Setenv("CONFIG_DEBUG", "1")
	os.Setenv("CONFIG_HOSTS", "192.168.0.1,127.0.0.1")
	os.Setenv("CONFIG_TIMEOUT", "5s")

	os.Setenv("CONFIG_REDISVERSION", "3.2")
	os.Setenv("CONFIG_REDIS_HOST", "rdb")
	os.Setenv("CONFIG_REDIS_PORT", "6379")

	os.Setenv("CONFIG_MYSQL_HOST", "mysqldb")
	os.Setenv("CONFIG_MYSQL_PORT", "3306")

	defer os.Clearenv()

	cfg := new(config)
	err := Fill(cfg)
	if err != nil {
		t.Error(err)
	}

	assert := assert.New(t)

	assert.Equal(cfg.App, "ENV APP")
	assert.Equal(cfg.Port, 8000)
	assert.Equal(cfg.IsDebug, true)
	assert.Equal(cfg.Hosts, []string{"192.168.0.1", "127.0.0.1"})
	assert.Equal(cfg.Timeout, 5*time.Second)
	assert.Equal(cfg.Redis.Version, "3.2")
	assert.Equal(cfg.Redis.Host, "rdb")
	assert.Equal(cfg.MySQL.Version, "5.7")
	assert.Equal(cfg.MySQL.Host, "mysqldb")
	assert.Equal(cfg.MySQL.Port, 3306)

}



func TestNoPrefixEnv(t *testing.T) {
	os.Setenv("APP", "ENV_APP")
	// os.Setenv("PORT", "default") // default value
	os.Setenv("DEBUG", "1")
	os.Setenv("HOSTS", "192.168.1.1,127.0.0.1")
	os.Setenv("TIMEOUT", "5s")

	os.Setenv("REDISVERSION", "3.2")
	os.Setenv("REDIS_HOST", "rdb")
	os.Setenv("REDIS_PORT", "6379")

	os.Setenv("MYSQL_HOST", "mysqldb")
	os.Setenv("MYSQL_PORT", "3306")

	defer os.Clearenv()

	cfg := new(config)
	IgnorePrefix()
	err := Fill(cfg)
	if err != nil {
		t.Error(err)
	}

	assert := assert.New(t)

	assert.Equal(cfg.App, "ENV_APP")
	assert.Equal(cfg.Port, 8000)
	assert.Equal(cfg.IsDebug, true)
	assert.Equal(cfg.Hosts, []string{"192.168.1.1", "127.0.0.1"})
	assert.Equal(cfg.Timeout, 5*time.Second)
	assert.Equal(cfg.Redis.Version, "3.2")
	assert.Equal(cfg.Redis.Host, "rdb")
	assert.Equal(cfg.MySQL.Version, "5.7")
	assert.Equal(cfg.MySQL.Host, "mysqldb")
	assert.Equal(cfg.MySQL.Port, 3306)

}