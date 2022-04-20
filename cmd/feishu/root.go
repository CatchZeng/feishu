package feishu

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/CatchZeng/feishu/configs"
	"github.com/CatchZeng/feishu/pkg/feishu"
	"github.com/spf13/viper"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "feishu",
	Short: "feishu is a command line tool for feishu robot",
	Long:  "feishu is a command line tool for feishu robot",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func newClient() (*feishu.Client, error) {
	token := getAccessToken()
	secret := getSecret()

	if len(token) < 1 {
		return nil, errors.New("access_token can not be empty")
	}
	client := feishu.NewClient(token, secret)
	return client, nil
}

func getAccessToken() string {
	if len(accessToken) > 0 {
		return accessToken
	}

	value, err := configs.GetConfig(configs.AccessToken)
	if err == nil {
		return value
	}
	return ""
}

func getSecret() string {
	if len(secret) > 0 {
		return secret
	}

	value, err := configs.GetConfig(configs.Secret)
	if err == nil {
		return value
	}
	return ""
}

var accessToken, secret string
var debug bool

func init() {
	cobra.OnInitialize(configs.InitConfig)

	rootCmd.PersistentFlags().StringVarP(&accessToken, configs.AccessToken, "t", "", configs.AccessToken)
	rootCmd.PersistentFlags().StringVarP(&secret, configs.Secret, "s", "", configs.Secret)
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "D", false, "debug")

	if err := viper.BindPFlag(configs.AccessToken, rootCmd.PersistentFlags().Lookup(configs.AccessToken)); err != nil {
		log.Print(err)
	}
	if err := viper.BindPFlag(configs.Secret, rootCmd.PersistentFlags().Lookup(configs.Secret)); err != nil {
		log.Print(err)
	}
}
