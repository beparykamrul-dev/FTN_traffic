package dataplane
import("context";"fmt")
type BGPAdapter interface{RouterAdapter;SessionSummary(context.Context)([]BGPSession,error)}
type BGPSession struct{ID string `json:"id"`;RemoteASN uint32 `json:"remote_asn"`;Established bool `json:"established"`;IPv4 bool `json:"ipv4"`;IPv6 bool `json:"ipv6"`;RPKIValid bool `json:"rpki_valid"`;PrefixesIn uint64 `json:"prefixes_in"`;PrefixesOut uint64 `json:"prefixes_out"`}
func ParseBGPSummary(kind,out string)([]BGPSession,error){switch kind{case string(RouterFRR),"frr":return ParseFRRBGPSummary(out);case string(RouterBIRD),"bird":return ParseBIRDBGPSummary(out);case string(RouterGoBGP),"gobgp":return ParseGoBGPNeighborSummary(out);default:return nil,fmt.Errorf("unsupported BGP backend: %s",kind)}}
