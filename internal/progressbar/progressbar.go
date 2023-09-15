package progressbar

import (
	"bilibili-video-download/utils"
	"fmt"
	"strings"
	"sync"
)

type Option func(p *ProgressBar)

type ProgressBar struct {
	state  state
	config config
	lock   sync.Mutex
}

type state struct {
	maxBytes       float64
	currentBytes   float64
	currentPercent float64
}

type config struct {
	backgroundLayer string
	maskLayer       string
	prefix          string
	suffix          string
}

func NewProgressBar(maxBytes int64, options ...Option) *ProgressBar {
	progressBar := &ProgressBar{
		state: state{
			maxBytes: float64(maxBytes),
		},
	}

	defaultOptions := []Option{DefaultConfigWithOption()}
	options = append(defaultOptions, options...)

	for _, option := range options {
		option(progressBar)
	}

	return progressBar
}

func ConfigSuffixWithOption(suffix string) Option {
	return func(p *ProgressBar) {
		p.config.suffix = suffix
	}
}

func DefaultConfigWithOption() Option {
	return func(p *ProgressBar) {
		p.config = config{
			backgroundLayer: "▒",
			maskLayer:       "█",
		}
	}
}

func (bar *ProgressBar) Write(p []byte) (n int, err error) {
	n = len(p)
	bar.Add(n)
	return
}

func (bar *ProgressBar) Add(n int) {
	bar.lock.Lock()
	defer bar.lock.Unlock()
	bar.state.currentBytes += float64(n)
	bar.state.currentPercent = bar.currentPercent()
	bar.render()
}

func (bar *ProgressBar) render() {
	output := "\r\u001B[K\u001B[34m"

	if bar.config.prefix != "" {
		output = fmt.Sprintf("%s %s", output, bar.config.prefix)
	}

	if maskLayer := bar.MaskLayer(); maskLayer != "" {
		output = fmt.Sprintf("%s %-50s", output, maskLayer)
	}

	// 完成进度
	output = fmt.Sprintf("%s %.2f%s", output, bar.state.currentPercent, "%")

	// 文件大小
	output = fmt.Sprintf("%s %s", output, bar.showSize())

	if bar.config.suffix != "" {
		output = fmt.Sprintf("%s %s", output, bar.config.suffix)
	}

	fmt.Print(output)
}

// 已完成进度比例(尺寸)
func (bar *ProgressBar) showSize() string {
	currentSize, _ := utils.FormatFileSize(bar.state.currentBytes)
	maxSize, unit := utils.FormatFileSize(bar.state.maxBytes)
	return fmt.Sprintf("%s/%s %s", currentSize, maxSize, unit)
}

// MaskLayer 动态进度条
func (bar *ProgressBar) MaskLayer() string {
	mask := strings.Repeat(bar.config.maskLayer, int(bar.state.currentPercent/2))
	background := strings.Repeat(bar.config.backgroundLayer, int(50-(bar.state.currentPercent/2)))
	return fmt.Sprintf("%s%s", mask, background)
}

// 当前进度百分比
func (bar *ProgressBar) currentPercent() float64 {
	return bar.state.currentBytes / bar.state.maxBytes * 100
}
