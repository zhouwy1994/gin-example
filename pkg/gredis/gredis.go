package gredis

import (
	"errors"
	"github.com/zhouwy1994/gin-example/pkg/setting"
	"log"
	"strings"
	"time"

	"github.com/go-redis/redis"
)

var (
	clusterMode bool
	sClient     *redis.Client
	cClient     *redis.ClusterClient
)


func init() {
	sec, err := setting.Cfg.GetSection("redis")
	if err != nil {
		log.Fatal(2, "Fail to get section 'logger': %v", err)
	}

	address := sec.Key("ADDRESS").String()
	authInfo := sec.Key("AUTH_INFO").String()
	timeout := time.Duration(sec.Key("TIMEOUT").MustInt(3)) * time.Second

	err = newRedisClient(strings.Split(address, ","), authInfo, timeout)
	if err != nil {
		log.Println(err)
	}
}

func newRedisClient(address []string, authInfo string,
	timeout time.Duration) error {
	if len(address) < 1 {
		return errors.New("redis address not empty")
	}

	clusterMode = len(address) > 1
	if clusterMode {
		cClient = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:        address,
			Password:     authInfo,
			DialTimeout:  timeout,
			ReadTimeout:  timeout,
			WriteTimeout: timeout,
		})

		return nil
	}

	sClient = redis.NewClient(&redis.Options{
		Addr:         address[0],
		Password:     authInfo,
		DialTimeout:  timeout,
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
	})

	// 测试连通性
	//_, err := sClient.Ping().Result()

	return nil
}

func Set(key, value string, expire int) error {
	if clusterMode {
		if _,err := cClient.Do("set", key, value).Result();err != nil {
			return err
		}

		if expire > 0 {
			if _,err := cClient.Do("expire", key, expire).Result();err != nil {
				return err
			}
		}
	}

	if _,err := cClient.Do("set", key, value).Result();err != nil {
		return err
	}

	if expire > 0 {
		if _,err := cClient.Do("expire", key, expire).Result();err != nil {
			return err
		}
	}

	return nil
}

func Get(key string) (string, error) {
	if clusterMode {
		return cClient.Do("get", key).String()
	}

	return sClient.Do("get", key).String()
}

func Delete(key string) error {
	if clusterMode {
		_, err := cClient.Do("del", key).Result()
		return err
	}

	_, err := sClient.Do("del", key).Result()
	return err
}

func SetHash(hashKey, filed, value string, expire int) error {
	if clusterMode {
		if _,err := cClient.Do("hset", hashKey, filed, value).Result(); err != nil {
			return err
		}

		if expire > 0 {
			if _,err := cClient.Do("expire", hashKey, expire).Result();err != nil {
				return err
			}
		}
	}

	if _,err := cClient.Do("hset", hashKey, filed, value).Result(); err != nil {
		return err
	}

	if expire > 0 {
		if _,err := cClient.Do("expire", hashKey, expire).Result();err != nil {
			return err
		}
	}

	return nil
}

func GetHash(hashKey, filed string) (string, error) {
	if clusterMode {
		result, err := cClient.Do("hget", hashKey, filed).String()
		return result, err
	}

	result, err := sClient.Do("hget", hashKey, filed).String()
	return result, err
}


func DeleteHash(hashKey, filed string) error {
	if clusterMode {
		_, err := cClient.Do("hdel", hashKey, filed).Result()
		return err
	}

	_, err := sClient.Do("hdel", hashKey, filed).Result()
	return err
}
