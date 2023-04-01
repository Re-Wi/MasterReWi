# 关闭任务
supervisorctl stop MasterReWi
# 启动任务
supervisorctl start MasterReWi
# 更新配置文件并重启相关的程序
supervisorctl update
# 查看bluebell的运行状态
supervisorctl status MasterReWi

# 重新加载Nginx配置文件
bash /home/admin/system/restart.sh

#cd ..
#go run main.go &