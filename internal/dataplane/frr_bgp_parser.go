package dataplane

import("errors";"strconv";"strings")

var ErrBGPOutputUnrecognized=errors.New("unrecognized BGP session output")
func parseUint(s string)uint64{v,_:=strconv.ParseUint(strings.TrimSpace(s),10,64);return v}
func ParseFRRBGPSummary(out string) ([]BGPSession,error){var sessions []BGPSession;for _,line:=range strings.Split(out,"\n"){f:=strings.Fields(line);if len(f)<5||strings.EqualFold(f[0],"Neighbor")||strings.HasPrefix(f[0],"BGP"){continue};if strings.Contains(f[0],".")==false&&strings.Contains(f[0],":")==false{continue};s:=BGPSession{ID:f[0]};if n,e:=strconv.ParseUint(f[2],10,32);e==nil{s.RemoteASN=uint32(n)};last:=f[len(f)-1];if n,e:=strconv.ParseUint(last,10,64);e==nil{s.PrefixesIn=n;s.Established=true}else if len(f)>=2{s.Established=!strings.Contains(strings.ToLower(last),"idle")&&!strings.Contains(strings.ToLower(last),"active")};sessions=append(sessions,s)};if len(sessions)==0{return nil,ErrBGPOutputUnrecognized};return sessions,nil}
