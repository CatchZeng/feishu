package feishu

import (
	"log"

	"github.com/CatchZeng/feishu/pkg/feishu"
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
		log.Fatal("card is empty")
		return
	}

	client, err := newClient()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	msg := feishu.NewInteractiveMessage().SetCard(interactiveVars.card)

	req, _, err := client.Send(msg)
	if debug {
		log.Print(req)
	}
	if err != nil {
		log.Fatal(err.Error())
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
