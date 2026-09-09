package migration

import "sync"

type idempotencyEntry struct{ fingerprint string; state State }
type Idempotency struct{ mu sync.Mutex; entries map[string]idempotencyEntry }

func NewIdempotency()*Idempotency{return &Idempotency{entries:map[string]idempotencyEntry{}}}
func(i *Idempotency)Claim(key string,fps ...string)bool{fp:=key;if len(fps)>0{fp=fps[0]};if key==""||fp==""{return false};i.mu.Lock();defer i.mu.Unlock();if e,ok:=i.entries[key];ok{if e.state==StateFailed||e.state==StateRollback{delete(i.entries,key)}else{return false}};i.entries[key]=idempotencyEntry{fingerprint:fp,state:StatePreflight};return true}
func(i *Idempotency)Matches(key,fp string)bool{i.mu.Lock();defer i.mu.Unlock();e,ok:=i.entries[key];return ok&&e.fingerprint==fp}
func(i *Idempotency)SetState(key string,s State){i.mu.Lock();defer i.mu.Unlock();if e,ok:=i.entries[key];ok{e.state=s;i.entries[key]=e}}
func(i *Idempotency)State(key string)(State,bool){i.mu.Lock();defer i.mu.Unlock();e,ok:=i.entries[key];return e.state,ok}
func(i *Idempotency)Release(key,fp string){if key==""||fp==""{return};i.mu.Lock();defer i.mu.Unlock();if e,ok:=i.entries[key];ok&&e.fingerprint==fp{delete(i.entries,key)}}
