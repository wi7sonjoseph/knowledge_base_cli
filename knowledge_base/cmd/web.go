package cmd

import (
	"context"
	"fmt"
	"knowledge_base_cli/service"
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
)

// webCmd represents the web command
var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Web tool for the Knowledge Base CLI",
	Long: `Web tool for the Knowledge Base CLI`,
	Run: func(cmd *cobra.Command, args []string) {
		keyword, _ := cmd.Flags().GetString("keyword")
		
		//Testing config load
		fmt.Println("Development Mode:", viper.Get("is_local_development"))
		
		
		service, err := service.New(context.Background())
		if err == nil {
			service.Web(context.Background(), keyword)
		}
	},
}

func init() {
	webCmd.Flags().StringP("keyword", "k", "", "Key Word Search")
	webCmd.MarkFlagRequired("keyword")
	rootCmd.AddCommand(webCmd)
}
