package test

import (
	elasticsearch "github.com/elastic/go-elasticsearch/v7"
	"log"
)

func run() {
	client, err := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://192.168.1.231:9200"},
		Username:  "elastic",
		Password:  "123456",
	})
	if err != nil {
		log.Fatal(err)
	}
	res, err := client.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	log.Println(res)
}

func main() {
	run()
}
