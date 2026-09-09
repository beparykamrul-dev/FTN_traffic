package dataplane

import "context"

type BGPSummarySource interface { SummaryOutput(context.Context) (string,error) }
func ParseBGPSummary(kind RouterKind,out string)([]BGPSession,error){switch kind{case RouterFRR:return ParseFRRBGPSummary(out);case RouterBIRD:return ParseBIRDBGPSummary(out);case RouterGoBGP:return ParseGoBGPNeighborSummary(out);default:return nil,ErrBGPOutputUnrecognized}}
