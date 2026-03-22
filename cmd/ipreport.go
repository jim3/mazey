// ipreport.go
package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/jim3/mazey/internal/ipreport"
	"github.com/spf13/cobra"
)

var ipreportCmd = &cobra.Command{
	Use:   "ipreport <IP_ADDRESS>",
	Short: "Get an IP address report",
	Args:  cobra.ExactArgs(1),

	RunE: func(cmd *cobra.Command, args []string) error {
		ipaddr := args[0]

		danger := color.New(color.FgRed).Add(color.Bold)
		warning := color.New(color.FgYellow)
		success := color.New(color.FgGreen)
		info := color.New(color.FgCyan)
		dim := color.New(color.FgHiBlack).SprintFunc()

		report, err := ipreport.MergeReports(ipaddr, 3)
		if err != nil {
			return err
		}

		divider := dim("------------------------------------------------------------")
		fmt.Println(divider)
		info.Printf("IP: %s\nASN: %d\nCountry: %s\nRep: %d\nNetwork: %s\n",
			report.IP, report.ASN, report.Country, report.Reputation, report.Network)
		fmt.Println(divider)
		if report.Stats.Malicious > 0 {
			danger.Printf("MALICIOUS:  %d\n", report.Stats.Malicious)
		} else {
			success.Println("✅ MALICIOUS:  0 (Clean)")
		}
		if report.Stats.Suspicious > 0 {
			warning.Printf("SUSPICIOUS: %d\n", report.Stats.Suspicious)
		} else {
			success.Println("✅ SUSPICIOUS:  0 (Clean)")
		}
		fmt.Println(divider)
		info.Println("Recent domain names that have pointed to this IP:")
		for _, v := range report.Resolutions {
			info.Printf("- %s\n", v.HostName)
		}
		fmt.Println(divider)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(ipreportCmd)
	// Here you will define your flags and configuration settings.
}
