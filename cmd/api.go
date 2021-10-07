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
	"github.com/seanlan/lazy/generator"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
	"path/filepath"
)

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "generate api service model",
	Run: func(cmd *cobra.Command, args []string) {
		tmplDir := ""
		packageName := "lazy"
		modelName := "test"
		apiName := "TestApi2"
		apiOutPath := "./app/api/v1"
		modelOutPath := "./app/model"
		serviceOutPath := "./app/service"
		err := os.MkdirAll(apiOutPath, os.ModePerm)
		if err != nil {
			zap.S().Info("mkdir error: ", err, apiOutPath)
			return
		}
		err = os.MkdirAll(modelOutPath, os.ModePerm)
		if err != nil {
			zap.S().Info("mkdir error: ", err, modelOutPath)
			return
		}
		err = os.MkdirAll(serviceOutPath, os.ModePerm)
		if err != nil {
			zap.S().Info("mkdir error: ", err, serviceOutPath)
			return
		}

		apiMap := map[string]interface{}{
			"PackageName":   packageName,
			"ApiName":       apiName,
			"ReqStructName": apiName + "Req",
			"ServiceName":   apiName,
		}
		apiOutFile := filepath.Join(apiOutPath, modelName+".go")
		generator.RenderWithMap(apiOutFile, tmplDir, "gin_api.tmpl", apiMap, false)
		modelMap := map[string]interface{}{
			"ReqStructName":  apiName + "Req",
			"RespStructName": apiName + "Resp",
		}
		modelOutFile := filepath.Join(modelOutPath, modelName+".go")
		generator.RenderWithMap(modelOutFile, tmplDir, "gin_model.tmpl", modelMap, false)
		serviceMap := map[string]interface{}{
			"PackageName":    packageName,
			"ServiceName":    apiName,
			"ReqStructName":  apiName + "Req",
			"RespStructName": apiName + "Resp",
		}
		serviceOutFile := filepath.Join(serviceOutPath, modelName+".go")
		generator.RenderWithMap(serviceOutFile, tmplDir, "gin_service.tmpl", serviceMap, false)

	},
}

func init() {
	rootCmd.AddCommand(apiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// apiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// apiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
