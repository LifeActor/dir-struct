package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var (
	excludeDirs  = flag.String("exclude-dirs", "node_modules,.git", "逗号分隔的目录列表，排除这些目录")
	excludeFiles = flag.String("exclude-files", "package-lock.json", "逗号分隔的文件列表，排除这些文件")
	showContent  = flag.Bool("show-content", true, "是否显示源码文件的内容")
)

func printDirectoryStructure(path string, prefix string) {
	entries, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("读取目录错误:", err)
		return
	}

	for i, entry := range entries {
		fullPath := filepath.Join(path, entry.Name())
		if shouldExclude(fullPath) {
			continue
		}

		isLastEntry := i == len(entries)-1
		newPrefix := prefix
		if isLastEntry {
			fmt.Print(prefix, "└── ")
			newPrefix += "    "
		} else {
			fmt.Print(prefix, "├── ")
			newPrefix += "│   "
		}

		fmt.Println(entry.Name())

		if entry.IsDir() {
			newPath := filepath.Join(path, entry.Name())
			printDirectoryStructure(newPath, newPrefix)
		}
	}
}

func findSourceFiles(root string) []string {
	var sourceFiles []string

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println("访问路径错误:", path, err)
			return nil
		}

		if shouldExclude(path) {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		if !info.IsDir() {
			ext := filepath.Ext(path)
			if isSourceCodeFile(ext) {
				sourceFiles = append(sourceFiles, path)
			}
		}
		return nil
	})

	return sourceFiles
}

func shouldExclude(path string) bool {
	for _, dir := range strings.Split(*excludeDirs, ",") {
		if strings.Contains(path, filepath.Join(dir)) {
			return true
		}
	}
	for _, file := range strings.Split(*excludeFiles, ",") {
		if filepath.Base(path) == file {
			return true
		}
	}
	return false
}

func isSourceCodeFile(ext string) bool {
	sourceExtensions := []string{
		".js", ".go", ".py", ".php", ".html", ".css", ".json",
		".c", ".cpp", ".h", ".hpp", ".java", ".rb", ".sh", ".pl",
		".swift", ".kt", ".rs", ".scala", ".cs", ".ts", ".jsx", ".tsx",
		".vue", ".scss", ".less", ".md", ".yml", ".yaml", ".toml",
	}
	for _, sourceExt := range sourceExtensions {
		if ext == sourceExt {
			return true
		}
	}
	return false
}

func main() {
	flag.Parse()

	root := "."
	if len(flag.Args()) > 0 {
		root = flag.Arg(0)
	}

	fmt.Println(root)
	printDirectoryStructure(root, "")

	if *showContent {
		sourceFiles := findSourceFiles(root)
		for _, file := range sourceFiles {
			relPath, _ := filepath.Rel(root, file)
			fmt.Println("\n文件:", relPath, "的内容是:")

			content, err := os.ReadFile(file)
			if err != nil {
				fmt.Println("文件内容读取失败:", file, err)
				continue
			}
			fmt.Println(string(content))
		}
	}
}
