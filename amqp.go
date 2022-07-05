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
}

func (goAmqp *GoAmqp) OpenConn() *goCommon.Resp {
	var resp = &goCommon.Resp{}
	//work, _ := os.Getwd()
	//viper.SetConfigName("amqp")
	//viper.SetConfigType("yml")
	//viper.AddConfigPath(work + "/conf")
	//viper.ReadInConfig()
	//host := viper.GetString("amqp.host")
	//port := viper.GetString("amqp.port")
	//user := viper.GetString("amqp.user")
	//pass := viper.GetString("amqp.pass")
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

func (goAmqp *GoAmqp) DeclareQueue(queueName string) *goCommon.Resp {
	var resp = &goCommon.Resp{}
	//work, _ := os.Getwd()
	//viper.SetConfigName("amqp")
	//viper.SetConfigType("yml")
	//viper.AddConfigPath(work + "/conf")
	//viper.ReadInConfig()
	//durable := viper.GetBool("amqp.durable")
	//autoDelete := viper.GetBool("amqp.autoDelete")
	//exclusive := viper.GetBool("amqp.exclusive")
	//noWait := viper.GetBool("amqp.noWait")
	//xMessageTtl := viper.GetInt("amqp.x-message-ttl")

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
