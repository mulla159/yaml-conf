package yamlconf

import (
	"fmt"
	"log"
	"net/url"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type CustomURL url.URL

func (cu *CustomURL) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var urlString string

	err := unmarshal(&urlString)
	if err != nil {
		return err
	}

	u, err := url.ParseRequestURI(urlString)
	if err != nil {
		return err
	}

	*cu = CustomURL(*u)
	return nil
}

type AppConfig struct {
	Port        int       `yaml:"port"`
	DbURL       CustomURL `yaml:"db_url"`
	JaegerURL   CustomURL `yaml:"jaeger_url"`
	SentryURL   CustomURL `yaml:"sentry_url"`
	KafkaBroker string    `yaml:"kafka_broker"`
	AppID       string    `yaml:"some_app_id"`
}

func (conf *AppConfig) Get(fileName string) {
	if fileName == "" {
		log.Fatalf("Не указано имя файла конфигурации")
	}

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Ошибка при открытии файла конфигурации: %v", err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatalf("Ошибка при закрыттии файла конфигурации: %v\n", err)
		}
	}()

	if err := yaml.NewDecoder(file).Decode(&conf); err != nil {
		log.Fatalf("Ошибка при декодировании yaml-файла в структуру: %v\n", err)
	}
}

func (conf *AppConfig) Print() {
	fmt.Println("Configs:")
	fmt.Println("Port: ", conf.Port)
	fmt.Println("DBUrl: ", conf.DbURL)
	fmt.Println("JaegerUrl: ", conf.JaegerURL)
	fmt.Println("SentryUrl: ", conf.SentryURL)
	fmt.Println("KafkaBroker: ", conf.KafkaBroker)
}
