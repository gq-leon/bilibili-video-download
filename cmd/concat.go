package cmd

import (
	"bilibili-video-download/internal/concat"
	"bilibili-video-download/utils"
	"github.com/spf13/cobra"
	"log"
)

var concatCmd = &cobra.Command{
	Use:     "concat",
	Short:   "合并指定文件夹中的视频（根据创建时间）",
	Long:    "",
	Example: "bilibili concat -f ./video -o output.mp4",
	Run:     VideoConcat,
}

func VideoConcat(cmd *cobra.Command, args []string) {
	folder, _ := cmd.Flags().GetString("folder")
	output, _ := cmd.Flags().GetString("output")
	if folder == "" {
		log.Fatal("缺少文件夹参数")
	}

	if utils.PathExists(output) {
		log.Fatal(output, "文件已存在")
	}

	c := concat.New(folder, output, "video-concat.txt")
	if err := c.Start(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	concatCmd.Flags().StringP("folder", "f", "", "合并视频文件夹")
	concatCmd.Flags().StringP("output", "o", "output.mp4", "合并后文件名")

	RootCmd.AddCommand(concatCmd)
}
