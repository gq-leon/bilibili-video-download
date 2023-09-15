package cmd

import (
	"bilibili-video-download/internal/downloader"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var downloadCmd = &cobra.Command{
	Use:     "download",
	Short:   "根据指定BV下载视频",
	Long:    "",
	Example: "bilibili download -b xxx -a true -o ./video",
	Run: func(cmd *cobra.Command, args []string) {
		bv, _ := cmd.Flags().GetString("bvid")
		if bv == "" {
			log.Fatalln("缺少视频BV编号")
		}

		all, _ := cmd.Flags().GetBool("isAll")
		output, _ := cmd.Flags().GetString("output")
		if output == "" {
			output = "./" + bv
		}

		// 创建文件夹
		if err := os.MkdirAll(output, 0666); err != nil {
			log.Fatalln("文件夹创建失败:", err)
		}

		d := downloader.New(bv, output, all)
		if err := d.Start(); err != nil {
			log.Fatalln("视频下载失败:", err)
		}
	},
}

func init() {
	downloadCmd.Flags().StringP("output", "o", "", "文件存放目录")
	downloadCmd.Flags().StringP("bvid", "b", "", "视频BV编号")
	downloadCmd.Flags().BoolP("isAll", "a", false, "下载全部视频")

	RootCmd.AddCommand(downloadCmd)
}
