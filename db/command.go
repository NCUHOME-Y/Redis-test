package db

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

type Info struct {
	Value    string `json:"value,omitempty"`
	ExpireAt int64  `json:"expire_at,omitempty"`
}

func SET(matchString []string) {

	intoMap, err := LoadJsonTransformIntoMap()
	if err != nil {
		log.Println(err)
		return
	}
	var expireAt int
	if len(matchString) == 4 {
		expireAt, err = strconv.Atoi(matchString[3])
		if err != nil {
			log.Println(err)
			return
		}
	}
	expireAt = 999999
	info := Info{Value: matchString[2], ExpireAt: time.Now().Add(time.Minute * time.Duration(expireAt)).Unix()}

	intoMap[matchString[1]] = info
	if err = Write(intoMap); err != nil {
		log.Println(err)
		return
	}
	fmt.Println(1)
}

func SETNX(matchString []string) {
	intoMap, err := LoadJsonTransformIntoMap()
	if err != nil {
		log.Println(err)
		return
	}
	if _, ok := intoMap[matchString[1]]; ok {
		fmt.Println(0)
		return
	}
	var expireAt int
	if len(matchString) == 4 {
		expireAt, err = strconv.Atoi(matchString[3])
		if err != nil {
			log.Println(err)
			return
		}
	}
	expireAt = 9999
	info := &Info{Value: matchString[2], ExpireAt: time.Now().Add(time.Minute * time.Duration(expireAt)).Unix()}
	intoMap[matchString[1]] = info
	if err = Write(intoMap); err != nil {
		log.Println(err)
		return
	}
	fmt.Println(1)
}

func DEL(matchString []string) {
	intoMap, err := LoadJsonTransformIntoMap()
	if err != nil {
		log.Println(err)
		return
	}
	delete(intoMap, matchString[1])
	if err = Write(intoMap); err != nil {
		log.Println(err)
		return
	}
	fmt.Println("success")
}

func GET(matchString []string) {
	intoMap, err := LoadJsonTransformIntoMap()
	if err != nil {
		log.Println(err)
		return
	}
	if val, ok := intoMap[matchString[1]]; ok {
		m := val.(map[string]interface{})
		if int64(m["expire_at"].(float64)) > time.Now().Unix() {
			delete(intoMap, matchString[1])
			err = Write(intoMap)
			if err != nil {
				log.Println(err)
			}
		}
		fmt.Println(m["value"])
	}
	return
}

func LPUSH(matchString []string) {
	intoMap, err := LoadJsonTransformIntoMap()
	if err != nil {
		log.Println(err)
	}
	list, ok := intoMap[matchString[1]].([]interface{})
	if !ok {
		list = make([]interface{}, 0)
	}
	list = append(list, matchString[2])
	intoMap[matchString[1]] = list
	err = Write(intoMap)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(1)
}

func LRANGE(matchString []string) {
	intoMap, err := LoadJsonTransformIntoMap()
	if err != nil {
		log.Println(err)
	}
	list, ok := intoMap[matchString[1]].([]interface{})
	if !ok {
		fmt.Println("list not exist")
		return
	}
	begin, err := strconv.Atoi(matchString[2])
	if err != nil {
		log.Println(err)
		return
	}
	end, err := strconv.Atoi(matchString[3])
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(list[begin:end]...)
}

func SADD(matchString []string) {
	intoMap, err := LoadJsonTransformIntoMap()
	if err != nil {
		log.Println(err)
	}
	set, ok := intoMap[matchString[1]].(map[string]interface{})
	if !ok {
		set = make(map[string]interface{})
	}

	set[matchString[2]] = true
	intoMap[matchString[1]] = set
	err = Write(intoMap)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(1)
}

func SMEMBER(matchString []string) {
	intoMap, err := LoadJsonTransformIntoMap()
	if err != nil {
		log.Println(err)
		return
	}
	for key := range intoMap[matchString[1]].(map[string]interface{}) {
		fmt.Println(key)
	}
}
