package ngin

import "net/http"

// SessionCurrentAccount ...
const SessionCurrentAccount = "current_account"

// Options 可选参数列表
type Options struct {
	StatusCode               int
	Layout                   string
	PageName                 string
	Template                 string
	Pjax                     bool
	PjaxLayout               string
	GlobalVariable           map[string]interface{}
	GlobalConstant           map[string]interface{}
	SessionCurrentAccountKey string
}

// newOptions 创建可选参数
func newOptions(opts ...Option) Options {
	opt := Options{
		StatusCode:               http.StatusOK,
		Pjax:                     false,
		PjaxLayout:               "pjax_layout.tmpl",
		Layout:                   "layout.tmpl",
		Template:                 "default",
		SessionCurrentAccountKey: SessionCurrentAccount,
		GlobalVariable:           map[string]interface{}{},
		GlobalConstant:           map[string]interface{}{},
	}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}

// Option 为可选参数赋值的函数
type Option func(*Options)

// StatusCode ...
func StatusCode(statusCode int) Option {
	return func(o *Options) {
		o.StatusCode = statusCode
	}
}

// Layout ...
func Layout(layout string) Option {
	return func(o *Options) {
		o.Layout = layout
	}
}

// PageName ...
func PageName(pageName string) Option {
	return func(o *Options) {
		o.PageName = pageName
	}
}

// Template ...
func Template(template string) Option {
	return func(o *Options) {
		o.Template = template
	}
}

// GlobalVariable ...
func GlobalVariable(variable map[string]interface{}) Option {
	return func(o *Options) {
		o.GlobalVariable = variable
	}
}

// GlobalConstant ...
func GlobalConstant(constant map[string]interface{}) Option {
	return func(o *Options) {
		o.GlobalConstant = constant
	}
}

// Pjax ...
func Pjax(pjax bool) Option {
	return func(o *Options) {
		o.Pjax = pjax
	}
}

// PjaxLayout ...
func PjaxLayout(pjaxLayout string) Option {
	return func(o *Options) {
		o.PjaxLayout = pjaxLayout
	}
}
