package util

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

// GetIntSliceFromFile 从文件中读取[]int结构
// 用\n作为行分隔符, "splitString"作为列分隔符
func GetIntSliceFromFile(file, splitString string) ([]int, error) {
	s := make([]int, 0)
	f, err := os.Open(file)
	if err != nil {
		return s, err
	}
	defer f.Close()

	// 读取文件到buffer里边
	buf := bufio.NewReader(f)
	for {
		// 按照换行读取每一行
		l, err := buf.ReadString('\n')
		// 跳过空行
		if l == "\n" {
			continue
		}

		lineSplit := strings.SplitN(l, splitString, 1024)
		for _, v := range lineSplit {
			v = strings.TrimSpace(v)
			if v == "" {
				continue
			}
			value, _ := strconv.Atoi(v)
			s = append(s, value)
		}
		if err != nil {
			break
		}
	}
	return s, nil
}

func MakeName(name, version string) string {
	return fmt.Sprintf("%s/v%s/%s/%s", name, version, runtime.GOOS, runtime.Version())
}

func FileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil && os.IsNotExist(err) {
		return false
	}

	return true
}

func AbsolutePath(datadir string, filename string) string {
	if filepath.IsAbs(filename) {
		return filename
	}
	return filepath.Join(datadir, filename)
}

// 列出指定路径中的文件和目录
// 如果目录不存在，则返回空slice
func ScanDir(directory string) []string {
	file, err := os.Open(directory)
	if err != nil {
		return []string{}
	}
	names, err := file.Readdirnames(-1)
	if err != nil {
		return []string{}
	}
	return names
}

// 判断给定文件名是否是一个目录
// 如果文件名存在并且为目录则返回 true。如果 filename 是一个相对路径，则按照当前工作目录检查其相对路径。
func IsDir(filename string) bool {
	return isFileOrDir(filename, true)
}

// 判断给定文件名是否为一个正常的文件
// 如果文件存在且为正常的文件则返回 true
func IsFile(filename string) bool {
	return isFileOrDir(filename, false)
}

// 判断是文件还是目录，根据decideDir为true表示判断是否为目录；否则判断是否为文件
func isFileOrDir(filename string, decideDir bool) bool {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return false
	}
	isDir := fileInfo.IsDir()
	if decideDir {
		return isDir
	}
	return !isDir
}
