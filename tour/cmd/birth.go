package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

const year = 1997

var birthCmd = &cobra.Command{
	Use:   "version",
	Short: "-",
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		version := time.Now().Year() - year

		if version > 10 && version < 100 {
			content = fmt.Sprintf("%d.%d.0", version/10%10, version/1%10)
		}

		fmt.Printf("4wen version %s\n", content)
		panic("no girlfriend panic~")
	},
}
