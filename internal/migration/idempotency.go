package migration

import "sync"

type Idempotency struct { mu sync.Mutex; seen map[string]struct{} }
func NewIdempotency() *Idempotency{return &Idempotency{seen:map[string]struct{}{}}}
func (i *Idempotency) Claim(key string) bool { if key=="" { return false }; i.mu.Lock(); defer i.mu.Unlock(); if _,ok:=i.seen[key];ok{return false}; i.seen[key]=struct{}{};return true }
