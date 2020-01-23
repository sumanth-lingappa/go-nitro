package ns

type Nsip struct {
	Advertiseondefaultpartition string      `json:"advertiseondefaultpartition,omitempty"`
	Arp                         string      `json:"arp,omitempty"`
	Arpowner                    int         `json:"arpowner,omitempty"`
	Arpresponse                 string      `json:"arpresponse,omitempty"`
	Bgp                         string      `json:"bgp,omitempty"`
	Decrementttl                string      `json:"decrementttl,omitempty"`
	Dynamicrouting              string      `json:"dynamicrouting,omitempty"`
	Flags                       int         `json:"flags,omitempty"`
	Freeports                   int         `json:"freeports,omitempty"`
	Ftp                         string      `json:"ftp,omitempty"`
	Gui                         string      `json:"gui,omitempty"`
	Hostroute                   string      `json:"hostroute,omitempty"`
	Hostrtgw                    string      `json:"hostrtgw,omitempty"`
	Hostrtgwact                 string      `json:"hostrtgwact,omitempty"`
	Icmp                        string      `json:"icmp,omitempty"`
	Icmpresponse                string      `json:"icmpresponse,omitempty"`
	Ipaddress                   string      `json:"ipaddress,omitempty"`
	Iptype                      interface{} `json:"iptype,omitempty"`
	Metric                      int         `json:"metric,omitempty"`
	Mgmtaccess                  string      `json:"mgmtaccess,omitempty"`
	Netmask                     string      `json:"netmask,omitempty"`
	Networkroute                string      `json:"networkroute,omitempty"`
	Operationalarpowner         int         `json:"operationalarpowner,omitempty"`
	Ospf                        string      `json:"ospf,omitempty"`
	Ospfarea                    int         `json:"ospfarea,omitempty"`
	Ospfareaval                 int         `json:"ospfareaval,omitempty"`
	Ospflsatype                 string      `json:"ospflsatype,omitempty"`
	Ownerdownresponse           string      `json:"ownerdownresponse,omitempty"`
	Ownernode                   int         `json:"ownernode,omitempty"`
	Restrictaccess              string      `json:"restrictaccess,omitempty"`
	Rip                         string      `json:"rip,omitempty"`
	Riserhimsgcode              int         `json:"riserhimsgcode,omitempty"`
	Snmp                        string      `json:"snmp,omitempty"`
	Ssh                         string      `json:"ssh,omitempty"`
	State                       string      `json:"state,omitempty"`
	Tag                         int         `json:"tag,omitempty"`
	Td                          int         `json:"td,omitempty"`
	Telnet                      string      `json:"telnet,omitempty"`
	Type                        string      `json:"type,omitempty"`
	Viprtadv2bsd                bool        `json:"viprtadv2bsd,omitempty"`
	Vipvsercount                int         `json:"vipvsercount,omitempty"`
	Vipvserdowncount            int         `json:"vipvserdowncount,omitempty"`
	Vipvsrvrrhiactivecount      int         `json:"vipvsrvrrhiactivecount,omitempty"`
	Vipvsrvrrhiactiveupcount    int         `json:"vipvsrvrrhiactiveupcount,omitempty"`
	Vrid                        int         `json:"vrid,omitempty"`
	Vserver                     string      `json:"vserver,omitempty"`
	Vserverrhilevel             string      `json:"vserverrhilevel,omitempty"`
	Vserverrhimode              string      `json:"vserverrhimode,omitempty"`
}
