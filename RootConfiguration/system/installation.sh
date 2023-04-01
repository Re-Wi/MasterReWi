# 如有需要，先新建用户
ls /home
useradd admin -m -d /home/admin -s /bin/bash
chmod -R 777 /home/admin
passwd admin

# 更新安装环境
apt-get upgrade -y
apt-get update -y

# 安装编辑器nano
apt-get install nano -y

# 安装ssh服务
apt-get install openssh-server -y
# 启动sshserver
/etc/init.d/ssh start
# 重启ssh服务
/etc/init.d/ssh restart

# 安装gcc
apt install -y gcc
# 查看新版本
gcc -v

##########   安装中文语言包  ################
# 查看当前语言
locale
# 查看当前已安装的语言
locale -a
# ubuntu系统更新软件包列表
apt-get update  
apt-get install  -y language-pack-zh-hans
apt-get install -y language-pack-zh-hant
# 字符集
locale-gen zh_CN.UTF-8
# 再次查看
locale -a
echo "export LC_ALL=zh_CN.UTF-8" >> /etc/profile
source /etc/profile
# 如果这里添加失败，提示没有这种语言包，退出容器，再重新进入，就可以添加了
# 完成
locale

# 安装systemctl
apt install -y systemctl

# 启动ssh服务并配置开机自启
systemctl enable ssh #开机启动 
systemctl restart ssh #重启ssh 
systemctl status ssh #查看ssh状态

# 安装Go环境
apt-get install -y golang
echo "export GOPROXY=https://goproxy.io,direct" >> /etc/profile

# 部署程序
apt-get install -y supervisor

# 安装nginx
apt-get install -y nginx
# 显示版本信息并退出
nginx -v 
