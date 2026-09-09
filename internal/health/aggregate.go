package health

func All(statuses ...Status) Status {
	for _, s := range statuses {
		if !s.Healthy { return Status{Name: "aggregate", Healthy: false, Detail: s.Name + ": " + s.Detail} }
	}
	return Status{Name: "aggregate", Healthy: true}
}
