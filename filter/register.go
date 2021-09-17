package filter

var (
	// GlobalGroups is the middle stage of filtering environment variables, this is for filtering on a per-key basis
	// GlobalFilterGroups is a map, and a given filter will only be applied by Environment if the incoming
	// key matches the key used in the map.
	GlobalGroups = make(GroupMap)

	// GlobalEntranceGroup is the first stage of filtering environment variables, this is for filtering all variables
	// regardless of their key.
	GlobalEntranceGroup = make(Group, 0)

	// GlobalExitGroup is the final stage of filtering environment variables, like GlobalEntranceGroup it
	GlobalExitGroup = make(Group, 0)
)
