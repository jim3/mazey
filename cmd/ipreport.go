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
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ipaddr := args[0]
		// Setup color  palette
		header := color.New(color.FgHiCyan).Add(color.Bold)
		danger := color.New(color.FgRed).Add(color.Bold)
		warning := color.New(color.FgYellow)
		success := color.New(color.FgGreen)
		label := color.New(color.FgHiBlack)

		// API Call
		report := &ipreport.IpAddrReport{}
		res, err := report.GetIpReport(ipaddr)
		if err != nil {
			return err
		}

		stats := res.Data.Attributes.LastAnalysisStats

		fmt.Println()
		header.Printf("IP ADDRESS REPORT: %s\n", res.Data.Id)
		label.Println(" --------------------------------------------------")

		if stats.Malicious > 0 {
			danger.Printf("MALICIOUS:  %d\n", stats.Malicious)
		} else {
			success.Println(" ✅ MALICIOUS:  0 (Clean)")
		}

		// Warning for Suspicious Hits
		if stats.Suspicious > 0 {
			warning.Printf("SUSPICIOUS: %d\n", stats.Suspicious)
		}

		// Subtle details for the rest
		label.Printf("  • Undetected: %d\n", stats.Undetected)
		label.Printf("  • Harmless:   %d\n", stats.Harmless)
		label.Printf("  • Timeout:    %d\n", stats.Timeout)

		fmt.Println()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(ipreportCmd)
	// Here you will define your flags and configuration settings.
}
