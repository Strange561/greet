package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var adminToken string

// adminCmd 是管理命令入口
var adminCmd = &cobra.Command{
	Use:   "admin",
	Short: "Admin operations for greet",
	Long:  `Admin 子命令，用于执行管理类操作。使用 -t/--token 传入管理令牌。`,
	Run:   runAdmin,
}

func runAdmin(cmd *cobra.Command, args []string) {
	if adminToken == "" {
		fmt.Fprintln(os.Stderr, "error: admin token is required, use -t/--token")
		os.Exit(1)
	}
	// TODO: 在这里做真正的管理操作，比如校验 token、调用管理接口等
	fmt.Println("admin: authenticated, running admin tasks...")
}

func init() {
	rootCmd.AddCommand(adminCmd)
	adminCmd.Flags().StringVarP(&adminToken, "token", "t", "", "Admin token (required)")
}
