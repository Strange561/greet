# greet

一个基于 [Cobra](https://github.com/spf13/cobra) 的 Go 命令行工具。

## 安装

```bash
go install github.com/Strange561/greet@latest
```

安装后二进制文件位于 `$(go env GOPATH)/bin/greet`，请确保该目录已加入 `PATH`。

## 使用

```bash
greet --help          # 查看所有命令
greet hello           # 启动一个监听 :8080 的 HTTP 服务，访问 /hello 返回问候
greet admin -t <token> # 管理命令（需要 token）
```

## 本地开发

```bash
git clone https://github.com/Strange561/greet.git
cd greet
go build .
./greet --help
```

## 发布新版本

```bash
git tag v0.x.0
git push --tags
```

用户执行 `go install github.com/Strange561/greet@latest` 即可升级到最新版本。

## License

见 [LICENSE](LICENSE)。
