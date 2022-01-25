package feishu

import (
	"log"

	"github.com/CatchZeng/feishu/pkg/feishu"
	"github.com/spf13/cobra"
)

var shareChatCmd = &cobra.Command{
	Use:   "shareChat",
	Short: "send shareChat message with feishu robot",
	Long:  `send shareChat message with feishu robot`,
	Args:  cobra.MinimumNArgs(0),
	Run:   runShareChatCmd,
}

func runShareChatCmd(_ *cobra.Command, _ []string) {
	if len(shareChatVars.shareChatID) < 1 {
		log.Fatal("shareChatID can not be empty")
		return
	}

	client, err := newClient()
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	msg := feishu.NewShareChatMessage().
		SetShareChatID(shareChatVars.shareChatID)
	req, _, err := client.Send(msg)
	if debug {
		log.Print(req)
	}
	if err != nil {
		log.Fatal(err.Error())
	}
}

// ShareChatVars struct
type ShareChatVars struct {
	shareChatID string
}

var shareChatVars ShareChatVars

func init() {
	rootCmd.AddCommand(shareChatCmd)
	shareChatCmd.Flags().StringVarP(&shareChatVars.shareChatID, "shareChatID", "a", "", "shareChatID, see more https://open.feishu.cn/document/ukTMukTMukTM/ucjMxEjL3ITMx4yNyETM?op_tracking=hc")
}
