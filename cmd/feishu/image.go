package feishu

import (
	"github.com/CatchZeng/feishu/pkg/feishu"
	"github.com/go-ecosystem/utils/log"
	"github.com/spf13/cobra"
)

var imageCmd = &cobra.Command{
	Use:   "image",
	Short: "send image message with feishu robot",
	Long:  `send image message with feishu robot`,
	Args:  cobra.MinimumNArgs(0),
	Run:   runImageCmd,
}

func runImageCmd(_ *cobra.Command, _ []string) {
	if len(imageVars.imageKey) < 1 {
		log.E("imageKey is empty")
		return
	}

	client, err := newClient()
	if err != nil {
		log.E(err.Error())
		return
	}

	msg := feishu.NewImageMessage().SetImageKey(imageVars.imageKey)

	req, _, err := client.Send(msg)
	if debug {
		log.I(req)
	}
	if err != nil {
		log.E(err.Error())
	}
}

type ImageVars struct {
	imageKey string
}

var imageVars ImageVars

func init() {
	rootCmd.AddCommand(imageCmd)
	imageCmd.Flags().StringVarP(&imageVars.imageKey, "imageKey", "m", "", "imageKey from https://open.feishu.cn/document/ukTMukTMukTM/uEDO04SM4QjLxgDN")
}
