#重启ssh服务
/etc/init.d/ssh restart
systemctl restart ssh

# 重启环境
# 注意启用本脚本使用 source 命令，否则不生效
source /etc/profile

# 优雅的关闭nginx,意为当前访问结束后，关闭nginx
mkdir -p /home/admin/logs
nginx -s stop
nginx -tc /home/admin/conf/nginx.conf
nginx -c /home/admin/conf/nginx.conf

# 启动supervisor服务
supervisord -c /etc/supervisor/supervisord.conf
