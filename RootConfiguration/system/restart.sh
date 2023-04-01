# 更新配置文件并重启相关的程序
supervisorctl update

# 查看所有任务状态
sleep 1
supervisorctl status

# 重新加载Nginx配置文件
nginx -s reload
