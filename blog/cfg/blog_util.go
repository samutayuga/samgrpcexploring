package cfg

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	ServerPort   int
	EndpointPort int
	DbHost       string
	DbPort       int
	DbUser       string
	DbName       string
	ResourceBlog string
}

func LoadConfig() Config {
	configDir := os.Getenv("CONFIG_DIR")
	configFile := fmt.Sprintf("%s/%s", configDir, "blog.yaml")
	//viper.AddConfigPath(configFile)
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configFile)
	if errRead := viper.ReadInConfig(); errRead != nil {
		panic(errRead)
	} else {

		svrPort := viper.GetInt("server.port")
		httpPort := viper.GetInt("endpoint.port")
		//database properties
		dbHost := viper.GetString("database.host")
		dbPort := viper.GetInt("database.port")
		dbUser := viper.GetString("database.username")
		dbName := viper.GetString("database.dbname")
		resBlog := viper.GetString("endpoint.blogResource")
		blogConfig := Config{ServerPort: svrPort, EndpointPort: httpPort,
			DbHost: dbHost, DbPort: dbPort,
			DbUser: dbUser, DbName: dbName,
			ResourceBlog: resBlog}
		log.Printf("Got config %v\n", blogConfig)
		return blogConfig
	}

}
