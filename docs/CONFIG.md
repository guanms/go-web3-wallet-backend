# 配置说明

## 环境变量配置

为了安全起见，敏感信息（如私钥）应该通过环境变量配置，而不是直接写在配置文件中。

### 1. 复制配置文件模板

```bash
cp config.yaml.example config.yaml
cp .env.example .env
```

### 2. 设置环境变量

**方式一：使用 .env 文件（推荐本地开发）**

创建 `.env` 文件：

```bash
# 链私钥
CHAIN_PRIVATE_KEY=your_private_key_here

# JWT 密钥
JWT_SECRET=your_jwt_secret_here
```

**方式二：直接在终端设置环境变量**

```bash
export CHAIN_PRIVATE_KEY=your_private_key_here
export JWT_SECRET=your_jwt_secret_here
```

**方式三：在启动命令前设置**

```bash
CHAIN_PRIVATE_KEY=your_key JWT_SECRET=your_secret go run cmd/server/main.go
```

### 3. 加载 .env 文件

在项目中使用 godotenv 库自动加载 .env 文件：

```bash
go get github.com/joho/godotenv
```

然后在代码中：

```go
import _ "github.com/joho/godotenv/autoload"
```

## 配置文件说明

### config.yaml

```yaml
server:
  port: "8080"  # 服务端口

chain:
  rpc_url: "https://rpc.ankr.com/eth_sepolia"  # 以太坊 RPC 节点
  chain_id: 11155111  # 链 ID (Sepolia 测试网)
  # private_key 从环境变量 CHAIN_PRIVATE_KEY 读取

jwt:
  secret: "replace_with_your_own_secret"  # JWT 签名密钥
```

## 安全提示

1. **永远不要**将私钥提交到 Git 仓库
2. **永远不要**在生产环境中使用 .env 文件（应使用环境变量或密钥管理服务）
3. .env 文件已在 .gitignore 中，不会被提交
4. 定期轮换密钥和 JWT secret
5. 使用强密码和足够长的私钥

## 获取私钥

从 MetaMask 导出私钥：

1. 打开 MetaMask
2. 点击账户 → 账户详情
3. 点击 "导出私钥"
4. 输入密码，复制私钥

**⚠️ 警告：私钥一旦泄露，资产将无法挽回！**

## 测试网 Sepolia

本项目默认使用 Sepolia 测试网。获取测试币：

- https://sepoliafaucet.com/
- https://faucet.quicknode.com/ethereum/sepolia
- https://cloud.google.com/application/web3/faucet/ethereum/sepolia
