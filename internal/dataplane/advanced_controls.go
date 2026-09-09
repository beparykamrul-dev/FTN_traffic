package dataplane

import (
 "errors"
 "net/netip"
)

var (
 ErrBFDInvalid = errors.New("invalid BFD session")
 ErrRPKIInvalid = errors.New("invalid RPKI route state")
 ErrVRFInvalid = errors.New("invalid VRF")
 ErrVLANInvalid = errors.New("invalid VLAN")
)

type BFDSession struct { ID string `json:"id"`; Local string `json:"local"`; Remote string `json:"remote"`; MinRxMS uint32 `json:"min_rx_ms"`; MinTxMS uint32 `json:"min_tx_ms"`; Multiplier uint8 `json:"multiplier"`; Up bool `json:"up"`; Authorized bool `json:"authorized"` }
func (b BFDSession) Validate() error { if !b.Authorized || b.ID=="" { return ErrUnauthorized }; l,e:=netip.ParseAddr(b.Local); if e!=nil{return ErrBFDInvalid}; rr,e:=netip.ParseAddr(b.Remote); if e!=nil||l.Is4()!=rr.Is4(){return ErrBFDInvalid}; if b.MinRxMS==0||b.MinTxMS==0||b.Multiplier==0{return ErrBFDInvalid}; return nil }

// RPKIState represents the validated origin state used by route policy.
type RPKIState string
const ( RPKIValid RPKIState="valid"; RPKIInvalid RPKIState="invalid"; RPKIUnknown RPKIState="unknown" )
func ValidateRPKI(state RPKIState, required bool) error { if required && state!=RPKIValid{return ErrRPKIInvalid}; return nil }

type VRF struct { Name string `json:"name"`; Table uint32 `json:"table"`; Authorized bool `json:"authorized"` }
func (v VRF) Validate() error { if !v.Authorized||v.Name==""||v.Table==0{return ErrVRFInvalid}; return nil }

type VLAN struct { ID uint16 `json:"id"`; Parent string `json:"parent"`; Authorized bool `json:"authorized"` }
func (v VLAN) Validate() error { if !v.Authorized||v.ID<1||v.ID>4094||v.Parent==""{return ErrVLANInvalid}; return nil }

type QoSPolicy struct { ID string `json:"id"`; RateMbps float64 `json:"rate_mbps"`; BurstKB uint32 `json:"burst_kb"`; DSCP []uint8 `json:"dscp"`; Authorized bool `json:"authorized"` }
func (q QoSPolicy) Validate() error { if !q.Authorized||q.ID==""||q.RateMbps<=0||q.BurstKB==0{return errors.New("invalid QoS policy")}; for _,d:=range q.DSCP{if d>63{return errors.New("invalid DSCP")}}; return nil }

type NATPolicy struct { ID string `json:"id"`; SourcePrefix string `json:"source_prefix"`; EgressInterface string `json:"egress_interface"`; Authorized bool `json:"authorized"` }
func (n NATPolicy) Validate() error { if !n.Authorized||n.ID==""||n.EgressInterface==""{return errors.New("invalid NAT policy")}; if _,e:=netip.ParsePrefix(n.SourcePrefix);e!=nil{return errors.New("invalid NAT source prefix")}; return nil }
