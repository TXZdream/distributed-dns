package main

import (
	"fmt"
	"github.com/txzdream/distributed-dns/dns"
	"github.com/txzdream/distributed-dns/logger"
	"log"
	"net"
	"time"

	flag "github.com/spf13/pflag"
	"google.golang.org/grpc"

	mygrpc "github.com/txzdream/distributed-dns/grpc"

	"github.com/txzdream/distributed-dns/server/dnsUdp"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
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
	k := 8
	kad := dns.Init(uint16(k), id, *access, *otherNodeAccess)
	// 每隔3分钟进行一次更新
	go func() {
		for {
			updateInterval := time.Duration(3)
			kad.Update()
			time.Sleep(updateInterval * time.Second)
		}
	}()
	// 启动grpc服务器
	lis, err := net.Listen("tcp", fmt.Sprintf("%v", *access))
	if err != nil {
		log.Fatalf("failed to listen grpc: %v", err)
	}
	log.Printf("Listening on: %s", *access)
	// 添加grpc日志
	defer logger.Logger.Sync()
	gs := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			// grpc_zap.StreamServerInterceptor(logger.Logger),
			grpc_recovery.StreamServerInterceptor(),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			// grpc_zap.UnaryServerInterceptor(logger.Logger),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)
	ddns := kad.(*dns.DistributeDNS)
	// 启动监听udp的53端口
	dnsUdp.Start(ddns)
	mygrpc.RegisterKademilaServer(gs, ddns)
	gs.Serve(lis)
}
