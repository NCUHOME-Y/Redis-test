package db

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

var file *os.File

func init() {
	var err error
	file, err = os.OpenFile("redis.json", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Println(err)
		return
	}
}

func LoadJsonTransformIntoMap() (map[string]interface{}, error) {
	m := make(map[string]interface{})
	fileStat, err := file.Stat()
	if err != nil {
		return nil, err
	}
	if fileStat.Size() == 0 {
		return m, nil
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return nil, err
	}

	all, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(all, &m)
	if err != nil {
		if err == io.EOF {
			return m, nil
		}
		return nil, err
	}
	return m, nil
}

func Write(jsonString map[string]interface{}) error {
	marshal, err := json.Marshal(jsonString)
	if err != nil {
		return err
	}

	err = file.Truncate(0) // 清空文件内容
	if err != nil {
		return err
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return err
	}
	_, err = file.Write(marshal)
	if err != nil {
		return err
	}
	return nil
}
