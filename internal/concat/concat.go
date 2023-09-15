package concat

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
	"syscall"
)

var (
	ErrFolderEmpty = errors.New("该文件夹下暂无视频")
)

type Concat struct {
	folder     string
	output     string
	concatFile string
	entries    []os.DirEntry
}

func New(folder, output, concatFile string) *Concat {
	return &Concat{
		folder:     folder,
		output:     output,
		concatFile: concatFile,
	}
}

func (c *Concat) Start() error {
	log.Println("检查文件夹合法性")
	if err := c.checkFolder(); err != nil {
		return err
	}

	log.Println("创建ffmpeg规则文件")
	if err := c.ruleFile(".mp4"); err != nil {
		return errors.New(fmt.Sprintf("规则文件创建失败:%s", err))
	}

	if err := c.videoConcat(c.output); err != nil {
		return errors.New(fmt.Sprintf("视频合并失败:%s", err))
	}
	log.Println("视频合并成功,文件名:", c.output)

	//log.Println("删除ffmpeg规则文件")
	//if err := os.Remove(c.concatFile); err != nil {
	//	return errors.New(fmt.Sprintf("规则文件删除失败:%s", err))
	//}
	return nil
}

func (c *Concat) checkFolder() (err error) {
	if c.entries, err = os.ReadDir(c.folder); err != nil {
		return
	}

	if len(c.entries) == 0 {
		return ErrFolderEmpty
	}

	return
}

func (c *Concat) ruleFile(suffix string) error {
	var (
		err        error
		collection Videos
	)

	for _, entry := range c.entries {
		if filepath.Ext(entry.Name()) == suffix {
			fileInfo, _ := os.Stat(c.folder + string(os.PathSeparator) + entry.Name())
			attribute := fileInfo.Sys().(*syscall.Win32FileAttributeData)
			collection = append(collection, Video{Name: entry.Name(), CreateTime: attribute.CreationTime.Nanoseconds()})
		}
	}

	if len(collection) == 0 {
		return ErrFolderEmpty
	}

	file, _ := os.OpenFile(c.concatFile, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	defer file.Close()
	_ = os.Truncate(c.concatFile, 0)

	sort.Sort(collection)
	for _, video := range collection {
		_, _ = file.Write([]byte(fmt.Sprintf("file '%s'", c.folder+string(os.PathSeparator)+video.Name)))
		_, _ = file.WriteString("\n")
	}

	return err
}

func (c *Concat) videoConcat(output string) error {
	cmdStr := fmt.Sprintf("ffmpeg -f concat -safe 0 -i %s -c copy %s", c.concatFile, output)
	log.Println("执行视频合并指令:", cmdStr)
	cmdArg := strings.Split(cmdStr, " ")
	command := exec.Command(cmdArg[0], cmdArg[1:]...)
	//command.Stderr = log.Writer()
	return command.Run()
}
