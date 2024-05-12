package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("No .env file found")
	}
}
func LoadConfigs(cfg *Config) {
	cfg.MongoConnection = os.Getenv("MONGODB_URI")
	cfg.MongoDBName = os.Getenv("MONGODB_DB_NAME")
	cfg.AppPort = os.Getenv("APP_PORT")
	log.Println(cfg)
}
