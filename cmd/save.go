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

// saveCmd represents the save command
var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "save packrd templates to path",
	Run: func(cmd *cobra.Command, args []string) {
		savePath, _ := cmd.Flags().GetString("path")
		box := packr.New("lazy", "../templates")
		_ = os.MkdirAll(savePath, os.ModePerm)
		for _, t := range []string{
			"dao_gorm.tmpl",
			"dao_gorm_base.tmpl",
			"dao_gorm_model.tmpl",
			"gin_model.tmpl",
			"gin_api.tmpl",
			"gin_service.tmpl",
		} {
			outPath := filepath.Join(savePath, t)
			content, err := box.FindString(t)
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
	},
}

func init() {
	saveCmd.Flags().String("path", "./templates", "template save path")
	rootCmd.AddCommand(saveCmd)
}
