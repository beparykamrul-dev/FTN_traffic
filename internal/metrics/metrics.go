package metrics

type Snapshot struct { Requests uint64 `json:"requests"`; Bytes uint64 `json:"bytes"`; Errors uint64 `json:"errors"` }

func (s *Snapshot) Record(bytes uint64, failed bool) { s.Requests++; s.Bytes += bytes; if failed { s.Errors++ } }
