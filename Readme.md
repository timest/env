# env
This lib is used for dealing with environment variables in Golang .

```
CONFIG_APP=ENVAPP
CONFIG_DEBUG=1
CONFIG_HOSTS=192.168.0.1,127.0.0.1
CONFIG_TIMEOUT=5s

CONFIG_REDISVERSION=3.2
CONFIG_REDIS_HOST=rdb
CONFIG_REDIS_PORT=6379

CONFIG_MYSQL_HOST=mysqldb
CONFIG_MYSQL_PORT=3306
```


```golang

import (
	"fmt"
	"time"
	"os"
	"github.com/timest/env"
)

type config struct {
	App     string
	Port    int      `default:"8000"`
	IsDebug bool     `env:"DEBUG"`
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

func main() {
	cfg := new(config)
	err := env.Fill(cfg)
	if err != nil {
		panic(err)
	}
	fmt.Println("Home:", cfg.App)
	fmt.Println("Port:", cfg.Port)
	fmt.Println("IsDebug:", cfg.IsDebug)
	fmt.Println("Hosts:", cfg.Hosts, len(cfg.Hosts))
	fmt.Println("Duration:", cfg.Timeout)
	fmt.Println("Redis_Version:", cfg.Redis.Version)
	fmt.Println("Redis_Host:", cfg.Redis.Host)
	fmt.Println("Redis_Port:", cfg.Redis.Port)
	fmt.Println("MySQL_Version:", cfg.MySQL.Version)
	fmt.Println("MySQL_Name:", cfg.MySQL.Host)
	fmt.Println("MySQL_port:", cfg.MySQL.Port)
}

// output:
// Home: ENV APP
// Port: 8000
// IsDebug: true
// Hosts: [192.168.0.1 127.0.0.1] 2
// Duration: 5s
// Redis_Version: 3.2
// Redis_Host: rdb
// Redis_Port: 6379
// MySQL_Version: 5.7
// MySQL_Name: mysqldb
// MySQL_port: 3306
```


## or

```
APP=ENVAPP
DEBUG=1
HOSTS=192.168.0.1,127.0.0.1
TIMEOUT=5s

REDISVERSION=3.2
REDIS_HOST=rdb
REDIS_PORT=6379

MYSQL_HOST=mysqldb
MYSQL_PORT=3306
```


```
func main() {
    cfg := new(config)
    env.IgnorePrefix()
    err := env.Fill(cfg)
    ...
}
```

## tag supported
* default
* env
* require
* slice_sep
* sep
