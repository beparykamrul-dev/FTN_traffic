package health

type Status struct { Name string `json:"name"`; Healthy bool `json:"healthy"`; Detail string `json:"detail,omitempty"` }

func OK(name string) Status { return Status{Name:name, Healthy:true} }
