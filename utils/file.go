package utils

import (
	"fmt"
	"os"
)

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	return false
}

const (
	B  = 1
	KB = B * 1024
	MB = KB * 1024
	GB = MB * 1024
	TB = GB * 1024
)

// FormatFileSize 根据字节数转换文件大小
func FormatFileSize(fileSize float64) (string, string) {
	unit := ""
	value := fileSize

	switch {
	case fileSize >= TB:
		unit = "TB"
		value = value / TB
	case fileSize >= GB:
		unit = "GB"
		value = value / GB
	case fileSize >= MB:
		unit = "MB"
		value = value / MB
	case fileSize >= KB:
		unit = "KB"
		value = value / KB
	default:
		unit = "B"
	}

	return fmt.Sprintf("%.2f", value), unit
}
