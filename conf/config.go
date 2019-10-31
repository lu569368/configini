package conf

//AppConf ...
type AppConf struct {
	EdctConf `ini:"etcd"`
	Kafka    `ini:"kafka"`
}

//EdctConf ...
type EdctConf struct {
	//Address ...
	Address string `ini:"address"`
	//Timeout ...
	Timeout int `ini:"timeout"`
}

//Kafka ...
type Kafka struct {
	//Address ...
	Address string `ini:"address"`
	//Timeout ...
	Topic string `ini:"topic"`
}