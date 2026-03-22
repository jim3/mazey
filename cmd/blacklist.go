package cmd

import (
	"fmt"
	"strconv"

	"github.com/jim3/mazey/internal/blacklist"
	"github.com/spf13/cobra"
)

var blacklistCmd = &cobra.Command{
	Use:   "blacklist [COUNT]",
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

		for i, v := range ipSlc {
			fmt.Printf("Blacklist IP# %d %s\n", i, v)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(blacklistCmd)
}
