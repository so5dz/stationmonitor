package config

type StationMonitorConfig struct {
	Redis  RedisConfig  `json:"redis"`
	APRSIS APRSISConfig `json:"aprsis"`
}

type RedisConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type APRSISConfig struct {
	Host   string `json:"host"`
	Port   int    `json:"port"`
	User   string `json:"user"`
	Pass   int    `json:"pass"`
	Filter string `json:"filter"`
}
