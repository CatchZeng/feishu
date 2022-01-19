package feishu

import (
	"github.com/CatchZeng/feishu/pkg/feishu"
	"github.com/go-ecosystem/utils/log"
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
		log.E("shareChatID can not be empty")
		return
	}

	client, err := newClient()
	if err != nil {
		log.E(err.Error())
		return
	}

	msg := feishu.NewShareChatMessage().
		SetShareChatID(shareChatVars.shareChatID)
	req, _, err := client.Send(msg)
	if debug {
		log.I(req)
	}
	if err != nil {
		log.E(err.Error())
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
