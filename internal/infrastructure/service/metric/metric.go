package metric

import (
	"fmt"
	"log"

	"github.com/DataDog/datadog-go/statsd"
)

var client *statsd.Client

// TODO finish implementation
func InitDatadogStatsd() *statsd.Client {
	// Crear un cliente statsd
	// options := statsd.Options{
	// 	Namespace: "apiPayment.service.",
	// 	Tags:      []string{"env:local"},
	// }
	var err error
	client, err = statsd.New("", nil)
	if err != nil {
		log.Fatalf("Cannot connect to statsd: %s", err)
	}
	defer client.Close()

	return client
}

// TODO simulation count function
func Count(name string, value int64, tags []string, rate float64) {
	fmt.Printf("counting values")
}
