package cache

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"log"
	"time"
)

const ReqresKey = "reqres_"
const TimeStamp = ":time"

func (redis *RedisClient) GetByKey(key string) (string, string, error) {
	reqresKey := ReqresKey + key
	timestampKey := reqresKey + TimeStamp

	result, err := redis.Get(reqresKey).Bytes()
	timestamp, err := redis.Get(timestampKey).Result()

	if err != nil {
		return "", "", err
	}

	rData := bytes.NewReader(result)

	r, err := gzip.NewReader(rData)

	if err != nil {
		log.Print("Cannot build gzip reader: ", err)
		return "", "", err
	}

	s, err := ioutil.ReadAll(r)

	if err != nil {
		log.Print("Cannot use ioutil.ReadAll: ", err)
		r.Close()
		return "", "", err
	}

	r.Close()

	return string(s), timestamp, err
}

func (redis *RedisClient) SetValueByKey(key string, value string) error {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)

	if _, err := gz.Write([]byte(value)); err != nil {
		panic(err)
	}

	if err := gz.Flush(); err != nil {
		panic(err)
	}

	if err := gz.Close(); err != nil {
		panic(err)
	}

	reqresKey := ReqresKey + key
	timestampKey := reqresKey + TimeStamp
	t := time.Now().Unix()


	err := redis.Set(reqresKey, b.Bytes(), 0).Err()

	if err == nil {
		err = redis.Set(timestampKey, t, 0).Err()
		if err != nil {
			redis.Del(reqresKey)
		}
	}

	return err
}
