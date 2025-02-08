package config

type Config struct {
	AppConfig           `mapstructure:"app"`
	DatabaseConfig      `mapstructure:"database"`
	ObjectStorageConfig `mapstructure:"object_storage"`
	MessageBrokerConfig `mapstructure:"message_broker"`
}

type AppConfig struct {
	Name    string
	Address string
	Port    string
}

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

type ObjectStorageConfig struct {
	Type      string
	AccessKey string
	ScretKey  string
	Host      string
	Region    string
}

type MessageBrokerConfig struct {
	Type     string
	Brokers  []string
	Username string
	Password string
}
