package ngin

// SessionCurrentAccount ...
const SessionCurrentAccount = "current_account"

//可选参数列表
type Options struct {
	Layout                   string
	PageName                 string
	Template                 string
	Pjax                     bool
	PjaxLayout               string
	SessionCurrentAccountKey string
}

// newOptions 创建可选参数
func newOptions(opts ...Option) Options {
	opt := Options{
		Pjax:                     false,
		PjaxLayout:               "pjax_layout.tmpl",
		Layout:                   "layout.tmpl",
		Template:                 "default",
		SessionCurrentAccountKey: SessionCurrentAccount,
	}
	for _, o := range opts {
		o(&opt)
	}
	return opt
}

// Option 为可选参数赋值的函数
type Option func(*Options)

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
