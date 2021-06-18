package feishu

import (
	"github.com/CatchZeng/feishu"
	"github.com/go-ecosystem/utils/log"
	"github.com/spf13/cobra"
)

var postCmd = &cobra.Command{
	Use:   "post",
	Short: "send post message with feishu robot",
	Long:  `send post message with feishu robot`,
	Args:  cobra.MinimumNArgs(0),
	Run:   runPostCmd,
}

func runPostCmd(_ *cobra.Command, _ []string) {
	client, err := newClient()
	if err != nil {
		log.E(err.Error())
		return
	}

	msg := feishu.NewPostMessage()

	if len(postVars.title) > 0 {
		msg.SetZHTitle(postVars.title)
	}

	line := []feishu.PostItem{}
	if len(postVars.text) > 0 {
		text := feishu.NewText(postVars.text)
		line = append(line, text)
	}
	if len(postVars.href) > 0 && len(postVars.hrefText) > 0 {
		a := feishu.NewA(postVars.hrefText, postVars.href)
		line = append(line, a)
	}
	if len(postVars.at) > 0 {
		at := feishu.NewAT(postVars.at)
		line = append(line, at)
	}
	if len(postVars.imageKey) > 0 {
		image := feishu.NewImage(postVars.imageKey, postVars.height, postVars.width)
		line = append(line, image)
	}
	msg.AppendZHContent(line)

	if _, err := client.Send(msg); err != nil {
		log.E(err.Error())
	}
}

type PostVars struct {
	title    string
	text     string
	href     string
	hrefText string
	at       string
	imageKey string
	height   int
	width    int
}

var postVars PostVars

func init() {
	rootCmd.AddCommand(postCmd)
	postCmd.Flags().StringVarP(&postVars.title, "title", "i", "", "title")
	postCmd.Flags().StringVarP(&postVars.text, "text", "e", "", "text")
	postCmd.Flags().StringVarP(&postVars.href, "href", "r", "", "href")
	postCmd.Flags().StringVarP(&postVars.hrefText, "hrefText", "f", "", "href text")
	postCmd.Flags().StringVarP(&postVars.at, "at", "a", "", "at user_id, if you want @all, just put `all`")
	postCmd.Flags().StringVarP(&postVars.imageKey, "imageKey", "m", "", "imageKey from https://open.feishu.cn/document/ukTMukTMukTM/uEDO04SM4QjLxgDN")
	postCmd.Flags().IntVarP(&postVars.height, "height", "g", 300, "image height")
	postCmd.Flags().IntVarP(&postVars.width, "width", "w", 300, "image width")
}
