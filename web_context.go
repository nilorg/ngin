package ngin

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// CurrentAccount ...
const CurrentAccount = "current_account"

// WebContext Web上下文
type WebContext struct {
	*gin.Context
	layout, pageName, template string
}

// RenderPage 渲染页面
func (ctx *WebContext) RenderPage(data gin.H) {
	layout := ctx.layout
	if ctx.GetHeader("X-PJAX") == "true" {
		layout = "pjax_layout.tmpl"
	}
	tmplName := fmt.Sprintf("%s/%s_pages/%s", ctx.template, layout, ctx.pageName)
	if data == nil {
		data = gin.H{
			CurrentAccount: ctx.GetCurrentAccount(),
		}
	} else {
		data[CurrentAccount] = ctx.GetCurrentAccount()
	}
	ctx.HTML(http.StatusOK, tmplName, data)
}

// RenderSinglePage 渲染单页面
func (ctx *WebContext) RenderSinglePage(data gin.H) {
	tmplName := fmt.Sprintf("%s/singles/%s.tmpl", ctx.template, ctx.pageName)
	if data == nil {
		data = gin.H{
			CurrentAccount: ctx.GetCurrentAccount(),
		}
	} else {
		data[CurrentAccount] = ctx.GetCurrentAccount()
	}
	ctx.HTML(http.StatusOK, tmplName, data)
}

// SetCurrentAccount 设置当前账户
func (ctx *WebContext) SetCurrentAccount(data interface{}) error {
	session := sessions.Default(ctx.Context)
	session.Set(CurrentAccount, data)
	return session.Save()
}

// GetCurrentAccount 设置当前账户
func (ctx *WebContext) GetCurrentAccount() interface{} {
	session := sessions.Default(ctx.Context)
	return session.Get(CurrentAccount)
}

// DelCurrentAccount 删除当前账户
func (ctx *WebContext) DelCurrentAccount() error {
	session := sessions.Default(ctx.Context)
	session.Delete(CurrentAccount)
	return session.Save()
}

// WebControllerFunc Web控制器函数
func WebControllerFunc(ctlFunc func(ctx *WebContext), template, pageName string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tmplCtx := &WebContext{
			Context:  ctx,
			layout:   "layout.tmpl",
			template: template,
			pageName: pageName,
		}
		ctlFunc(tmplCtx)
	}
}

// WebControllerLayoutFunc Layout,Web控制器函数
func WebControllerLayoutFunc(ctlFunc func(ctx *WebContext), layout, template, pageName string) gin.HandlerFunc {
	if layout == "" {
		layout = "layout.tmpl"
	}
	return func(ctx *gin.Context) {
		tmplCtx := &WebContext{
			Context:  ctx,
			layout:   layout,
			template: template,
			pageName: pageName,
		}
		ctlFunc(tmplCtx)
	}
}
