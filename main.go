// @Author sxd
// @Date 2023/11/29 09:58:00
// @Desc
package main

import (
	"fmt"
	"github.com/spf13/viper"
	"net/http"
	"sslexporter/service"
	"sslexporter/settings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// @Title init
// @Description
// @Author sxd 2023-11-29 14:42:26 ${time}
func init() {
	prometheus.MustRegister(service.SSLxpiration, service.Domainxpiration)
}

func main() {
	//1、加载配置文件

	if err := settings.Init(); err != nil {
		fmt.Println("init settings failed")
		return
	}

	testurl := viper.GetStringSlice("web.url")
	testdomain := viper.GetStringSlice("web.domain")

	go func() {
		for {
			for _, urlToChecks := range testurl {
				service.CheckSSLCertificateExpiration(urlToChecks)
			}
			for _, dominToChecks := range testdomain {
				service.CheckDomainExpiration(dominToChecks)
			}

			time.Sleep(1 * time.Hour) // Check every hour, adjust as needed
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("Starting server at :28080")
	http.ListenAndServe(":28080", nil)
}
