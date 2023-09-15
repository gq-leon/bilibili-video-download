package downloader

import (
	"bilibili-video-download/utils"
	"fmt"
)

type VideoBase struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Bvid  string `json:"bvid"`
		Aid   int    `json:"aid"`
		Cid   int    `json:"cid"`
		Title string `json:"title"`
	}
}

func (d *Downloader) VideoInfo() (*VideoBase, error) {
	var (
		err       error
		client    *utils.Client
		url       = fmt.Sprintf("https://api.bilibili.com/x/web-interface/view?bvid=%s", d.bv)
		videoBase = new(VideoBase)
	)

	if client, err = utils.NewClient("GET", url, nil); err != nil {
		return nil, err
	}

	if err = client.SetHeader().SetCookie().SetReferer().Struct(&videoBase); err != nil {
		return nil, err
	}

	return videoBase, nil
}
