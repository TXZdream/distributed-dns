package dnsUdp

import (
	"bytes"
	"encoding/binary"
	"log"
	"net"
	"strconv"

	"github.com/txzdream/distributed-dns/dns"
	"github.com/txzdream/distributed-dns/server/resolution"

	"golang.org/x/net/dns/dnsmessage"
)

const (
	// DNS server default port
	udpPort int = 9053
	// DNS packet max length
	packetLen int = 512
)

// DNSServer will do Listen, Query and Send.
type DNSServer interface {
	Listen()
	Query(Packet)
}

// DNSService is an implementation of DNSServer interface.
type DNSService struct {
	conn *net.UDPConn
	dns  *dns.DistributeDNS
	//forwarders []net.UDPAddr
}

// Packet carries DNS packet payload and sender address.
type Packet struct {
	addr    net.UDPAddr
	message dnsmessage.Message
}

// question to string，已测试，可用
func quesToString(q dnsmessage.Question) string {
	// 将报文中Name的解析成域名再返回
	fullQues := q.Name.String()
	fullQuesLen := len(fullQues)
	var arrStr []string
	var ques string
	// 把字符串中的每一个字符都当做字符串存到一个新的字符串数组中
	for _, s := range fullQues {
		arrStr = append(arrStr, string(s))
	}
	// 从新的字符串数组中提取字符得到ques
	for i := 0; i < fullQuesLen; i++ {
		tmpLen, _ := strconv.Atoi(arrStr[i])
		if tmpLen == 0 {
			break
		}
		for j := 1; j <= tmpLen; j++ {
			ques += arrStr[i+j]
		}
		i += tmpLen
		if i != fullQuesLen-2 {
			ques += "."
		}
	}
	return ques
}

// Start conveniently init every parts of DNS service.
func Start(disDns *dns.DistributeDNS) *DNSService {
	newDNSService := DNSService{
		dns: disDns,
	}
	go newDNSService.Listen()
	log.Printf("DNS listening on %d", udpPort)
	return &newDNSService
}

// Listen starts a DNS server on port given in config file
func (s *DNSService) Listen() {
	var err error
	s.conn, err = net.ListenUDP("udp", &net.UDPAddr{Port: udpPort})
	if err != nil {
		log.Fatal(err)
	}
	defer s.conn.Close()

	for {
		buf := make([]byte, packetLen)
		_, addr, err := s.conn.ReadFromUDP(buf)
		if err != nil {
			log.Println(err)
			continue
		}
		var m dnsmessage.Message
		err = m.Unpack(buf)
		if err != nil {
			log.Println(err)
			continue
		}
		if len(m.Questions) == 0 {
			continue
		}
		go s.Query(Packet{*addr, m})
	}
}

// Query lookup answers for DNS message.
func (s *DNSService) Query(p Packet) {
	q := p.message.Questions[0]
	// qStr := quesToString(q)
	s.dns.RtLock.Lock()
	defer s.dns.RtLock.Unlock()
	s.dns.DataLock.Lock()
	defer s.dns.DataLock.Unlock()
	// 根据域名寻找ip
	ok, ip, err := s.dns.GetDataRecursive(q.Name.String()[0 : len(q.Name.String())-1])
	if err != nil {
		log.Println(err)
		p.message.Header.RCode = 2 // RCodeServerFailure RCode = 2
		go sendPacket(s.conn, p.message, p.addr)
		return
	}
	if ok == false {
		p.message.Header.RCode = 3 // RCodeNameError RCode = 3
		go sendPacket(s.conn, p.message, p.addr)
		return
	}
	// 这里需要完成Answers
	p.message.Header.RCode = 0 // RCodeSuccess RCode = 0
	resource := generateAnswerMsg(ip, q.Name.String())
	p.message.Answers = append(p.message.Answers, resource)
	go sendPacket(s.conn, p.message, p.addr)
}

func sendPacket(conn *net.UDPConn, message dnsmessage.Message, addr net.UDPAddr) {
	packed, err := message.Pack()
	if err != nil {
		log.Println(err)
		return
	}
	_, err = conn.WriteToUDP(packed, &addr)
	if err != nil {
		log.Println(err)
	}
}

// 已测试，可用
func IntToByte(n int) byte {
	data := int64(n)
	//好像这个bytebuf长度默认为8
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, data)
	return bytebuf.Bytes()[7]
}

// 生成ipv4类型的资源信息
func generateAnswerMsg(ipAddr, domain string) dnsmessage.Resource {
	ipArr := resolution.Ipv4ToArr(ipAddr)
	resource := dnsmessage.Resource{
		Header: dnsmessage.ResourceHeader{
			Name:  dnsmessage.MustNewName(domain),
			Type:  dnsmessage.TypeA,
			Class: dnsmessage.ClassINET,
		},
		Body: &dnsmessage.AResource{A: [4]byte{IntToByte(ipArr[0]), IntToByte(ipArr[1]), IntToByte(ipArr[2]), IntToByte(ipArr[3])}},
	}
	return resource
}
