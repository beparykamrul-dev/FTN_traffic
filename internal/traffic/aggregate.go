package traffic

type Aggregate struct { Bytes uint64 `json:"bytes"`; Packets uint64 `json:"packets"`; Events uint64 `json:"events"` }

func Add(a Aggregate, e Event) Aggregate { a.Bytes += e.Bytes; a.Packets += e.Packets; a.Events++; return a }
