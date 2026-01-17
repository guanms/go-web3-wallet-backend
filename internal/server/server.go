package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// Server 封装了 HTTP 引擎（gin.Engine）。
// 这个结构体负责持有应用的路由引擎并提供启动方法。
// 将来可以在此添加中间件、依赖注入的字段（例如 DB、缓存 客户端等）。
type Server struct {
	engine *gin.Engine
}

// NewServer 初始化并返回一个 *Server 实例。
// 工作流程：
// 1. 调用 loadConfig 加载配置（使用 viper）
// 2. 创建 gin 默认引擎（包含 Logger、Recovery 等中间件）
// 3. 调用 registerRoutes 将应用的路由注册到 gin 引擎上（registerRoutes 在项目其它位置定义）
// 4. 返回封装后的 Server 指针
//
// 注意：loadConfig 在读取配置失败时会触发 panic，适用于开发阶段快速失败；
// 生产环境可以改为返回错误并由调用方处理。
func NewServer() *Server {
	// 读取并加载配置文件到 viper 中
	loadConfig()

	// 创建带默认中间件（Logger, Recovery）的 gin 引擎
	engine := gin.Default()

	// 在工程中的其它文件定义 registerRoutes(engine) 来注册具体路由和处理函数
	registerRoutes(engine)

	return &Server{engine: engine}
}

// Run 启动 HTTP 服务并监听来自配置的端口。
//
// 它会从 viper 中读取配置键 `server.port`：
//   - 如果配置为空，则回退到默认端口 8080。
//   - 最终调用 gin 的 Run 方法启动服务器，Run 会阻塞当前 goroutine，
//     并返回可能的启动或运行时错误（例如端口被占用）。
func (s *Server) Run() error {
	port := viper.GetString("server.port")
	if port == "" {
		// 如果配置中没有指定端口，使用默认值
		port = "8080"
	}
	// gin.Engine.Run 接收形如 ":8080" 的地址字符串
	return s.engine.Run(fmt.Sprintf(":%s", port))
}

// loadConfig 使用 viper 加载 YAML 配置。
// 配置查找顺序与设置：
// - 配置文件名为 `config`（不带扩展名）
// - 配置类型为 `yaml`
// - 在相对路径 `./config` 下查找（常用于项目内的 config 目录）
// - 还会在当前工作目录 `.` 下查找（方便开发时将 config 放在项目根目录）
//
// 如果读取配置失败，会 panic 并打印错误信息。
// 在更严格的错误处理场景下，建议将该函数改为返回 error 并由调用方决定如何处理。
func loadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		// 这里选择直接 panic 以便尽早暴露配置问题
		panic("Failed to load config: " + err.Error())
	}
}
