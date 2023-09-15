package downloader

import (
	"bilibili-video-download/internal/progressbar"
	"bilibili-video-download/utils"
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

type Downloader struct {
	bv     string
	all    bool
	output string
}

type VideoDownloadInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    struct {
		From              string   `json:"from"`
		Result            string   `json:"result"`
		Message           string   `json:"message"`
		Quality           int      `json:"quality"`
		Format            string   `json:"format"`
		Timelength        int      `json:"timelength"`
		AcceptFormat      string   `json:"accept_format"`
		AcceptDescription []string `json:"accept_description"`
		AcceptQuality     []int    `json:"accept_quality"`
		VideoCodecid      int      `json:"video_codecid"`
		SeekParam         string   `json:"seek_param"`
		SeekType          string   `json:"seek_type"`
		Durl              []struct {
			Order     int      `json:"order"`
			Length    int      `json:"length"`
			Size      int      `json:"size"`
			Ahead     string   `json:"ahead"`
			Vhead     string   `json:"vhead"`
			Url       string   `json:"url"`
			BackupUrl []string `json:"backup_url"`
		} `json:"durl"`
		SupportFormats []struct {
			Quality        int         `json:"quality"`
			Format         string      `json:"format"`
			NewDescription string      `json:"new_description"`
			DisplayDesc    string      `json:"display_desc"`
			Superscript    string      `json:"superscript"`
			Codecs         interface{} `json:"codecs"`
		} `json:"support_formats"`
		HighFormat   interface{} `json:"high_format"`
		LastPlayTime int         `json:"last_play_time"`
		LastPlayCid  int         `json:"last_play_cid"`
	} `json:"data"`
}

func (d *Downloader) Start() error {
	var (
		err       error
		videoBase *VideoBase
	)

	if videoBase, err = d.VideoInfo(); err != nil {
		return err
	}

	if d.all {
		collection, err := d.GetVideoDetail(videoBase.Data.Aid)
		if err != nil {
			return err
		}

		if len(collection.Data.View.UgcSeason.Sections) == 0 {
			return errors.New(fmt.Sprintf("%s 下没有视频集合", d.bv))
		}

		var (
			downloadCollection []struct {
				Name string
				Url  string
			}
			episodes = collection.Data.View.UgcSeason.Sections[0].Episodes
		)

		for _, episode := range episodes {
			downloadInfo, err := d.VideoDownloadInfo(episode.Bvid, episode.Cid)
			if err != nil {
				return err
			}

			downloadCollection = append(downloadCollection, struct {
				Name string
				Url  string
			}{Name: episode.Title, Url: downloadInfo.Data.Durl[0].Url})

			if err := d.download(episode.Title, downloadInfo.Data.Durl[0].Url); err != nil {
				return err
			}
		}
	} else {
		downloadInfo, err := d.VideoDownloadInfo(videoBase.Data.Bvid, videoBase.Data.Cid)
		if err != nil {
			return err
		}

		if err := d.download(videoBase.Data.Title, downloadInfo.Data.Durl[0].Url); err != nil {
			return err
		}
	}

	return nil
}

func New(bv, output string, all bool) *Downloader {
	return &Downloader{bv: bv, output: output, all: all}
}

func (d *Downloader) VideoDownloadInfo(bid string, cid int) (*VideoDownloadInfo, error) {
	var (
		err       error
		client    *utils.Client
		uri       = fmt.Sprintf("https://api.bilibili.com/x/player/playurl?bvid=%s&cid=%d&qn=0&fnval=0&fnver=0&fourk=1", bid, cid)
		videoInfo *VideoDownloadInfo
	)

	if client, err = utils.NewClient("GET", uri, nil); err != nil {
		return nil, err
	}

	if err = client.SetHeader().SetCookie().Struct(&videoInfo); err != nil {
		return nil, err
	}

	return videoInfo, nil
}

func (d *Downloader) download(name, uri string) error {
	parse, err := url.Parse(uri)
	path := parse.Path
	segments := strings.Split(path, "/")
	filename := fmt.Sprintf("%s/%s%s", d.output, name, filepath.Ext(segments[len(segments)-1]))

	//name += d.output + string(os.PathSeparator) + filepath.Ext(segments[len(segments)-1])
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	defer file.Close()

	video, err := utils.NewClient("GET", parse.String(), nil)
	if err != nil {
		return err
	}

	response, err := video.SetReferer().Do()
	if err != nil {
		return err
	}
	defer response.Body.Close()

	bar := progressbar.NewProgressBar(
		response.ContentLength,
		progressbar.ConfigSuffixWithOption(name),
	)

	if _, err := io.Copy(io.MultiWriter(file, bar), response.Body); err != nil {
		return err
	}

	return nil
}
