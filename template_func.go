package ngin

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
)

// DefaultLoadTemplate ...
// prefix 可以省略不写。默认情况下不需要写，一个gin多模板的时候需要设置prefix
func DefaultLoadTemplate(templatesDir, prefix string, funcMap template.FuncMap) multitemplate.Render {
	if prefix != "" {
		prefix += "_"
	}
	r := multitemplate.New()
	// 加载布局
	layouts, err := filepath.Glob(filepath.Join(templatesDir, "layouts/*.tmpl"))
	if err != nil {
		panic(err)
	}
	// 加载错误页面
	errors, err := filepath.Glob(filepath.Join(templatesDir, "errors/*.tmpl"))
	if err != nil {
		panic(err)
	}
	for _, errPage := range errors {
		tmplName := fmt.Sprintf("%serror_%s", prefix, filepath.Base(errPage))
		r.AddFromFilesFuncs(tmplName, funcMap, errPage)
	}

	// 加载局部页面
	partials, err := filepath.Glob(filepath.Join(templatesDir, "partials/*.tmpl"))
	if err != nil {
		panic(err)
	}

	// 页面文件夹
	pages, err := ioutil.ReadDir(filepath.Join(templatesDir, "pages"))
	if err != nil {
		panic(err)
	}
	for _, page := range pages {
		if !page.IsDir() {
			continue
		}
		for _, layout := range layouts {
			pageItems, err := filepath.Glob(filepath.Join(templatesDir, fmt.Sprintf("pages/%s/*.tmpl", page.Name())))
			if err != nil {
				panic(err)
			}
			files := []string{
				layout,
			}
			files = append(files, partials...)
			files = append(files, pageItems...)
			tmplName := fmt.Sprintf("%s%s_pages_%s", prefix, filepath.Base(layout), page.Name())
			r.AddFromFilesFuncs(tmplName, funcMap, files...)
		}
	}
	// 加载单页面
	singles, err := filepath.Glob(filepath.Join(templatesDir, "singles/*.tmpl"))
	if err != nil {
		panic(err)
	}
	for _, singlePage := range singles {
		tmplName := fmt.Sprintf("%ssingles_%s", prefix, filepath.Base(singlePage))
		r.AddFromFilesFuncs(tmplName, funcMap, singlePage)
	}
	return r
}
