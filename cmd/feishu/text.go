package feishu

import (
	"github.com/CatchZeng/feishu"
	"github.com/go-ecosystem/utils/log"
	"github.com/spf13/cobra"
)

var textCmd = &cobra.Command{
	Use:   "text",
	Short: "send text message with feishu robot",
	Long:  `send text message with feishu robot`,
	Args:  cobra.MinimumNArgs(0),
	Run:   runTextCmd,
}

func runTextCmd(_ *cobra.Command, _ []string) {
	if len(textVars.text) < 1 {
		log.E("text can not be empty")
		return
	}

	client, err := newClient()
	if err != nil {
		log.E(err.Error())
		return
	}

	msg := feishu.NewTextMessage().
		SetText(textVars.text)
	if _, err := client.Send(msg); err != nil {
		log.E(err.Error())
	}
}

// TextVars struct
type TextVars struct {
	text string
}

var textVars TextVars

func init() {
	rootCmd.AddCommand(textCmd)
	textCmd.Flags().StringVarP(&textVars.text, "text", "e", "", "text")
}
