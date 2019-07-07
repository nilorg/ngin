package ngin

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
)

func loadTemplate(r *multitemplate.Render, templatesDir, template string, funcMap template.FuncMap) {
	// 加载布局
	layouts, err := filepath.Glob(filepath.Join(templatesDir, template, "layouts/*.tmpl"))
	if err != nil {
		panic(err)
	}
	// 加载错误页面
	errors, err := filepath.Glob(filepath.Join(templatesDir, template, "errors/*.tmpl"))
	if err != nil {
		panic(err)
	}
	for _, errPage := range errors {
		tmplName := fmt.Sprintf("%s_error_%s", template, filepath.Base(errPage))
		r.AddFromFilesFuncs(tmplName, funcMap, errPage)
	}

	// 加载局部页面
	partials, err := filepath.Glob(filepath.Join(templatesDir, template, "partials/*.tmpl"))
	if err != nil {
		panic(err)
	}

	// 页面文件夹
	pages, err := ioutil.ReadDir(filepath.Join(templatesDir, template, "pages"))
	if err != nil {
		panic(err)
	}
	for _, page := range pages {
		if !page.IsDir() {
			continue
		}
		for _, layout := range layouts {
			pageItems, err := filepath.Glob(filepath.Join(templatesDir, template, fmt.Sprintf("pages/%s/*.tmpl", page.Name())))
			if err != nil {
				panic(err)
			}
			files := []string{
				layout,
			}
			files = append(files, partials...)
			files = append(files, pageItems...)
			tmplName := fmt.Sprintf("%s/%s_pages/%s", template, filepath.Base(layout), page.Name())
			r.AddFromFilesFuncs(tmplName, funcMap, files...)
		}
	}
	// 加载单页面
	singles, err := filepath.Glob(filepath.Join(templatesDir, template, "singles/*.tmpl"))
	if err != nil {
		panic(err)
	}
	for _, singlePage := range singles {
		tmplName := fmt.Sprintf("%s/singles/%s", template, filepath.Base(singlePage))
		r.AddFromFilesFuncs(tmplName, funcMap, singlePage)
	}
}

// DefaultLoadTemplate ...
func DefaultLoadTemplate(templatesDir string, funcMap template.FuncMap, templates ...string) multitemplate.Render {
	r := multitemplate.New()
	for _, template := range templates {
		loadTemplate(&r, templatesDir, template, funcMap)
	}
	return r
}
