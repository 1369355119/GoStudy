package gin

import (
	"github.com/gin-gonic/gin"
)

type Option func(engine *gin.Engine)

type optionsHolder struct {
	options []Option // 存储各个模块的路由配置
}

// Include 注册路由配置
func (oh *optionsHolder) Include(opts ...Option) {
	oh.options = append(oh.options, opts...)
}

// Init 初始化路由
func (oh *optionsHolder) Init(r *gin.Engine) {
	// 每个注册的模块都进行初始化
	for _, opt := range oh.options {
		opt(r)
	}
}
