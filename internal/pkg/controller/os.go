package controller

import (
	"net"
	"runtime"
	"strings"

	"github.com/shirou/gopsutil/mem"
)

// OSInformation define all needed information like the operating system, ip address and memory
type OSInformation struct {
	OperatingSystem string `json:"operatingSystem,omitempty"`
	IP              string `json:"ip,omitempty"`
	Memory          Memory `json:"memory,omitempty"`
}

// Memory defines the memory information
type Memory struct {
	Total       uint64  `json:"total,omitempty"`
	Free        uint64  `json:"free,omitempty"`
	UsedPercent float64 `json:"usedPercent,omitempty"`
}

// GetOSInformation return all needed information like the operating system, ip address and memory
func GetOSInformation() OSInformation {
	return OSInformation{
		OperatingSystem: operatingSystem(),
		IP:              externalIP(),
		Memory:          memory(),
	}
}

func operatingSystem() string {
	switch strings.ToLower(runtime.GOOS) {
	case "windows":
		return "Windows"
	case "linux":
		return "Linux"
	case "darwin":
		return "MacOS"
	default:
		return "Unknown"
	}
}

func externalIP() string {
	ifaces, _ := net.Interfaces()
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return ""
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String()
		}
	}
	return ""
}

func memory() Memory {
	v, _ := mem.VirtualMemory()
	return Memory{
		Total:       v.Total,
		Free:        v.Free,
		UsedPercent: v.UsedPercent,
	}
}