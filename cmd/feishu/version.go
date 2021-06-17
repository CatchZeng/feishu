package feishu

import (
	"log"

	v "github.com/go-ecosystem/utils/version"
	"github.com/spf13/cobra"
)

const (
	version   = "0.0.1"
	buildTime = "2021/06/17"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "feishu version",
	Long:  `feishu version`,
	Run:   runVersionCmd,
}

func runVersionCmd(_ *cobra.Command, _ []string) {
	v := v.Stringify(version, buildTime)
	log.Println(v)
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
