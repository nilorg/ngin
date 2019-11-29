package ngin

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// WebContext Web上下文
type WebContext struct {
	*gin.Context
	opts Options
}

func (ctx *WebContext) loadOptions(opts ...Option) {
	for _, o := range opts {
		o(&ctx.opts)
	}
}

// RenderPage 渲染页面
func (ctx *WebContext) RenderPage(data gin.H, opts ...Option) {
	ctx.loadOptions(opts...)
	layout := ctx.opts.Layout
	if ctx.opts.Pjax && ctx.GetHeader("X-PJAX") == "true" {
		layout = ctx.opts.PjaxLayout
	}
	pageName := ctx.opts.PageName
	tmplName := fmt.Sprintf("%s/%s_pages/%s", ctx.opts.Template, layout, pageName)
	fmt.Printf("tmplName: %s\n", tmplName)
	if data == nil {
		data = gin.H{}
	}
	data[ctx.opts.SessionCurrentAccountKey] = ctx.GetCurrentAccount()
	data["constant"] = ctx.opts.GlobalConstant
	data["variable"] = ctx.opts.GlobalVariable
	ctx.HTML(http.StatusOK, tmplName, data)
}

// RenderSinglePage 渲染单页面
func (ctx *WebContext) RenderSinglePage(data gin.H, opts ...Option) {
	ctx.loadOptions(opts...)
	tmplName := fmt.Sprintf("%s/singles/%s.tmpl", ctx.opts.Template, ctx.opts.PageName)
	if data == nil {
		data = gin.H{}
	}
	data[ctx.opts.SessionCurrentAccountKey] = ctx.GetCurrentAccount()
	data["constant"] = ctx.opts.GlobalConstant
	data["variable"] = ctx.opts.GlobalVariable
	ctx.HTML(http.StatusOK, tmplName, data)
}

// SetCurrentAccount 设置当前账户
func (ctx *WebContext) SetCurrentAccount(data interface{}) error {
	session := sessions.Default(ctx.Context)
	session.Set(ctx.opts.SessionCurrentAccountKey, data)
	return session.Save()
}

// GetCurrentAccount 设置当前账户
func (ctx *WebContext) GetCurrentAccount() interface{} {
	session := sessions.Default(ctx.Context)
	return session.Get(ctx.opts.SessionCurrentAccountKey)
}

// DelCurrentAccount 删除当前账户
func (ctx *WebContext) DelCurrentAccount() error {
	session := sessions.Default(ctx.Context)
	session.Delete(ctx.opts.SessionCurrentAccountKey)
	return session.Save()
}

// NewWebControllerFunc Web控制器函数
func NewWebControllerFunc(ctlFunc func(ctx *WebContext), opts ...Option) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tmplCtx := &WebContext{
			Context: ctx,
			opts:    newOptions(opts...),
		}
		ctlFunc(tmplCtx)
	}
}
