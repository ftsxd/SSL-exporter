// @Author songxiandong
// @Date 2023/11/29 15:52:00
// @Desc
package service

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"
)

func CheckSSLCertificateExpiration(url string) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // Warning: InsecureSkipVerify should not be used in production
	}

	client := &http.Client{Transport: tr, Timeout: 10 * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("Error fetching %s: %s\n", url, err)
		return
	}
	defer resp.Body.Close()

	cert := resp.TLS.PeerCertificates[0]
	expiration := cert.NotAfter
	remainingTime := expiration.Sub(time.Now())
	fmt.Println(remainingTime)

	SSLxpiration.WithLabelValues(url).Set(remainingTime.Hours())
}
