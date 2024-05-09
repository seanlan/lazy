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
	"github.com/spf13/cobra"
)

func testFunc(cmd *cobra.Command, args []string) {
	//connStr := "root:q145145145@tcp(127.0.0.1:3306)/lucky?parseTime=true&loc=Local&charset=utf8mb4&collation=utf8mb4_unicode_ci"
	//db, _ := gorm.Open(mysql.Open(connStr),
	//	&gorm.Config{
	//		NamingStrategy: _schema.NamingStrategy{
	//			TablePrefix: "", SingularTable: true,
	//		},
	//	})
	//dao.DB = db.Debug()
	//ctx := context.Background()
	//q := sqlmodel.AdminUserColumns
	//var aus []sqlmodel.AdminUser
	//var au sqlmodel.AdminUser
	//zap.S().Info(dao.CountAdminUser(ctx, nil))
	//zap.S().Info(dao.SumAdminUser(ctx, q.ID, dao.And(q.ID.Eq(1), q.Username.Eq("seanlan"))))
	//zap.S().Info(dao.FetchAllAdminUser(
	//	ctx, &aus,
	//	dao.And(q.ID.Eq(1), q.Username.Eq("seanlan")), 2, 10,
	//	q.ID.Desc(), q.CreateAt.Desc()))
	//zap.S().Info(
	//	dao.FetchAdminUser(ctx, &au,
	//		dao.And(q.ID.Eq(1), q.Username.Eq("seanlan")),
	//		q.ID.Desc(), q.CreateAt.Desc()))
	//zap.S().Info(dao.AddAdminUser(ctx, &sqlmodel.AdminUser{
	//	Password: "seanlan",
	//	Username: "seanlan",
	//}))
	//zap.S().Info(dao.AddsAdminUser(ctx, &[]sqlmodel.AdminUser{
	//	{
	//		Password: "seanlan",
	//		Username: "seanlan",
	//	},
	//	{
	//		Password: "seanlan",
	//		Username: "seanlan",
	//	},
	//}))
	//au.Password = "111111"
	//zap.S().Info(dao.UpdateAdminUser(ctx, &au))
	//zap.S().Info(dao.UpsertAdminUser(ctx, &au, dao.M{q.Password.FieldName: "99999"}))
	//zap.S().Info(dao.DeleteAdminUser(ctx, q.Account.Eq("")))
}

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "test command",
	Run:   testFunc,
}

func init() {
	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
