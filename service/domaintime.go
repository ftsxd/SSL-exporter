// @Author songxiandong
// @Date 2023/11/29 16:00:00
// @Desc
package service

import (
	"fmt"
	"github.com/likexian/whois"
	"strings"
	"time"
)

// @Title CheckDomainExpiration
// @Description
// @Author sxd 2023-11-29 16:42:19 ${time}
// @Param domain
func CheckDomainExpiration(domain string) {
	whoisClient := whois.NewClient()
	result, err := whoisClient.Whois(domain)
	if err != nil {
		fmt.Printf("Error fetching WHOIS information: %s\n", err)
		return
	}
	expirationDate := extractExpirationDate(result)
	Domainxpiration.WithLabelValues(domain).Set(expirationDate)
}

// @Title extractExpirationDate
// @Description
// @Author sxd 2023-11-29 16:42:22 ${time}
// @Param whoisResult
// @Return float64
func extractExpirationDate(whoisResult string) float64 {
	var nilfloat float64
	var lines = splitLines(whoisResult)
	for _, line := range lines {
		if containsExpirationKeyword(line) {
			return OldDateFromLine(extractDateFromLine(line))
		}
	}
	return nilfloat
}

// @Title splitLines
// @Description
// @Author sxd 2023-11-29 16:42:27 ${time}
// @Param input
// @Return []string
func splitLines(input string) []string {
	return strings.Split(input, "\n")
}

// @Title containsExpirationKeyword
// @Description
// @Author sxd 2023-11-29 16:42:29 ${time}
// @Param line
// @Return bool
func containsExpirationKeyword(line string) bool {
	// You might need to modify this logic based on the WHOIS response structure for your specific domain registrar.
	keywords := []string{"Expiration Date", "expiry date", "Expiry", "expires on", "Registry Expiry Date"}
	for _, keyword := range keywords {
		if strings.Contains(strings.ToLower(line), strings.ToLower(keyword)) {
			return true
		}
	}
	return false
}

// @Title extractDateFromLine
// @Description
// @Author sxd 2023-11-29 16:42:32 ${time}
// @Param line
// @Return string
func extractDateFromLine(line string) string {
	// You might need to customize this date extraction logic based on the structure of the WHOIS response for your domain registrar.
	dateFormats := []string{"2006-01-02", "02-Jan-2006", "2006.01.02", "02-January-2006", "2006/01/02", "2006-01-02T15:04:05Z"}
	for _, format := range dateFormats {
		expirationDate, err := time.Parse(format, extractDate(line))
		if err == nil {
			return expirationDate.Format("2006-01-02T15:04:05Z")
		}
	}
	return "Unable to parse expiration date"
}

// @Title extractDate
// @Description
// @Author sxd 2023-11-29 16:42:36 ${time}
// @Param line
// @Return string
func extractDate(line string) string {
	parts := strings.Fields(line)
	if len(parts) > 1 {
		return parts[len(parts)-1]
		//fmt.Println(parts[len(parts)-1])
	}
	return ""
}

// @Title OldDateFromLine
// @Description
// @Author sxd 2023-11-29 16:42:38 ${time}
// @Param times
// @Return float64
func OldDateFromLine(times string) float64 {
	var nilfloat float64
	remainingTime, err := time.Parse(time.RFC3339, times)
	if err != nil {
		fmt.Println("解析时间出错:", err)
		return nilfloat
	}
	oldtime := remainingTime.Sub(time.Now())
	return oldtime.Hours()
}
