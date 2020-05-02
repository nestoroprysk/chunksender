package config

type Config struct {
	Email *EmailConfig `json:"email"`
}

type EmailConfig struct {
	Sender    SenderConfig `json:"sender"`
	Receivers []string     `json:"receivers"`
}

type SenderConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}
