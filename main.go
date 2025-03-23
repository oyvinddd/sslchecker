package main

import (
)

func main() {
	cfg := app.NewConfig(":8888", nil)
	log.Fatal(app.New(cfg))
/*
	var client = http.Client{}
	res, err := client.Head("https://dummyjson.com/test")
	if err != nil {
		panic(err)
	}

	if len(res.TLS.PeerCertificates) > 0 {
		cert := res.TLS.PeerCertificates[0]

		issuedAt := cert.NotBefore
		expiresAt := cert.NotAfter

		fmt.Println("Issued At: ", issuedAt)
		fmt.Println("Expires At: ", expiresAt)
	}

	fmt.Println("Status: ", res.StatusCode)
	fmt.Println("Content-Length: ", res.ContentLength)
*/
}

