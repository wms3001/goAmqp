package goAmqp

import (
	"fmt"
	"testing"
)

func TestOpenConn(t *testing.T) {
	goAmqp := GoAmqp{}
	goAmqp.Host = "192.168.4.81"
	goAmqp.Port = "5672"
	goAmqp.User = "logistics"
	goAmqp.Pass = "logistics"
	resp := goAmqp.OpenConn()
	fmt.Println(resp)
}
