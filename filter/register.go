package filter

var (
	// GlobalFilterGroups is the latter half of the collection of filters applied when setting environment variables.
	// GlobalFilterGroups is a map, and a given filter will only be applied by Environment if the incoming
	// key matches the key used in the map.
	GlobalFilterGroups = make(Groups)

	// UniversalFilterGroup is the former half of the collection of filters applied when setting environment variables.
	// Unlike GlobalFilterGroups, UniversalFilterGroup consists of a collection of filters that are applied to all
	// environment variables.
	UniversalFilterGroup = make(Group, 0)
)
