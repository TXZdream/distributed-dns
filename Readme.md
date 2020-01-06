# 基于DHT的分布式DNS系统

## 功能说明

## 代码结构说明

- dns/

包含一个实现了kademlia接口和grpc接口的结构体，用以在节点间提供grpc服务

- grpc/

根据Kademlia算法要求提供了四个RPC接口用以实现各种操作

- kademlia/

定义了kademlia接口，该接口包含了一系列方法用以实现Kademlia算法

- main.go
