package goAmqp

import (
	amqp "github.com/rabbitmq/amqp091-go"
	//amqp "github.com/wms3001/goAmqp/model"
	"github.com/wms3001/goCommon"
)

type GoAmqp struct {
	Host        string
	Port        string
	User        string
	Pass        string
	Durable     bool
	AutoDelete  bool
	Exclusive   bool
	NoWait      bool
	Xmessagettl int64
	Connection  *amqp.Connection
	Channel     *amqp.Channel
	Queue       *amqp.Queue
	Delivery    <-chan amqp.Delivery
}

func (goAmqp *GoAmqp) OpenConn() *goCommon.Resp {
	var resp = &goCommon.Resp{}
	connnect, err := amqp.Dial("amqp://" + goAmqp.User + ":" + goAmqp.Pass + "@" + goAmqp.Host + ":" + goAmqp.Port + "/")
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 1
		resp.Message = "connected"
		goAmqp.Connection = connnect
	}
	return resp
}

func (goAmqp *GoAmqp) CloseConn() {
	goAmqp.Connection.Close()
}

func (goAmqp *GoAmqp) OpenChannel() *goCommon.Resp {
	var resp = &goCommon.Resp{}
	ch, err := goAmqp.Connection.Channel()
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 1
		resp.Message = "success"
		goAmqp.Channel = ch
	}
	return resp
}

func (goAmqp *GoAmqp) CloseChannel() {
	goAmqp.Channel.Close()
}

func (goAmqp *GoAmqp) DeclareQueue(queueName string) *goCommon.Resp {
	var resp = &goCommon.Resp{}
	var args amqp.Table = map[string]interface{}{}
	args["x-message-ttl"] = goAmqp.Xmessagettl
	q, err := goAmqp.Channel.QueueDeclare(
		queueName,         // name
		goAmqp.Durable,    // durable 持久化
		goAmqp.AutoDelete, // delete when unused 是否自动删除队列
		goAmqp.Exclusive,  // exclusive 排他
		goAmqp.NoWait,     // no-wait
		args,              // arguments
	)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 1
		resp.Message = "success"
		goAmqp.Queue = &q
	}
	return resp
}

func (goAmqp *GoAmqp) SendMessage(msg string) *goCommon.Resp {
	var resp = &goCommon.Resp{}
	err := goAmqp.Channel.Publish(
		"",                //交换
		goAmqp.Queue.Name, //队列名称
		false,             //强制为
		false,             //立即
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 1
		resp.Message = "success"
	}
	return resp
}

func (goAmqp *GoAmqp) Consume() *goCommon.Resp {
	var resp = &goCommon.Resp{}
	msgs, err := goAmqp.Channel.Consume(
		goAmqp.Queue.Name, //队列名称
		"",                //消费者
		true,              //自动确认
		false,             //排他
		false,
		false,
		nil,
	)
	if err != nil {
		resp.Code = -1
		resp.Message = err.Error()
	} else {
		resp.Code = 1
		resp.Message = "success"
		goAmqp.Delivery = msgs
	}
	return resp
}
