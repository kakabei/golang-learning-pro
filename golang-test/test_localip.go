package main

import (
	"fmt"
	"net"
	"strings"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/v3/mem"
)

func getLocalIpV4() string {
	inters, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	for _, inter := range inters {
		// 判断网卡是否开启，过滤本地环回接口
		if inter.Flags&net.FlagUp != 0 && !strings.HasPrefix(inter.Name, "lo") {
			// 获取网卡下所有的地址
			addrs, err := inter.Addrs()
			if err != nil {
				continue
			}
			for _, addr := range addrs {
				if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					//判断是否存在IPV4 IP 如果没有过滤
					if ipnet.IP.To4() != nil {
						fmt.Printf(ipnet.IP.String())
						//return ipnet.IP.String()
					}
				}
			}
		}
	}
	return ""
}

func test_Men() {
	v, _ := mem.VirtualMemory()

	// almost every return value is a struct
	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	fmt.Println(v)
}

func test_CPU() {

	// infos, _ := cpu.Times(false)
	// for _, info := range infos {
	// 	data, _ := json.MarshalIndent(info, "", " ")
	// 	fmt.Print(string(data))
	// }

	totalPercent, _ := cpu.Percent(3*time.Second, false)
	fmt.Printf("total percent:%v", totalPercent)

}

func GetMemPercent() float64 {
	memInfo, _ := mem.VirtualMemory()
	return memInfo.UsedPercent
}

func GetDiskPercent() float64 {
	parts, _ := disk.Partitions(true)
	diskInfo, _ := disk.Usage(parts[0].Mountpoint)
	return diskInfo.UsedPercent
}

func main() {
	fmt.Printf("GetMemPercent: %v\n", GetMemPercent())
	fmt.Printf("GetDiskPercent: %v", GetDiskPercent())
}
