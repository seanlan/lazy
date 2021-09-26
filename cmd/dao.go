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
	"os/exec"
)

// daoCmd represents the dao command
var daoCmd = &cobra.Command{
	Use:   "dao",
	Short: "generate gorm dao model",
	Run: func(cmd *cobra.Command, args []string) {
		zap.S().Info()
		dbStr, _ := cmd.Flags().GetString("conn")
		database, _ := cmd.Flags().GetString("database")
		packageName, _ := cmd.Flags().GetString("package")
		tmplPath, _ := cmd.Flags().GetString("template")
		modelPackage, _ := cmd.Flags().GetString("model")
		modelPath, _ := cmd.Flags().GetString("model-path")
		daoPackage, _ := cmd.Flags().GetString("dao")
		daoPath, _ := cmd.Flags().GetString("dao-path")
		//dbStr := "root:q145145145@tcp(127.0.0.1:3306)/mutual?parseTime=true&loc=Local&charset=utf8mb4&collation=utf8mb4_unicode_ci"
		//database := "mutual"
		//packageName := "github.com/seanlan/lazy"
		//tmplPath := "./templates"
		//modelPackage := "sqlmodel"
		//modelPath := "dao/sqlmodel"
		//daoPackage := "dao"
		//daoPath := "./dao"
		g := generator.NewGormGenerator(dbStr, database, packageName, tmplPath,
			modelPackage, modelPath, daoPackage, daoPath)
		if g == nil {
			return
		}
		g.Gen()
		osCmd := exec.Command("gofmt", "-s", "-d", "-w", ".")
		osCmd.Run()
	},
}

func init() {
	daoCmd.Flags().String("conn", "", "mysql grom conn dial")
	daoCmd.MarkFlagRequired("conn")
	daoCmd.Flags().String("database", "", "mysql select database name")
	daoCmd.MarkFlagRequired("database")
	daoCmd.Flags().String("package", "", "project package name")
	daoCmd.MarkFlagRequired("package")
	daoCmd.Flags().String("template", "", "template dir")
	daoCmd.Flags().String("model", "model", "db model package name")
	daoCmd.Flags().String("model-path", "model", "db model out dir")
	daoCmd.Flags().String("dao", "dao", "dao package name")
	daoCmd.Flags().String("dao-path", "dao", "dao out dir")
	rootCmd.AddCommand(daoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// daoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// daoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
