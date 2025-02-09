package config

type Config struct {
	App           AppConfig           `mapstructure:"app"`
	Database      DatabaseConfig      `mapstructure:"database"`
	Cache         CacheConfig         `mapstructure:"cache"`
	ObjectStorage ObjectStorageConfig `mapstructure:"object_storage"`
	MessageBroker MessageBrokerConfig `mapstructure:"message_broker"`
	Jwt           JwtConfig           `mapstructure:"jwt"`
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
	SslMode  string
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

type CacheConfig struct {
	Type     string
	Host     string
	Port     int
	Username string
	Password string
}

type JwtConfig struct {
	ExpirationInSecond int64
	Issuer             string
	SigningMethod      string
	SignatureKey       string
}
