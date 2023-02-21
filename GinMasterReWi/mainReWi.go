package main

//
//import (
//	"GinMasterReWi/apps/job/jobs"
//	"GinMasterReWi/pkg/config"
//	"GinMasterReWi/pkg/global"
//	"GinMasterReWi/pkg/initialize"
//	"GinMasterReWi/pkg/middleware"
//	"context"
//	"fmt"
//	"github.com/Re-Wi/GoKitReWi/databases"
//	"github.com/Re-Wi/GoKitReWi/logger"
//	"github.com/XM-GO/PandaKit/restfulx"
//	"github.com/spf13/cobra"
//	"os"
//	"os/signal"
//	"syscall"
//)
//
//var (
//	configFile string
//)
//
//var rootCmd = &cobra.Command{
//	Use:   "GinMasterReWi is the main component in the GinMasterReWi.",
//	Short: `GinMasterReWi is go gin frame`,
//	PreRun: func(cmd *cobra.Command, args []string) {
//		if configFile != "" {
//			global.Conf = config.InitConfig(configFile)
//			global.Log = logger.InitLog(global.Conf.Log.File.GetFilename(), global.Conf.Log.Level)
//			//fmt.Println("okkkk " + global.Conf.Upload.PlanPath)
//			dbGorm := databases.DbGorm{Type: global.Conf.Server.DbType}
//			if global.Conf.Server.DbType == "mysql" {
//				dbGorm.Dsn = global.Conf.Mysql.Dsn()
//				dbGorm.MaxIdleConns = global.Conf.Mysql.MaxIdleConns
//				dbGorm.MaxOpenConns = global.Conf.Mysql.MaxOpenConns
//			} else if global.Conf.Server.DbType == "postgresql" {
//				dbGorm.Dsn = global.Conf.Postgresql.PgDsn()
//				dbGorm.MaxIdleConns = global.Conf.Postgresql.MaxIdleConns
//				dbGorm.MaxOpenConns = global.Conf.Postgresql.MaxOpenConns
//			} else {
//				dbGorm.Dsn = global.Conf.Sqlite3.SqliteDsn()
//				dbGorm.MaxIdleConns = global.Conf.Sqlite3.MaxIdleConns
//				dbGorm.MaxOpenConns = global.Conf.Sqlite3.MaxOpenConns
//			}
//			global.Db = dbGorm.GormInit()
//			initialize.InitTable()
//		} else {
//			global.Log.Panic("请配置config")
//		}
//	},
//	Run: func(cmd *cobra.Command, args []string) {
//		restfulx.UseAfterHandlerInterceptor(middleware.OperationHandler)
//		//// 前置 函数
//		//restfulx.UseBeforeHandlerInterceptor(middleware.PermissionHandler)
//		//// 后置 函数
//		restfulx.UseAfterHandlerInterceptor(middleware.LogHandler)
//		go func() {
//			// 启动系统调度任务
//			global.Log.Info("启动系统调度任务")
//			jobs.InitJob()
//			jobs.Setup()
//		}()
//
//		app := initialize.InitRouter()
//		global.Log.Info("路由初始化完成")
//		app.Start(context.TODO())
//		stop := make(chan os.Signal, 1)
//		signal.Notify(stop, syscall.SIGTERM, os.Interrupt)
//		<-stop
//
//		if err := app.Stop(context.TODO()); err != nil {
//			fmt.Printf("fatal app stop: %s", err)
//			os.Exit(-3)
//		}
//	},
//}
//
//func init() {
//	rootCmd.Flags().StringVar(&configFile, "config", getEnvStr("GinMasterReWi_CONFIG", "./config/config.yml"), "GinMasterReWi config file path.")
//}
//
//func getEnvStr(env string, defaultValue string) string {
//	v := os.Getenv(env)
//	if v == "" {
//		return defaultValue
//	}
//	return v
//}
//
//func main() {
//	if err := rootCmd.Execute(); err != nil {
//		rootCmd.PrintErrf("GinMasterReWi root cmd execute: %s", err)
//		os.Exit(1)
//	}
//}
