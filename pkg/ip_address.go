package pkg

import (
	"bluebell/config"
	"github.com/fatih/color"
	"go.uber.org/zap"
	"net"
)

func PrintIPAddress() {
	// 获取所有网络接口
	interfaces, err := net.Interfaces()
	if err != nil {
		zap.L().Error("net.Interfaces err", zap.Error(err))
		return
	}

	// 遍历所有网络接口
	for _, iface := range interfaces {
		// 获取接口的地址
		addrs, err := iface.Addrs()
		if err != nil {
			zap.L().Error("ip address", zap.String("interface", iface.Name), zap.Error(err))
			continue
		}

		// 遍历接口的所有地址
		for _, addr := range addrs {
			// 检查地址是否为 IP 地址
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				// 检查是否为 IPv4 地址
				if ipnet.IP.To4() != nil && iface.Name == "WLAN" {
					color.Green("%v IP Address: %v:%v", iface.Name,
						ipnet.IP.String(), config.Cfg.WebConfig.Port)
				}
			}
		}
	}
	color.Green("Localhost Address: 127.0.0.1:%v", config.Cfg.WebConfig.Port)
}
