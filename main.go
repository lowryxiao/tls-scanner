package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
	"net/http"
)

func main() {

	// Output:
	// https://edesoft.com
	// 200 OK
	// version: 771
	// TLS 1.2
	var (
		conn *tls.Conn
		err  error
	)

	// tlsConfig := http.DefaultTransport.(*http.Transport).TLSClientConfig
	tlsConfig := &tls.Config{}
	//MinVersion: tls.VersionSSL30,
	//MaxVersion: tls.VersionSSL30,
	//MaxVersion: tls.VersionTLS12,
	tlsConfig.MinVersion = tls.VersionTLS12
	tlsConfig.MaxVersion = tlsConfig.MinVersion
	tlsConfig.InsecureSkipVerify = true
	//tlsConfig.CipherSuites = []uint16{
	//	tls.TLS_RSA_WITH_RC4_128_SHA,
	//}

	c := &http.Client{
		Transport: &http.Transport{
			DialTLSContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				conn, err = tls.Dial(network, addr, tlsConfig)
				return conn, err
			},
		},
	}

	res, err := c.Get("https://edesoft.com")
	if err != nil {
		fmt.Println(err)
		logrus.Fatal(err)
	}
	defer res.Body.Close()

	versions := map[uint16]string{
		//tls.VersionSSL30: "SSL",
		tls.VersionTLS10: "TLS 1.0",
		tls.VersionTLS11: "TLS 1.1",
		tls.VersionTLS12: "TLS 1.2",
		tls.VersionTLS13: "TLS 1.3",
	}

	logrus.Info(res.Request.URL)
	logrus.Info(res.Status)

	v := conn.ConnectionState().Version
	logrus.Infof("version: %d", v)
	logrus.Info(versions[v])

}
