package client

import (
	"Redis_test/db"
	"Redis_test/match"
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
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
func readSelect() string {
	line, _, err := reader.ReadLine()
	if err != nil {
		log.Println(err)
	}
	return string(line)
}
func Start() {
	for {
		exec.Command("powershell", "Clear-Host").Run()
		fmt.Println("1. 命令行")
		fmt.Println("2. 使用说明")
		fmt.Println("3. 退出程序")
		switch readSelect() {
		case "1":
			StartCmd()
			break
		case "2":
			UseCase()
			break
		case "3":
			return
		}
	}
}

func StartCmd() {
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
		case "exit":
			return
		default:
			fmt.Println("no such command")
		}
	}
}

func UseCase() {
	fmt.Print("SET Key value （将键值存储起来）\nSETNX Key value （如果键存在 返回 0 如果键不存在 返回 1   并存储值)\nGET Key （获得对应键所对应的值）\nDEl Key   （删除所对应的键和值）\nSADD SetName value （向一个集合（SetName）中添加一个元素（value））\nSMEMBER SetName   （获取一个集合内的所有元素）\nLPUSH ListName value （向一个列表（ListName）中添加一个元素（value））\nLRANGE  ListName  start（起始位置） end（结束位置）（获得一个列表 从 start到end 的所有元素）\n 按任意键继续。。。。")
	fmt.Scanln()
}
