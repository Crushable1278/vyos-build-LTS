package all

import (
	//Blank imports for plugins to register themselves
	_ "github.com/influxdata/telegraf/plugins/inputs/azure_storage_queue"
	_ "github.com/influxdata/telegraf/plugins/inputs/bond"
	_ "github.com/influxdata/telegraf/plugins/inputs/cgroup"
	_ "github.com/influxdata/telegraf/plugins/inputs/chrony"
	_ "github.com/influxdata/telegraf/plugins/inputs/conntrack"
	_ "github.com/influxdata/telegraf/plugins/inputs/cpu"
	_ "github.com/influxdata/telegraf/plugins/inputs/disk"
	_ "github.com/influxdata/telegraf/plugins/inputs/diskio"
	_ "github.com/influxdata/telegraf/plugins/inputs/disque"
	_ "github.com/influxdata/telegraf/plugins/inputs/dmcache"
	_ "github.com/influxdata/telegraf/plugins/inputs/dns_query"
	_ "github.com/influxdata/telegraf/plugins/inputs/docker"
	_ "github.com/influxdata/telegraf/plugins/inputs/docker_log"
	_ "github.com/influxdata/telegraf/plugins/inputs/ethtool"
	_ "github.com/influxdata/telegraf/plugins/inputs/exec"
	_ "github.com/influxdata/telegraf/plugins/inputs/execd"
	_ "github.com/influxdata/telegraf/plugins/inputs/file"
	_ "github.com/influxdata/telegraf/plugins/inputs/filecount"
	_ "github.com/influxdata/telegraf/plugins/inputs/filestat"
	_ "github.com/influxdata/telegraf/plugins/inputs/fireboard"
	_ "github.com/influxdata/telegraf/plugins/inputs/hddtemp"
	_ "github.com/influxdata/telegraf/plugins/inputs/hugepages"
	_ "github.com/influxdata/telegraf/plugins/inputs/influxdb"
	_ "github.com/influxdata/telegraf/plugins/inputs/influxdb_listener"
	_ "github.com/influxdata/telegraf/plugins/inputs/influxdb_v2_listener"
	_ "github.com/influxdata/telegraf/plugins/inputs/intel_pmu"
	_ "github.com/influxdata/telegraf/plugins/inputs/intel_powerstat"
	_ "github.com/influxdata/telegraf/plugins/inputs/intel_rdt"
	_ "github.com/influxdata/telegraf/plugins/inputs/internal"
	_ "github.com/influxdata/telegraf/plugins/inputs/internet_speed"
	_ "github.com/influxdata/telegraf/plugins/inputs/interrupts"
	_ "github.com/influxdata/telegraf/plugins/inputs/ipmi_sensor"
	_ "github.com/influxdata/telegraf/plugins/inputs/ipset"
	_ "github.com/influxdata/telegraf/plugins/inputs/iptables"
	_ "github.com/influxdata/telegraf/plugins/inputs/ipvs"
	_ "github.com/influxdata/telegraf/plugins/inputs/kernel"
	_ "github.com/influxdata/telegraf/plugins/inputs/kernel_vmstat"
	_ "github.com/influxdata/telegraf/plugins/inputs/mdstat"
	_ "github.com/influxdata/telegraf/plugins/inputs/mem"
	_ "github.com/influxdata/telegraf/plugins/inputs/net"
	_ "github.com/influxdata/telegraf/plugins/inputs/netstat"
	_ "github.com/influxdata/telegraf/plugins/inputs/nstat"
	_ "github.com/influxdata/telegraf/plugins/inputs/ping"
	_ "github.com/influxdata/telegraf/plugins/inputs/powerdns_recursor"
	_ "github.com/influxdata/telegraf/plugins/inputs/processes"
	_ "github.com/influxdata/telegraf/plugins/inputs/procstat"
	_ "github.com/influxdata/telegraf/plugins/inputs/sensors"
	_ "github.com/influxdata/telegraf/plugins/inputs/sflow"
	_ "github.com/influxdata/telegraf/plugins/inputs/slab"
	_ "github.com/influxdata/telegraf/plugins/inputs/smart"
	_ "github.com/influxdata/telegraf/plugins/inputs/snmp"
	_ "github.com/influxdata/telegraf/plugins/inputs/snmp_legacy"
	_ "github.com/influxdata/telegraf/plugins/inputs/snmp_trap"
	_ "github.com/influxdata/telegraf/plugins/inputs/socket_listener"
	_ "github.com/influxdata/telegraf/plugins/inputs/socketstat"
	_ "github.com/influxdata/telegraf/plugins/inputs/syslog"
	_ "github.com/influxdata/telegraf/plugins/inputs/sysstat"
	_ "github.com/influxdata/telegraf/plugins/inputs/system"
	_ "github.com/influxdata/telegraf/plugins/inputs/systemd_units"
	_ "github.com/influxdata/telegraf/plugins/inputs/tail"
	_ "github.com/influxdata/telegraf/plugins/inputs/tcp_listener"
	_ "github.com/influxdata/telegraf/plugins/inputs/temp"
	_ "github.com/influxdata/telegraf/plugins/inputs/twemproxy"
	_ "github.com/influxdata/telegraf/plugins/inputs/udp_listener"
	_ "github.com/influxdata/telegraf/plugins/inputs/wireguard"
	_ "github.com/influxdata/telegraf/plugins/inputs/wireless"
	_ "github.com/influxdata/telegraf/plugins/inputs/x509_cert"
)
