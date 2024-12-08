package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string
}
type Config struct {
	Env        string               `yaml:"env" env:"ENV" env-default:"production" env-required:"true"` // providing evn location/-env-default:"production"  setting defult env as production
	StoragPath string               `yaml:"storage_path" env-required:"true"`
	HTTPServer `yaml:"http_server"` // struct embedding
}

func MustLoad()*Config {
	var configPath string
	configPath = os.Getenv("CONFIG_PATH")
	if configPath ==""{
		flags := flag.String("config","","path to the config file")
		flag.Parse()
		configPath = *flags
		if configPath ==""{
		log.Fatal("config path is not set")
		}
	}
	if _, err:= os.Stat(configPath);os.IsNotExist(err){
		log.Fatalf("config file does not exist:%s", configPath)
	}

	//serializing the config struct
	var cfg Config 
	err :=cleanenv.ReadConfig(configPath,&cfg)
	if err != nil{
		log.Fatalf("cannot read config file %s",err.Error())
	}
	return &cfg
}