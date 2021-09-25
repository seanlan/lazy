package generator

import (
	"bufio"
	"bytes"
	"go.uber.org/zap"
	"os"
	"text/template"
)

func render(outPath, tmplPaht string, data interface{}, cover bool) {
	var buff bytes.Buffer
	var content string
	zap.S().Info(tmplPaht)
	tmpl, err := template.ParseFiles(tmplPaht)
	if err != nil {
		zap.S().Info(err)
		return
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
