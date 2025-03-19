package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// 获取当前目录
	dir := "E:\\golandproject\\golangTest\\chapter_0（面试必备）\\项目面试"
	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("无法读取目录:", err)
		return
	}

	// 用于存储文件信息的结构体
	type fileInfo struct {
		originalName string
		number       int
		name         string
		hasExt       bool // 是否有扩展名
		ext          string
	}

	var fileList []fileInfo

	// 遍历文件，解析文件名
	for _, file := range files {
		if file.IsDir() {
			continue // 跳过目录
		}

		filename := file.Name()
		parts := strings.SplitN(filename, ".", 2) // 按第一个点分割
		if len(parts) < 1 {
			fmt.Printf("文件 %s 不符合命名规则，跳过\n", filename)
			continue
		}

		number, err := strconv.Atoi(parts[0]) // 解析序号
		if err != nil {
			fmt.Printf("文件 %s 的序号无效，跳过\n", filename)
			continue
		}

		// 解析文件名和扩展名
		var name, ext string
		var hasExt bool
		if len(parts) > 1 {
			remaining := parts[1]
			lastDotIndex := strings.LastIndex(remaining, ".") // 查找最后一个点
			if lastDotIndex >= 0 {
				name = remaining[:lastDotIndex]
				ext = remaining[lastDotIndex+1:]
				hasExt = true
			} else {
				name = remaining
				hasExt = false
			}
		}

		fileList = append(fileList, fileInfo{
			originalName: filename,
			number:       number,
			name:         name,
			hasExt:       hasExt,
			ext:          ext,
		})
	}

	// 按原始序号排序
	sort.Slice(fileList, func(i, j int) bool {
		return fileList[i].number < fileList[j].number
	})

	// 重新编号并重命名文件
	for i, file := range fileList {
		newNumber := i + 1
		var newName string
		if file.hasExt {
			newName = fmt.Sprintf("%d.%s.%s", newNumber, file.name, file.ext)
		} else {
			newName = fmt.Sprintf("%d.%s", newNumber, file.name)
		}

		if newName == file.originalName {
			continue // 如果文件名未改变，跳过
		}

		// 重命名文件
		err := os.Rename(filepath.Join(dir, file.originalName), filepath.Join(dir, newName))
		if err != nil {
			fmt.Printf("无法重命名文件 %s 为 %s: %v\n", file.originalName, newName, err)
		} else {
			fmt.Printf("重命名文件 %s 为 %s\n", file.originalName, newName)
		}
	}

	fmt.Println("文件序号整理完成！")
}
