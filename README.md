# 目录结构打印工具

这是一个用 Go 语言编写的工具，用于打印目录结构并显示常见源码文件的内容。该工具支持排除特定目录和文件，并可以选择是否显示文件内容。

## 功能特性

- 递归打印目录结构。
- 排除特定目录和文件。
- 显示常见源码文件的内容。
- 支持多种源码文件后缀。

## 安装

### 前提条件

- 安装 [Go 语言环境](https://golang.org/dl/)（版本 1.16 或更高）。

### 编译和安装

1. 克隆仓库或下载源码：
    ```sh
    git clone https://github.com/LifeActor/dir-struct.git
    cd dir-struct
    ```

2. 编译项目：
    ```sh
    go build -o dir-struct
    ```

3. 将编译后的可执行文件移动到系统 PATH 中的目录（可选）：
    ```sh
    sudo mv dir-struct /usr/local/bin/
    ```

## 使用方法

### 命令行参数

- `-exclude-dirs`：逗号分隔的目录列表，排除这些目录。默认值为 `node_modules,.git`。
- `-exclude-files`：逗号分隔的文件列表，排除这些文件。默认值为 `package-lock.json`。
- `-show-content`：是否显示源码文件的内容。默认值为 `true`。

### 示例

1. 打印当前目录的结构，并显示源码文件内容：
    ```sh
    dir-struct
    ```

2. 打印指定目录的结构，并排除 `node_modules` 和 `.git` 目录：
    ```sh
    dir-struct -exclude-dirs=node_modules,.git /path/to/directory
    ```

3. 打印指定目录的结构，不显示源码文件内容：
    ```sh
    dir-struct -show-content=false /path/to/directory
    ```

## 支持的源码文件后缀

该工具支持以下源码文件后缀：

- `.js`, `.go`, `.py`, `.php`, `.html`, `.css`, `.json`
- `.c`, `.cpp`, `.h`, `.hpp`, `.java`, `.rb`, `.sh`, `.pl`
- `.swift`, `.kt`, `.rs`, `.scala`, `.cs`, `.ts`, `.jsx`, `.tsx`
- `.vue`, `.scss`, `.less`, `.md`, `.yml`, `.yaml`, `.toml`

## 贡献

欢迎贡献代码和提出问题。请遵循以下步骤：

1.  Fork 仓库。
2.  创建新的分支 (`git checkout -b feature/your-feature`）。
3.  提交更改 (`git commit -am '添加新功能'`）。
4.  推送到分支 (`git push origin feature/your-feature`）。
5.  创建 Pull Request。

## 许可证

该项目采用 MIT 许可证。有关详细信息，请参阅 [LICENSE](LICENSE) 文件。

## 联系

如有任何问题或建议，请通过 [email](mailto:9349540@qq.com) 或 [GitHub Issues](https://github.com/LifeActor/dir-struct/issues) 联系我们。

---
