package main

import (
	"flag"
)

func main() {
	// client := client.New("http://localhost:8080")
	// price, err := client.FetchPrice(context.Background(), "ET")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("price: %+v\n", price)
	// return
	listenAddr := flag.String("listenAddr", ":8080", "http listen address")
	flag.Parse()
	svc := NewLoggingService(NewMetricsService(&priceFetcher{}))
	server := NewJSONAPIServer(*listenAddr, svc)
	// price, err := svc.FetchPrice(context.Background(), "ETH")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("price: %f", price)
	server.Run()
}
