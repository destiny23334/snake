package global

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	Logger *zap.Logger  // 日志
	config *viper.Viper // 配置
	DB     *gorm.DB     // 数据库
)

func init() {
	Logger, _ = zap.NewProduction()
}

func LoadConfig() {
	config = viper.New()
	config.SetConfigName("config_dev") // name of config file (without extension)
	config.SetConfigType("yaml")       // REQUIRED if the config file does not have the extension in the name
	config.AddConfigPath("config")     // config.yaml存放的路径，这里就放到项目下，也可以添加多个路径
	err := config.ReadInConfig()       // Find and read the config file
	if err != nil {                    // Handle errors reading the config file
		Logger.Error("Fatal error config file", zap.String("message", err.Error()))
	}
}

func InitDatabase() {
	if config == nil {
		Logger.Error("config is nil")
	}
	// 连接数据库
	host := config.Get("pgsql.host").(string)
	user := config.Get("pgsql.user-name").(string)
	password := config.Get("pgsql.password").(string)
	dbName := config.Get("pgsql.db-name").(string)
	port := config.Get("pgsql.port").(int)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		host, user, password, dbName, port)
	//var err error
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		Logger.Error("connect database failed", zap.String("message", err.Error()))
	}
	DB = db
	Logger.Info("connect database success")
}
