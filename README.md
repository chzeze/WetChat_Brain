# wechat_brain
小程序王者头脑自动刷分工具

## 改进：
改进：https://github.com/sundy-li/wechat_brain


## 使用步骤：
Android 设备

	1. Android 手机一台，电脑上安装 ADB，连接上电脑后开启 USB 调试模式，开发者选项中有模拟触摸选项的请一并开启
	2. 安装证书。手机浏览器访问 `abc.com`安装证书,ios记得要信任证书 (或者将 `certs/goproxy.crt`传到手机, 点击安装证书)。
	3. 设置手机代理。手机连接wifi后进行代理设置，代理IP为个人pc的ip和端口为8998。
	4. 运行主程序。运行方法（二选一）：
		安装go(>=1.8)环境后, clone本repo源码到对应`$GOPATH/src/github.com/sundy-li/`下, 进入源码目录后,执行 `go run cmd/main.go`。
	5. 打开微信并启动王者头脑小程序。
	6. 正确的答案将在小程序的选项中以【标准答案】或【数字】字样标识，abd模拟提交






