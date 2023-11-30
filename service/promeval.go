// @Author songxiandong
// @Date 2023/11/29 15:57:00
// @Desc
package service

import "github.com/prometheus/client_golang/prometheus"

var (
	SSLxpiration = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "ssl_certificate_expiration",
			Help: "SSL certificate expiration in hours",
		},
		[]string{"url"},
	)
)

var (
	Domainxpiration = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "domain_expiration",
			Help: "domain expiration in hours",
		},
		[]string{"domain"},
	)
)
