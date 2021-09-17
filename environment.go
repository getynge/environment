package environment

import (
    "fmt"
    "github.com/getynge/environment/filter"
    "os"
    "strings"
)

type Environment struct {
    m map[string]string
}

// New creates an empty Environment
func New() (e Environment) {
    e = Environment{
        m:       make(map[string]string),
    }

    return e
}

// Shell creates a new Environment, with the environment variables of the current process added to it
func Shell() (e Environment) {
    e = New()

    for _, pair := range os.Environ() {
        var v string
        kv := strings.Split(pair, "=")
        k := kv[0]

        if len(kv) >= 2 {
            v = kv[1]
        }

        e.m[k] = v
    }

    return e
}

func (e Environment) String() string {
    b := new(strings.Builder)
    b.WriteString("export")
    for k, v := range e.m {
        b.WriteString(fmt.Sprintf(" %s=%s", k, v))
    }
    return b.String()
}

// Set sets the given key to the given value, after running all filters on the key value pair.
// If any of the filters fail, the variable is not added to the Environment and an error is returned
func (e Environment) Set(key, value string) (err error) {
    filters, _ := filter.GlobalFilters[key]

    for _, f := range filters {
        key, value, err = f.Filter(key, value)
        if err != nil {
            return err
        }
    }

    e.m[key] = value
    return err
}

func (e Environment) Get(key string) (variable string, has bool) {
    variable, has = e.m[key]
    return variable, has
}

func (e Environment) Remove(key string) {
    delete(e.m, key)
}