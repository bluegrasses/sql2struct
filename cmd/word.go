package cmd

import (
	"log"
	"strings"

	"github.com/bluegrasses/sql2struct/internal/word"

	"github.com/spf13/cobra"
)

//定义参数变量
var str string
var mode int8

const (
	//ioto 默认起始为0
	ModeUpper  = iota + 1 //全部转大写
	ModelLower            //全部转小写
)

var desc = strings.Join([]string{
	"该子命令支持各种单词格式转化,模式如下:",
	"1:全部转大写",
	"2:全部转换为小写",
}, "\n")

// wordCmd represents the word command
var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  `支持多种单词格式转换`,
	Run: func(cmd *cobra.Command, args []string) {
		//设置命令具体处理逻辑
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(str)
		case ModelLower:
			content = word.ToLower(str)
		default:
			log.Fatalf("暂不支持该转换模式,请执行help word 查看帮助文档")
		}
		log.Printf("输出结果:%s", content)
	},
}

func init() {
	//设置flags
	wordCmd.Flags().StringVarP(&str, "str", "s", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换的模式")
}
