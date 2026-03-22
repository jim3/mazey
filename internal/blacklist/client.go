package blacklist

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type BlacklistResponse struct {
	BannedIPs []string `json:"blacklist"`
}

// GetBlacklist returns a pool of recently blacklisted IP's
func GetBlacklist(count int) ([]string, error) {
	var b BlacklistResponse
	url := os.Getenv("API_ENDPOINT")
	if url == "" {
		return nil, fmt.Errorf("API_ENDPOINT is not set")
	}

	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("blacklist API returned %d", res.StatusCode)
	}

	if err := json.NewDecoder(res.Body).Decode(&b); err != nil {
		return nil, fmt.Errorf("decode failed: %w", err)
	}

	if count > len(b.BannedIPs) {
		count = len(b.BannedIPs)
	}
	return b.BannedIPs[:count], nil

}

// ----------------------------------------------------

type IpLookUp struct {
	IP        string   `json:"ip"`
	Ports     []int    `json:"ports"`
	CPES      []string `json:"cpes"`
	HostNames []string `json:"hostnames"`
	Tags      []string `json:"tags"`
	Vulns     []string `json:"vulns"`
}

// Fast IP Lookups for Open Ports and Vulnerabilities
func LookupIP(ipAddr string) error {
	var ip IpLookUp
	URL := fmt.Sprintf("https://internetdb.shodan.io/%s", ipAddr)
	res, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading HostInfo body: %v", err)
	}

	err = json.Unmarshal(body, &ip)
	if err != nil {
		return fmt.Errorf("error unmarshalling json data: %v", err)
	}
	return nil
}
