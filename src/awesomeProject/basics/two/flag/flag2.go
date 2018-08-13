package main

import (
	"math/big"
	"fmt"
	"flag"
)

type NodeInfo struct {
	Mode   string //1:三所的节点,2:解码服务连接的节点，3:TMS连接的节点,4:表示sdk连接的节点；5:经营平台所连接的box节点
	Domain string
}

type CFG struct {
	NInfo                   *NodeInfo
	BootNodeIP              string
	GrpcServerAddress       string
	ResponseTimeOutDuration int64
	PrivateKeyUuid          string
	PassPhrase              string
	KeystorePath            string
	MysqlPath               string
	EtherURL                string
	Retries                 int
	DelayedBlocks           *big.Int
	CursorBlocks            *big.Int
	BlockNoFilePath         string
	NonceFilePath           string
}

//同时编译 go build ./flag2 ./jsonparse.go 后执行 ./flag2 -json config.json 或 ./flag2 -json abcd 可以测试效果
func main() {
	path := flag.String("json", "config.json", "ins nums")
	flag.Parse()

	jsonParser := NewJsonParser()
	cfg := new(CFG)
	err := jsonParser.Load(*path, cfg)
	if err != nil {
		fmt.Println("this is fail")
	}

	fmt.Println("------>", cfg)
}
