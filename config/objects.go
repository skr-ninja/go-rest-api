package config

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type App struct {
	Router *gin.Engine
	DB     *gorm.DB
	Config Config
}

type Config struct {
	Server   ServerConfiguration   `mapstructure:",squash"`
	Database DatabaseConfiguration `mapstructure:",squash"`
	JWTCnf   JWTConfiguration      `mapstructure:",squash"`
}

type ServerConfiguration struct {
	Debug                bool   `mapstructure:"DEBUG"`
	Port                 string `mapstructure:"SERVER_PORT"`
	VPCProxyCIDR         string `mapstructure:"VPC_CIDR"`
	MessengerURL         string `mapstructure:"MESSENGER_URL"`
	LimitCountPerRequest int64
}

type DatabaseConfiguration struct {
	Name     string `mapstructure:"DB_NAME"`
	UserName string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	Host     string `mapstructure:"DB_HOST"`
	Port     string `mapstructure:"DB_PORT"`
	DbDriver string `mapstructure:"DB_DRIVER"`
	//SSLMode  string `mapstructure:"MAIN_DB_SSL_MODE"`
}

type JWTConfiguration struct {
	JWTSecret         string `mapstructure:"API_SECRET"`
	JWTValidTimeInMin int32  `mapstructure:"TOKEN_HOUR_LIFESPAN"`
}
