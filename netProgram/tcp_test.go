package netProgram

//1234567
import "testing"

func TestTcpServer(t *testing.T) {
	// 定义一个测试函数 TestTcpServer，接受一个指向 testing.T 的指针作为参数
	TcpServer()
	// 调用 TcpServer 函数，进行 TCP 服务器的测试
}

func TestTcpClient(t *testing.T) {
	TcpClient()
}
