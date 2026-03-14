package ipreport

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Get an IP address report
type IpAddrReport struct {
	Data struct {
		Id         string `json:"id"`
		Type       string `json:"type"`
		Attributes struct {
			WhoIs        string `json:"whois"`
			LastAnalysis int    `json:"last_analysis_date"`
			TotalVotes   struct {
				Harmless  int `json:"harmless"`
				Malicious int `json:"malicious"`
			} `json:"total_votes"`
			LastAnalysisStats struct {
				Malicious  int `json:"malicious"`
				Suspicious int `json:"suspicious"`
				Undetected int `json:"undetected"`
				Harmless   int `json:"harmless"`
				Timeout    int `json:"timeout"`
			} `json:"last_analysis_stats"`
		} `json:"attributes"`
	} `json:"data"`
}

// GetIpReport fetches an IP address report from VirusTotal.
func (ip IpAddrReport) GetIpReport(ipaddr string) (IpAddrReport, error) {
	if ipaddr == "" {
		return ip, fmt.Errorf("hash cannot be empty")
	}
	url := fmt.Sprintf("https://www.virustotal.com/api/v3/ip_addresses/%s", ipaddr)
	apiKey := os.Getenv("VT_API_KEY")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return IpAddrReport{}, fmt.Errorf("malformed URL, check for misspellings %s", err)
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("x-apikey", apiKey)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client request failed: %s\n", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return ip, fmt.Errorf("IP report API returned %d", res.StatusCode)
	}

	if err := json.NewDecoder(res.Body).Decode(&ip); err != nil {
		return ip, fmt.Errorf("decode failed: %w", err)
	}

	return ip, nil
}
