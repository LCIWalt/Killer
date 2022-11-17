
# 让你的舍友游戏体验极为良好
**golang使用略**

### 编译
用管理员powershell打开文件夹 编译
go build -ldflags "-s -w -H=windowsgui"

### 别忘修改xml
conf.txt中填写你想要kill的游戏名(进程名)xxx(不知道的可以开一个然后ctrl+alt+delete设备管理器查看xxx.exe)
*这里使用了csgo*

### command settings

cmd用管理员权限开启
schtasks /CREATE /TN "输入你想要命名计划的名字 与XML中的尽量一致" /XML + XML配置文件的绝对路径