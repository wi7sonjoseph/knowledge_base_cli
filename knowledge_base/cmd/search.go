package cmd

import (
	"context"
	"fmt"
	"knowledge_base_cli/service"
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search tool for the Knowledge Base CLI",
	Long: `Search tool for the Knowledge Base CLI`,
	Run: func(cmd *cobra.Command, args []string) {
		keyword, _ := cmd.Flags().GetString("keyword")
		
		//Testing config load
		fmt.Println("Development Mode:", viper.Get("is_local_development"))
		
		
		service, err := service.New(context.Background())
		if err == nil {
			service.Search(context.Background(), keyword)
		}
	},
}

func init() {
	searchCmd.Flags().StringP("keyword", "k", "", "Key Word Search")
	searchCmd.MarkFlagRequired("keyword")
	rootCmd.AddCommand(searchCmd)
}
