package main

import (
	"distributed-dns/dns"
	"fmt"
	"log"
	"net"
	"time"

	flag "github.com/spf13/pflag"
	"google.golang.org/grpc"

	mygrpc "distributed-dns/grpc"
	uuid "github.com/satori/go.uuid"
)

func main() {
	// 从用户输入解析当前节点的相关参数
	access := flag.StringP("access", "a", "", "ip:port格式，表示当前节点kademlia的访问地址")
	otherNodeAccess := flag.StringP("node", "n", "", "若需要加入一个集群，则需输入ip:port格式的访问地址，用以加入集群；若为空则新建一个集群")
	flag.Parse()
	// 生成id号
	id, err := dns.CalculateHash(uuid.NewV4().String())
	if err != nil {
		panic(err)
	}
	k := 3
	kad := dns.Init(uint16(k), id, *access, *otherNodeAccess)
	// 每隔3分钟进行一次更新
	go func() {
		updateInterval := time.Duration(3)
		kad.Update()
		time.Sleep(updateInterval * time.Minute)
	}()
	// 启动grpc服务器
	lis, err := net.Listen("tcp", fmt.Sprintf("%v", *access))
	if err != nil {
		log.Fatalf("failed to listen grpc: %v", err)
	}
	log.Printf("Listening on: %s", *access)
	gs := grpc.NewServer()
	ddns := kad.(dns.DistributeDNS)
	mygrpc.RegisterKademilaServer(gs, &ddns)
	gs.Serve(lis)
}
