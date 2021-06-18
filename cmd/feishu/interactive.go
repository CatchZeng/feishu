package feishu

import (
	"github.com/CatchZeng/feishu"
	"github.com/go-ecosystem/utils/log"
	"github.com/spf13/cobra"
)

var interactiveCmd = &cobra.Command{
	Use:   "interactive",
	Short: "send interactive message with feishu robot",
	Long:  `send interactive message with feishu robot`,
	Args:  cobra.MinimumNArgs(0),
	Run:   runInteractiveCmd,
}

func runInteractiveCmd(_ *cobra.Command, _ []string) {
	if len(interactiveVars.card) < 1 {
		log.E("card is empty")
		return
	}

	client, err := newClient()
	if err != nil {
		log.E(err.Error())
		return
	}

	msg := feishu.NewInteractiveMessage().SetCard(interactiveVars.card)

	if _, err := client.Send(msg); err != nil {
		log.E(err.Error())
	}
}

type InteractiveVars struct {
	card string
}

var interactiveVars InteractiveVars

func init() {
	rootCmd.AddCommand(interactiveCmd)
	interactiveCmd.Flags().StringVarP(&interactiveVars.card, "card", "c", "", "card from cardbuilder https://open.feishu.cn/tool/cardbuilder?from=custom_bot_doc")
}
