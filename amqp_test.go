package goAmqp

import (
	"fmt"
	"log"
	"strconv"
	"testing"
)

func TestOpenConn(t *testing.T) {
	goAmqp := GoAmqp{}
	goAmqp.Host = "192.168.4.81"
	goAmqp.Port = "5672"
	goAmqp.User = "logistics"
	goAmqp.Pass = "logistics"
	goAmqp.Durable = true
	goAmqp.AutoDelete = false
	goAmqp.NoWait = false
	goAmqp.Xmessagettl = 60000
	resp := goAmqp.OpenConn()
	defer goAmqp.CloseConn()
	fmt.Println(resp)
}

func TestGoAmqp_OpenChannel(t *testing.T) {
	goAmqp := GoAmqp{}
	goAmqp.Host = "192.168.4.81"
	goAmqp.Port = "5672"
	goAmqp.User = "logistics"
	goAmqp.Pass = "logistics"
	goAmqp.Durable = true
	goAmqp.AutoDelete = false
	goAmqp.NoWait = false
	goAmqp.Xmessagettl = 60000
	resp := goAmqp.OpenConn()
	defer goAmqp.CloseConn()
	fmt.Println(resp)
	resp1 := goAmqp.OpenChannel()
	defer goAmqp.CloseChannel()
	fmt.Println(resp1)
}

func TestGoAmqp_DeclareQueue(t *testing.T) {
	goAmqp := GoAmqp{}
	goAmqp.Host = "192.168.4.81"
	goAmqp.Port = "5672"
	goAmqp.User = "logistics"
	goAmqp.Pass = "logistics"
	goAmqp.Durable = true
	goAmqp.AutoDelete = false
	goAmqp.NoWait = false
	goAmqp.Xmessagettl = 60000
	resp := goAmqp.OpenConn()
	defer goAmqp.CloseConn()
	fmt.Println(resp)
	resp1 := goAmqp.OpenChannel()
	defer goAmqp.CloseChannel()
	fmt.Println(resp1)
	resp2 := goAmqp.DeclareQueue("mQueue")
	fmt.Println(resp2)

}

func TestGoAmqp_SendMessage(t *testing.T) {
	goAmqp := GoAmqp{}
	goAmqp.Host = "192.168.4.81"
	goAmqp.Port = "5672"
	goAmqp.User = "logistics"
	goAmqp.Pass = "logistics"
	goAmqp.Durable = true
	goAmqp.AutoDelete = false
	goAmqp.NoWait = false
	goAmqp.Xmessagettl = 60000
	resp := goAmqp.OpenConn()
	defer goAmqp.CloseConn()
	fmt.Println(resp)
	resp1 := goAmqp.OpenChannel()
	defer goAmqp.CloseChannel()
	fmt.Println(resp1)
	resp2 := goAmqp.DeclareQueue("mQueue")
	fmt.Println(resp2)
	resp3 := goAmqp.SendMessage("test")
	fmt.Println(resp3)
	for n := 1; n <= 50000; n++ {
		goAmqp.SendMessage(strconv.Itoa(n) + "test")
	}
}

func BenchmarkSendMessage(b *testing.B) {
	goAmqp := GoAmqp{}
	goAmqp.Host = "192.168.4.81"
	goAmqp.Port = "5672"
	goAmqp.User = "logistics"
	goAmqp.Pass = "logistics"
	goAmqp.Durable = true
	goAmqp.AutoDelete = false
	goAmqp.NoWait = false
	goAmqp.Xmessagettl = 60000
	resp := goAmqp.OpenConn()
	defer goAmqp.CloseConn()
	fmt.Println(resp)
	resp1 := goAmqp.OpenChannel()
	defer goAmqp.CloseChannel()
	fmt.Println(resp1)
	resp2 := goAmqp.DeclareQueue("mQueue")
	fmt.Println(resp2)

	for i := 0; i < b.N; i++ { // b.N，测试循环次数
		resp3 := goAmqp.SendMessage("test")
		fmt.Println(resp3)
	}
}

func TestGoAmqp_Consume(t *testing.T) {
	goAmqp := GoAmqp{}
	goAmqp.Host = "192.168.4.81"
	goAmqp.Port = "5672"
	goAmqp.User = "logistics"
	goAmqp.Pass = "logistics"
	goAmqp.Durable = true
	goAmqp.AutoDelete = false
	goAmqp.NoWait = false
	goAmqp.Xmessagettl = 60000
	resp := goAmqp.OpenConn()
	defer goAmqp.CloseConn()
	fmt.Println(resp)
	resp1 := goAmqp.OpenChannel()
	defer goAmqp.CloseChannel()
	fmt.Println(resp1)
	resp2 := goAmqp.DeclareQueue("mQueue")
	fmt.Println(resp2)
	resp3 := goAmqp.Consume()
	fmt.Println(resp3)
	for d := range goAmqp.Delivery {
		log.Printf("Received a message: %s", d.Body)
	}
}
