package dataplane

import (
	"sort"
	"strings"
)

func NormalizeRouteBatch(routes []RouteIntent) []RouteIntent {
	out := append([]RouteIntent(nil), routes...)
	sort.Slice(out, func(i, j int) bool {
		if out[i].Family != out[j].Family { return out[i].Family < out[j].Family }
		if out[i].Prefix != out[j].Prefix { return out[i].Prefix < out[j].Prefix }
		return out[i].NextHop < out[j].NextHop
	})
	for i := range out { for j := range out[i].Community { out[i].Community[j] = strings.TrimSpace(out[i].Community[j]) } }
	return out
}
