package netProgram

import (
	"log"
	"net"
	"sync"
)

const tcp = "tcp"

// 服务端
func TcpServer() {
	//A.基于某个地址建立监听
	//服务端地址
	address := "127.0.0.1:5678"
	listener, err := net.Listen(tcp, address)
	if err != nil {
		log.Fatalln(err)
	}
	//关闭监听的操作
	defer listener.Close()
	log.Printf("%s server is listening on %s\n", tcp, listener.Addr())

	//B.接收客户端的连接请求
	//循环接受
	for {
		//  阻塞接受
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		//处理连接，读写操作
		// 日志连接的远程地址(client addr)
		log.Printf("accept from %s\n", conn.RemoteAddr())

	}
}

// 客户端
func TcpClient() {
	//TCP服务端地址
	address := "127.0.0.1:5678"
	//模拟多客户端
	//并发的客户端请求
	num := 10
	wg := sync.WaitGroup{}
	wg.Add(num)
	for i := 0; i < num; i++ {
		//并发的请求
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			// A. 建立连接
			conn, err := net.Dial(tcp, address)
			if err != nil {
				log.Println(err)
			}
			// 保证关闭
			defer conn.Close()
			log.Printf("connection is establish,client addr is %s\n", conn.LocalAddr())
		}(&wg)
	}
	wg.Wait()
}
