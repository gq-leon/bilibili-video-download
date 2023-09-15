package downloader

import (
	"bilibili-video-download/utils"
	"fmt"
)

type Collection struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    struct {
		View struct {
			Bvid      string `json:"bvid"`
			Aid       int    `json:"aid"`
			Videos    int    `json:"videos"`
			Tid       int    `json:"tid"`
			Tname     string `json:"tname"`
			Copyright int    `json:"copyright"`
			Pic       string `json:"pic"`
			Title     string `json:"title"`
			Pubdate   int    `json:"pubdate"`
			Ctime     int    `json:"ctime"`
			Desc      string `json:"desc"`
			UgcSeason struct {
				Id        int    `json:"id"`
				Title     string `json:"title"`
				Cover     string `json:"cover"`
				Mid       int    `json:"mid"`
				Intro     string `json:"intro"`
				SignState int    `json:"sign_state"`
				Attribute int    `json:"attribute"`
				Sections  []struct {
					SeasonId int    `json:"season_id"`
					Id       int    `json:"id"`
					Title    string `json:"title"`
					Type     int    `json:"type"`
					Episodes []struct {
						SeasonId  int    `json:"season_id"`
						SectionId int    `json:"section_id"`
						Id        int    `json:"id"`
						Aid       int    `json:"aid"`
						Cid       int    `json:"cid"`
						Title     string `json:"title"`
						Attribute int    `json:"attribute"`
						Arc       struct {
							Aid       int    `json:"aid"`
							Videos    int    `json:"videos"`
							TypeId    int    `json:"type_id"`
							TypeName  string `json:"type_name"`
							Copyright int    `json:"copyright"`
							Pic       string `json:"pic"`
							Title     string `json:"title"`
							Pubdate   int    `json:"pubdate"`
							Ctime     int    `json:"ctime"`
							Desc      string `json:"desc"`
							State     int    `json:"state"`
							Duration  int    `json:"duration"`
							Rights    struct {
								Bp            int `json:"bp"`
								Elec          int `json:"elec"`
								Download      int `json:"downloader"`
								Movie         int `json:"movie"`
								Pay           int `json:"pay"`
								Hd5           int `json:"hd5"`
								NoReprint     int `json:"no_reprint"`
								Autoplay      int `json:"autoplay"`
								UgcPay        int `json:"ugc_pay"`
								IsCooperation int `json:"is_cooperation"`
								UgcPayPreview int `json:"ugc_pay_preview"`
								ArcPay        int `json:"arc_pay"`
								FreeWatch     int `json:"free_watch"`
							} `json:"rights"`
							Author struct {
								Mid  int    `json:"mid"`
								Name string `json:"name"`
								Face string `json:"face"`
							} `json:"author"`
							Stat struct {
								Aid        int    `json:"aid"`
								View       int    `json:"view"`
								Danmaku    int    `json:"danmaku"`
								Reply      int    `json:"reply"`
								Fav        int    `json:"fav"`
								Coin       int    `json:"coin"`
								Share      int    `json:"share"`
								NowRank    int    `json:"now_rank"`
								HisRank    int    `json:"his_rank"`
								Like       int    `json:"like"`
								Dislike    int    `json:"dislike"`
								Evaluation string `json:"evaluation"`
								ArgueMsg   string `json:"argue_msg"`
								Vt         int    `json:"vt"`
								Vv         int    `json:"vv"`
							} `json:"stat"`
							Dynamic   string `json:"dynamic"`
							Dimension struct {
								Width  int `json:"width"`
								Height int `json:"height"`
								Rotate int `json:"rotate"`
							} `json:"dimension"`
							DescV2             interface{} `json:"desc_v2"`
							IsChargeableSeason bool        `json:"is_chargeable_season"`
							IsBlooper          bool        `json:"is_blooper"`
							EnableVt           int         `json:"enable_vt"`
							VtDisplay          string      `json:"vt_display"`
						} `json:"arc"`
						Page struct {
							Cid       int    `json:"cid"`
							Page      int    `json:"page"`
							From      string `json:"from"`
							Part      string `json:"part"`
							Duration  int    `json:"duration"`
							Vid       string `json:"vid"`
							Weblink   string `json:"weblink"`
							Dimension struct {
								Width  int `json:"width"`
								Height int `json:"height"`
								Rotate int `json:"rotate"`
							} `json:"dimension"`
						} `json:"page"`
						Bvid string `json:"bvid"`
					} `json:"episodes"`
				} `json:"sections"`
			} `json:"ugc_season"`
		} `json:"View"`
	} `json:"data"`
}

func (d *Downloader) GetVideoDetail(aid int) (*Collection, error) {
	var (
		url        = fmt.Sprintf("https://api.bilibili.com/x/web-interface/wbi/view/detail?aid=%d", aid)
		err        error
		client     *utils.Client
		collection *Collection
	)

	if client, err = utils.NewClient("GET", url, nil); err != nil {
		return nil, err
	}

	if err := client.SetHeader().SetCookie().SetReferer().Struct(&collection); err != nil {
		return nil, err
	}

	return collection, err
}
