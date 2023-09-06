package client

import (
	"Redis_test/db"
	"Redis_test/match"
	"bufio"
	"fmt"
	"log"
	"os"
)

var reader *bufio.Reader

func init() {
	reader = bufio.NewReader(os.Stdin)
}

func readCmd() string {
	fmt.Print("redis-cli>")
	line, _, err := reader.ReadLine()
	if err != nil {
		log.Println(err)
	}
	return string(line)
}
func Start() {
	for {
		cmd := readCmd()
		parseCmd, s := match.ParseCmd(cmd)
		switch s {
		case "GET":
			db.GET(parseCmd)
			break
		case "DEL":
			db.DEL(parseCmd)
			break
		case "SET":
			db.SET(parseCmd)
			break
		case "SETNX":
			db.SETNX(parseCmd)
			break
		case "LPUSH":
			db.LPUSH(parseCmd)
			break
		case "LRANGE":
			db.LRANGE(parseCmd)
			break
		case "SADD":
			db.SADD(parseCmd)
			break
		case "SMEMBER":
			db.SMEMBER(parseCmd)
			break
		default:
			fmt.Println("no such command")
		}
	}
}
