package nats

import (
	"ecommerce-microservices/services/product/internal/domain/models"
	"encoding/json"
	"log"

	"github.com/nats-io/nats.go"
)

func PublishProduct(productModel *models.CreateProductModel) {
	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	js, err := nc.JetStream()
	if err != nil {
		log.Fatalf("JetStream başlatılamadı: %v", err)
	}

	// Stream tanımı (sadece bir kez çalışır)
	js.AddStream(&nats.StreamConfig{
		Name:     "PRODUCTS",             // Stream adı
		Subjects: []string{"products.*"}, // Konuları
		Storage:  nats.FileStorage,       // Diskte kalıcı
	})

	Publish(js, productModel)

}

func Publish(js nats.JetStreamContext, productModel *models.CreateProductModel) {

	subject := "products.created"

	publishedData, err := json.Marshal(productModel)
	if err != nil {
		log.Printf("Product JSON'a çevrilemedi: %v", err)
		return
	}

	ack, err := js.Publish(subject, publishedData)
	if err != nil {
		log.Printf("JetStream publish hatası: %v", err)
		return
	}

	log.Printf("JetStream mesaj yayınlandı. Stream: %s, Seq: %d", ack.Stream, ack.Sequence)

}
