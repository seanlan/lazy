/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bufio"
	"github.com/gobuffalo/packr/v2"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
	"path/filepath"
)

func saveFile(t, savePath string) {
	var boxName string
	switch t {
	case "sh":
		boxName = "lazy-scripts"
		break
	default:
		boxName = "lazy-templates"
	}
	err := os.MkdirAll(savePath, os.ModePerm)
	if err != nil {
		zap.S().Errorf("Failed to create directory %s", savePath)
		return
	}
	box := packr.New(boxName, "../templates")
	files := box.List()
	for _, file := range files {
		outPath := filepath.Join(savePath, file)
		content, err := box.FindString(file)
		if err != nil {
			zap.S().Info(err)
			return
		}
		var fs *os.File
		_, err = os.Stat(outPath)
		if err != nil { // file not exist
			fs, err = os.Create(outPath) //create file
		} else { // file exist
			fs, err = os.OpenFile(outPath, os.O_WRONLY|os.O_TRUNC, 0666)
		}
		writer := bufio.NewWriter(fs)
		_, _ = writer.WriteString(content)
		_ = writer.Flush()
		if fs != nil {
			_ = fs.Close()
		}
		zap.S().Info("save template success ", outPath)
	}
}

func saveFunc(cmd *cobra.Command, args []string) {
	t, _ := cmd.Flags().GetString("t")
	savePath, _ := cmd.Flags().GetString("path")
	saveFile(t, savePath)

}

// saveCmd represents the save command
var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "save packrd templates to path",
	Run:   saveFunc,
}

func init() {
	saveCmd.Flags().String("t", "tmpl", "save template type: tmpl or sh")
	saveCmd.Flags().String("path", "./tmp", "template save path")
	rootCmd.AddCommand(saveCmd)
}
