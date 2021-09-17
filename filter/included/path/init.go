//go:build !environment_nopath

package path

import "github.com/getynge/environment/filter"

func init() {
	filter.GlobalGroups.AddFilter("PATH", f{})
}
