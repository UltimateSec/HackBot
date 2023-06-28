package core

import (
	"fmt"
	"github.com/projectdiscovery/goflags"
	"github.com/projectdiscovery/gologger"
	"github.com/projectdiscovery/gologger/levels"
	"github.com/projectdiscovery/naabu/v2/pkg/result"
	"github.com/projectdiscovery/naabu/v2/pkg/runner"
)

type Naabu struct {
	Host  string `json:"host"`
	Ports string `json:"ports"`
}

var DefaultPorts = "80,81,82,83,84,85,86,87,88,89,90,91,92,98,99,443,800,801,808,880,888,889,1000,1010,1080,1081,1082,1099,1118,1888,2008,2020,2100,2375,2379,3000,3008,3128,3505,5555,5858,6080,6648,6868,7000,7001,7002,7003,7004,7005,7007,7008,7070,7071,7074,7078,7080,7088,7200,7680,7687,7688,7777,7890,8000,8001,8002,8003,8004,8006,8008,8009,8010,8011,8012,8016,8018,8020,8028,8030,8038,8042,8044,8046,8048,8053,8060,8069,8070,8080,8081,8082,8083,8084,8085,8086,8087,8088,8089,8090,8091,8092,8093,8094,8095,8096,8097,8098,8099,8100,8101,8108,8118,8161,8172,8180,8181,8200,8222,8244,8258,8280,8288,8300,8360,8443,8448,8484,8800,8834,8838,8848,8858,8868,8879,8880,8881,8888,8899,8983,8989,9000,9001,9002,9008,9010,9043,9060,9080,9081,9082,9083,9084,9085,9086,9087,9088,9089,9090,9091,9092,9093,9094,9095,9096,9097,9098,9099,9100,9200,9229,9443,9448,9800,9981,9986,9988,9998,9999,10000,10001,10002,10004,10008,10010,10250,12018,12443,14000,16080,18000,18001,18002,18004,18008,18080,18082,18088,18090,18098,19001,20000,20720,21000,21501,21502,28018,20880,45554,49155,49156,50050,61616"

func (n *Naabu) Run() (content string, err error) {
	gologger.DefaultLogger.SetMaxLevel(levels.LevelSilent)
	if n.Ports == "" {
		n.Ports = DefaultPorts
	}
	options := runner.Options{
		Host:      goflags.StringSlice{n.Host},
		ScanType:  runner.ConnectScan,
		Timeout:   runner.DefaultPortTimeoutConnectScan,
		Threads:   100,
		ResumeCfg: &runner.ResumeCfg{},
		OnResult: func(hr *result.HostResult) {
			var ports []int
			for _, port := range hr.Ports {
				ports = append(ports, port.Port)
			}
			content = fmt.Sprintf("%s open:%v", hr.IP, ports)
		},
		Ports: n.Ports,
	}

	r, err := runner.NewRunner(&options)
	if err != nil {
		return
	}
	defer r.Close()

	r.RunEnumeration()
	return
}
