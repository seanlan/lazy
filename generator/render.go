package generator

import (
	"bufio"
	"bytes"
	"github.com/gobuffalo/packr/v2"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"text/template"
)

func render(outPath, tmplDir, tmplName string, data interface{}, cover bool) {
	var err error
	var buff bytes.Buffer
	var content string
	var tmpl *template.Template
	if tmplDir != "" {
		tmpl, err = template.ParseFiles(filepath.Join(tmplDir, tmplName))
		if err != nil {
			zap.S().Info(err)
			return
		}
	} else {
		content, err := packr.New("lazy", "../templates").FindString(tmplName)
		if err != nil {
			zap.S().Info(err)
			return
		}
		tmpl, err = template.New(tmplName).Parse(content)
		if err != nil {
			zap.S().Info(err)
			return
		}
	}
	err = tmpl.Execute(&buff, data)
	if err != nil {
		zap.S().Info(err)
		return
	}
	var fs *os.File
	defer fs.Close()
	_, err = os.Stat(outPath)
	content = buff.String()
	if err != nil { // file not exist
		fs, err = os.Create(outPath) //create file
	} else { // file exist
		if !cover {
			fs, err = os.OpenFile(outPath, os.O_WRONLY|os.O_APPEND, 0666)
		} else {
			fs, err = os.OpenFile(outPath, os.O_WRONLY|os.O_TRUNC, 0666)
		}
	}
	if err != nil {
		zap.S().Info(err)
		return
	}
	writer := bufio.NewWriter(fs)
	writer.WriteString(content)
	writer.Flush()
}
