package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	INPUT_FILE  = "d.txt"
	OUTPUT_FILE = "Last3_And_Last1"
)

func main() {
	var filename string

	if len(os.Args) == 2 {
		filename = os.Args[1]
	}

	if filename == "" {
		filename = INPUT_FILE
	}

	fmt.Println("数据文件名:", filename)

	// 打开文件
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return
	}
	defer file.Close()

	//OUTPUT_FILE 加上时间后缀
	output := OUTPUT_FILE + "_" + time.Now().Format("2006-01-02-15-04-05") + ".txt"

	// 创建一个新文件，用于保存提取到的数据
	f, err := os.OpenFile(output, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	// 创建一个 Scanner 来逐行读取文件
	scanner := bufio.NewScanner(file)
	const maxTokenSize = 10 * 1024 * 1024 // Set your desired maximum token size
	buf := make([]byte, maxTokenSize)
	scanner.Buffer(buf, maxTokenSize)

	// 计数器，用于跟踪当前行号
	lineNumber := 1

	// 逐行扫描文件
	for scanner.Scan() {
		// 如果行号大于等于18，则提取内容
		if lineNumber >= 18 {
			line := scanner.Text()
			// 用空格分割每行数据
			parts := strings.Fields(line)

			// 如果至少有三列数据
			if len(parts) >= 3 {
				// 提取倒数第一列和倒数第三列
				lastColumn := parts[len(parts)-1]
				thirdToLastColumn := parts[len(parts)-3]

				// 输出提取到的数据
				// fmt.Printf("倒数第一列: %s, 倒数第三列: %s\n", lastColumn, thirdToLastColumn)

				// 将提取到的数据写入文件
				if _, err = f.WriteString(thirdToLastColumn + " " + lastColumn + "\n"); err != nil {
					fmt.Println(err)
					return
				}

			}
		}

		// 增加行号计数
		lineNumber++
	}

	//保存文件
	f.Sync()

	fmt.Printf("数据提取完成: %s", OUTPUT_FILE)

	// 检查扫描过程是否出错
	if err := scanner.Err(); err != nil {
		fmt.Println("文件扫描错误:", err)
	}
}
