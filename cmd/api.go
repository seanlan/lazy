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
		tmplDir, _ := cmd.Flags().GetString("template")
		packageName, _ := cmd.Flags().GetString("package")
		modelName, _ := cmd.Flags().GetString("model")
		apiName, _ := cmd.Flags().GetString("api")
		apiOutPath, _ := cmd.Flags().GetString("api-out")
		modelOutPath, _ := cmd.Flags().GetString("model-out")
		serviceOutPath, _ := cmd.Flags().GetString("service-out")
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
		var (
			_modelOutPath   string
			_serviceOutPath string
			_apiOutPath     string
		)
		_modelOutPath = modelOutPath
		_serviceOutPath = serviceOutPath
		_apiOutPath = apiOutPath
		apiMap := map[string]interface{}{
			"PackageName":   packageName,
			"ApiName":       apiName,
			"ReqStructName": apiName + "Req",
			"ServiceName":   apiName,
			"ModelPath":     _modelOutPath,
			"ApiPath":       _apiOutPath,
			"ServicePath":   _serviceOutPath,
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
			"ModelPath":      _modelOutPath,
			"ApiPath":        _apiOutPath,
			"ServicePath":    _serviceOutPath,
		}
		serviceOutFile := filepath.Join(serviceOutPath, modelName+".go")
		generator.RenderWithMap(serviceOutFile, tmplDir, "gin_service.tmpl", serviceMap, false)

	},
}

func init() {
	var err error
	apiCmd.Flags().String("model", "model", "request/response model package name")
	err = apiCmd.MarkFlagRequired("model")
	if err != nil {
		zap.S().Error(err)
		return
	}
	apiCmd.Flags().String("api", "api", "api name")
	err = apiCmd.MarkFlagRequired("api")
	if err != nil {
		zap.S().Error(err)
		return
	}
	apiCmd.Flags().String("api-out", "api-out", "api output path")
	err = apiCmd.MarkFlagRequired("api-out")
	if err != nil {
		zap.S().Error(err)
		return
	}
	apiCmd.Flags().String("model-out", "model-out", "model output path")
	err = apiCmd.MarkFlagRequired("model-out")
	if err != nil {
		zap.S().Error(err)
		return
	}
	apiCmd.Flags().String("service-out", "service-out", "service output path")
	err = apiCmd.MarkFlagRequired("service-out")
	if err != nil {
		zap.S().Error(err)
		return
	}
	apiCmd.Flags().String("package", "", "project package name")
	err = apiCmd.MarkFlagRequired("package")
	if err != nil {
		zap.S().Error(err)
		return
	}
	apiCmd.Flags().String("template", "", "template dir")
	rootCmd.AddCommand(apiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// apiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// apiCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
