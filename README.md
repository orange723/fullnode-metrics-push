# fullnode-metrics-push

## 简介

`fullnode-metrics-push` 是一个使用 Go 语言编写的轻量级监控工具。它定期从指定的区块链全节点（Full Node）获取最新的区块高度，并将这些数据作为指标（Metrics）推送到兼容 Prometheus 的监控服务中。

## 功能特性

- **多链支持**: 可通过配置文件同时监控多个不同的区块链节点。
- **并发获取**: 使用 Go 的并发特性，能够高效地从多个 RPC 端点并行获取数据。
- **Prometheus 集成**: 将区块高度数据格式化并推送到任何兼容 Prometheus `remote_write` 协议的监控后端。
- **简单配置**: 使用 TOML 文件进行配置，简单明了。
- **命令行操作**: 提供简单的命令行界面来启动服务。

## 配置

您可以通过创建 `.toml` 配置文件来指定需要监控的节点和推送的目标地址。默认配置文件路径为 `config/config.example.toml`。

以下是一个配置文件的示例：

```toml
# config/config.example.toml

# 监控数据推送的目标服务器
[server]
host = "127.0.0.1"  # 监控服务器地址
port = 9090         # 监控服务器 remote_write 端口

# 需要监控的 RPC 节点列表
[[rpc]]
chain = "ethereum"  # 链名称，将作为 Prometheus 指标的名称
list = [
  "https://eth.public-node.com",
  "https://rpc.ankr.com/eth"
]

[[rpc]]
chain = "bsc"
list = [
  "https://bsc-dataseed1.binance.org",
  "https://bsc-dataseed2.defibit.io"
]
```

### 配置说明

- `[server]`:
  - `host`: 您的 Prometheus 或兼容服务的地址。
  - `port`: 服务的 `remote_write` 端口。
- `[[rpc]]`:
  - `chain`: 区块链的名称。这个名称将用作 Prometheus 指标的 `__name__`。
  - `list`: 一个包含该链多个 RPC 端点 URL 的列表。

## 使用方法

### 1. 构建

您可以使用 Go 编译器直接构建此项目：

```bash
go build -o fullnode-metrics-push .
```

### 2. 运行

使用 `-config` 标志来指定您的配置文件并运行程序。

```bash
./fullnode-metrics-push --config /path/to/your/config.toml
```

如果您不指定配置文件，程序将默认尝试加载 `config/config.example.toml`。

程序会读取配置文件，获取所有节点的区块高度，然后将指标推送到指定的服务器后退出。您可以配合 `crontab` 或其他调度工具来定期执行它。

例如，每分钟执行一次：

```bash
* * * * * /path/to/your/fullnode-metrics-push --config /path/to/your/config.toml
```

## 指标格式

程序会将区块高度数据推送为以下格式的 Prometheus 指标：

- **指标名称**: 您在配置文件中定义的 `chain` 名称（例如 `ethereum`、`bsc`）。
- **标签 (Labels)**:
  - `provider`: 节点的 RPC 地址。
- **值 (Value)**:
  - 对应节点的当前区块高度。

**Prometheus 中的查询示例:**

```promql
# 查询以太坊所有节点的区块高度
ethereum

# 查询 bsc 链特定节点的区块高度
bsc{provider="https-dataseed1.binance.org"}
```
