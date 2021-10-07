package main

import (
	"bytes"
	_ "github.com/seanlan/lazy/init"
	"go.uber.org/zap"
	"html/template"
)

func main() {

	var buff bytes.Buffer
	tmp := `{{if result}} 111 {{else}} 222 {{end}}`
	data := map[string]interface{}{"result": true}
	tmpl, _ := template.New("name").Parse(tmp)
	tmpl.Execute(&buff, data)
	zap.S().Info(buff.String())

}
