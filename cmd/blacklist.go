package cmd

import (
	"fmt"
	"strconv"

	"github.com/jim3/mazey/internal/blacklist"
	"github.com/spf13/cobra"
)

var blacklistCmd = &cobra.Command{
	Use:   "blacklist [count]",
	Short: "Fetches blacklisted IPs and and enriches them using various threat intelligence API's",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		count := 10
		if len(args) > 0 {
			parsedCount, err := strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("invalid count %q: %w", args[0], err)
			}
			count = parsedCount
		}

		client := &blacklist.BlacklistResponse{}
		ipSlc, err := client.GetBlacklist(count)
		if err != nil {
			return err
		}
		for _, v := range ipSlc {
			var resp blacklist.IpLookUp
			err := resp.LookupIP(v)
			if err != nil {
				fmt.Fprintf(cmd.ErrOrStderr(), "IP lookup failed for %s: %v\n", v, err)
				continue
			}
			fmt.Println("---------------------------------------------")
			fmt.Println("Looking up blacklisted ip address: ", v)
			fmt.Printf("CPEs: %v\n", resp.CPES)
			fmt.Printf("Hostname: %v\n", resp.HostNames)
			fmt.Printf("IP: %v\n", resp.IP)
			fmt.Printf("Ports: %v\n", resp.Ports)
			fmt.Printf("Tags: %v\n", resp.Tags)
			fmt.Printf("Vulns: %v\n", resp.Vulns)
			fmt.Println()
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(blacklistCmd)
}
