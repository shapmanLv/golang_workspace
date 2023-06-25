package dfa

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

// 加载准备好的一份腾讯敏感词词库，有48603条数据
func wordGenerate() ([]string, error) {
	//打开文件
	file, err := os.Open("./testdata/tencent_sensitive_words1.txt")
	if err != nil {
		return nil, err
	}

	buffer := bufio.NewScanner(file)
	// 循环读取
	var result = make([]string, 0, 50000)
	for {
		if !buffer.Scan() {
			break //文件读完了,退出for
		}
		line := buffer.Text() //获取每一行
		result = append(result, line)
	}

	err = file.Close()
	if err != nil {
		return nil, err
	}

	return result, nil
}

// 加载准备好的一篇一万字的论文
func thesisGenerate() (string, error) {
	//打开文件
	file, err := os.Open("./testdata/thesis1.txt")
	if err != nil {
		return "", err
	}
	buffer := bufio.NewScanner(file)
	var result string
	for {
		if !buffer.Scan() {
			break //文件读完了,退出for
		}
		line := buffer.Text() //获取每一行
		strings.Trim(line, "\n")
		strings.Trim(line, " ")
		result += line
	}

	err = file.Close()
	if err != nil {
		return "", err
	}

	return result, nil
}

// 基准测试——多叉树的构建
func BenchmarkBuildTree(b *testing.B) {
	words, err := wordGenerate()
	if err != nil {
		panic(err)
	}
	b.ResetTimer() // 重置计时器，不把上面初始化参数的部分纳入基准测试
	b.ReportAllocs()
	BuildSearchTree(words)
}

// 基准测试——匹配
func BenchmarkDetection(b *testing.B) {
	target, err := thesisGenerate()
	if err != nil {
		panic(err)
	}
	words, err := wordGenerate()
	if err != nil {
		panic(err)
	}
	tree := BuildSearchTree(words)
	b.ResetTimer() // 重置计时器，不把上面初始化参数的部分纳入基准测试
	b.ReportAllocs()
	Detection(target, tree)
}
