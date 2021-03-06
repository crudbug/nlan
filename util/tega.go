package util

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"

	"github.com/araobp/nlan/model/nlan"
	"github.com/araobp/tega/driver"

	netstat "github.com/drael/GOnetstat"
)

var ope *driver.Operation
var hostname string

const (
	IP_PATH          = "ip"
	HOSTS_PATH       = "hosts"
	RAW_PATH         = "raw"
	CONFIG_PATH      = "config"
	OPERATIONAL_PATH = "operational"
	STATS_PATH       = "stats"
)

type Self struct {
}

func (r *Self) OnInit() {
	ope.RegisterRpc(fmt.Sprintf("%s.%s", RAW_PATH, hostname), raw)
	ope.RegisterRpc(fmt.Sprintf("%s-%s.ip.hook.route", OPERATIONAL_PATH, hostname), ipRoute)
	ope.RegisterRpc(fmt.Sprintf("%s-%s.ip.hook.addr", OPERATIONAL_PATH, hostname), ipAddr)
	ope.RegisterRpc(fmt.Sprintf("%s-%s.hook.netstat", STATS_PATH, hostname), netstat_)

	registerHost()
}

func (r *Self) OnNotify(v *[]driver.Notification) {
}

func (r *Self) OnMessage(channel string, tegaId string, msg *driver.Message) {
}

func init() {
	var err error
	hostname = os.Getenv("HOSTNAME")
	if hostname == "" {
		hostname = "localhost"
	}
	tega := os.Getenv("TEGA_ADDRESS")
	self := &Self{}
	ope, err = driver.NewOperation(hostname, tega, 0, self, driver.LOCAL)
	if err != nil {
		log.Fatal(err)
	}
}

// Returns os.Getenv("HOSTNAME")
func GetHostname() string {
	return hostname
}

// Registers a host IP address with tega
func registerHost() {
	var err error
	if err != nil {
		log.Fatal(err)
	}
	path := fmt.Sprintf("%s.%s", HOSTS_PATH, hostname)
	interfaces, _ := net.Interfaces()
	var addrs []net.Addr
	for _, inter := range interfaces {
		if inter.Name == "eth0" {
			addrs, _ = inter.Addrs()
		}
	}
	value := addrs[0].String()
	err = ope.PutE(path, value)
	if err != nil {
		log.Fatal(err)
	}
}

// Gets a list of all host names and their addresses from tega
func listHosts(path string) map[string]interface{} {
	var nodes map[string]interface{}
	err := ope.Get(path, &nodes)
	if err != nil {
		log.Fatal(err)
	}
	hosts := make(map[string]interface{})
	for host, ipmask := range nodes {
		log.Print(host)
		log.Print(ipmask)
		hosts[host] = strings.Split(ipmask.(string), "/")[0]
	}
	return hosts
}

// Lists up all hosts on HOSTS_PATH
func ListHosts() map[string]interface{} {
	return listHosts(HOSTS_PATH)
}

// Lists up all hosts on IP_PATH
func ListIps() map[string]interface{} {
	return listHosts(IP_PATH)
}

// Sets NLAN config to tega
func SetModel(hostname string, model *nlan.Model) {
	path := fmt.Sprintf("%s-%s", CONFIG_PATH, hostname)
	err := ope.Put(path, model)
	if err != nil {
		log.Fatal(err)
	}
}

// Gets NLAN state from tega
func GetModel(hostname string, model *nlan.Model) {
	path := fmt.Sprintf("%s-%s", CONFIG_PATH, hostname)
	err := ope.Get(path, model)
	if err != nil {
		log.Fatal(err)
	}
}

// Resets NLAN state on tega
func ResetState() {
	// TODO: implementation
	/*
		err := ope.Delete("nlan.state")
		if err != nil {
			log.Print(err)
		}
	*/
}

// Gets a secondary IP address from tega
func GetSecondaryIp(hostname string) string {
	path := fmt.Sprintf("%s.%s", IP_PATH, hostname)
	var secondary string
	err := ope.Get(path, &secondary)
	if err != nil {
		log.Fatal(err)
	}
	return secondary
}

// Executes a raw command (i.e., shell command)
func raw(argsKwargs driver.ArgsKwargs) (driver.Result, error) {
	args := strings.Split(argsKwargs.Args[0].(string), " ")
	cmd := args[0]
	var cmdArgs []string
	if len(args) > 1 {
		cmdArgs = args[1:]
	}
	result, _ := OutputCmd(cmd, cmdArgs...) // Executes a raw command
	return driver.Result{Res: result}, nil
}

func ipRoute(argsKwargs driver.ArgsKwargs) (driver.Result, error) {
	value := driver.Result{Res: RouteMap()}
	path := fmt.Sprintf("%s-%s.ip.route", OPERATIONAL_PATH, hostname)
	ope.PutE(path, value.Res)
	return value, nil
}

func ipAddr(argsKwargs driver.ArgsKwargs) (driver.Result, error) {
	devMap, addrMap := AddrMap()
	devPath := fmt.Sprintf("%s-%s.ip.dev", OPERATIONAL_PATH, hostname)
	addrPath := fmt.Sprintf("%s-%s.ip.addr", OPERATIONAL_PATH, hostname)
	ope.PutE(devPath, devMap)
	ope.PutE(addrPath, addrMap)
	res := []interface{}{devMap, addrMap}
	value := driver.Result{Res: res}
	return value, nil
}

func netstat_(argsKwargs driver.ArgsKwargs) (driver.Result, error) {
	tcpstat := Netstat(TCP)
	udpstat := Netstat(UDP)
	data := make(map[string]*[]netstat.Process)
	data["tcp"] = tcpstat
	data["udp"] = udpstat
	netstatPath := fmt.Sprintf("%s-%s.netstat", STATS_PATH, hostname)
	ope.PutE(netstatPath, data)
	value := driver.Result{Res: data}
	return value, nil
}
