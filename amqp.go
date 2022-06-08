package service

import (
	ramqp "github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
	amqp "github.com/wms3001/goAmqp/model"
	"os"
)

func OpenConn() *amqp.Conn {
	var conn = &amqp.Conn{}
	work, _ := os.Getwd()
	viper.SetConfigName("amqp")
	viper.SetConfigType("yml")
	viper.AddConfigPath(work + "/conf")
	viper.ReadInConfig()
	host := viper.GetString("amqp.host")
	port := viper.GetString("amqp.port")
	user := viper.GetString("amqp.user")
	pass := viper.GetString("amqp.pass")
	connnect, err := ramqp.Dial("amqp://" + user + ":" + pass + "@" + host + ":" + port + "/")
	if err != nil {
		conn.Code = -1
		conn.Message = err.Error()
		return conn
	}
	conn.Code = 1
	conn.Message = "success"
	conn.Connection = connnect
	return conn
}

func OpenChannel(conn *amqp.Conn) *amqp.Channel {
	var channel = &amqp.Channel{}
	ch, err := conn.Connection.Channel()
	if err != nil {
		channel.Code = -1
		channel.Message = err.Error()
		return channel
	}
	channel.Code = 1
	channel.Message = "success"
	channel.Channel = ch
	return channel
}

func DeclareQueue(channel *amqp.Channel, queueName string) *amqp.Queue {
	var queue = &amqp.Queue{}
	work, _ := os.Getwd()
	viper.SetConfigName("amqp")
	viper.SetConfigType("yml")
	viper.AddConfigPath(work + "/conf")
	viper.ReadInConfig()
	durable := viper.GetBool("amqp.durable")
	autoDelete := viper.GetBool("amqp.autoDelete")
	exclusive := viper.GetBool("amqp.exclusive")
	noWait := viper.GetBool("amqp.noWait")
	xMessageTtl := viper.GetInt("amqp.x-message-ttl")
	var args ramqp.Table = map[string]interface{}{}
	args["x-message-ttl"] = xMessageTtl
	q, err := channel.Channel.QueueDeclare(
		queueName,  // name
		durable,    // durable 持久化
		autoDelete, // delete when unused 是否自动删除队列
		exclusive,  // exclusive 排他
		noWait,     // no-wait
		args,       // arguments
	)
	if err != nil {
		queue.Code = -1
		queue.Message = err.Error()
		return queue
	}
	queue.Code = 1
	queue.Message = "success"
	queue.Queue = &q
	return queue
}
