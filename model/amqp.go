package model

type AmqpData struct {
	Host string `json:"host" bson:"host"`
	Port string `json:"port" bson:"port"`
}
