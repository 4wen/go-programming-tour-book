package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var a int32
var b int32
var c int32

var triangleCmd = &cobra.Command{
	Use:   "triangle",
	Short: "给定三边计算是否符合三角形",
	Long:  "给定三边计算是否符合三角形",
	Run: func(cmd *cobra.Command, args []string) {
		if (a+b) < c || (a+c) < b || (b+c) < a {
			log.Println("不能组成三角形!请重新输入")
		} else if (a == b) && (a == c) {
			log.Println("等边三角形")
		} else if (a == b) || (a == c) || (b == c) {
			log.Println("等腰三角形")
		} else if a*a+b*b == c*c || a*a+c*c == b*b || b*b+c*c == a*a {
			log.Println("直角三角形")
		} else {
			log.Println("普通三角形")
		}

	},
}

func init() {
	triangleCmd.Flags().Int32VarP(&a, "a", "a", 0, "第一条边")
	triangleCmd.Flags().Int32VarP(&b, "b", "b", 0, "第二条边")
	triangleCmd.Flags().Int32VarP(&c, "c", "c", 0, "第三条边")
}
