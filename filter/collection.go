package filter

// Groups is a collection of filter "groups"
//
// Each key corresponds to a single group, and each group consists of every filter for the given key.
// Filters ought to be applied in the order they are added, and the results of one filter are the inputs of the next.
//
// e.g. if there is a filter that adds the letter A to a value, and another that adds B (addA and addB, respectively)
// and both filters are in the same group, then an Environment variable with the value 0 will be equal to 0AB after
// having run through every filter in the given group.
type Groups map[string][]Filter

func (g Groups) AddFilter(key string, filter Filter) {
    _, has := g[key]
    if !has {
        g[key] = make([]Filter, 0)
    }

    g[key] = append(g[key], filter)
}