package filter

// Identity is a filter which does nothing and returns what was passed to it
type Identity struct {}

func (Identity) Filter(keyIn, valueIn string) (keyOut, valueOut string, err error) {
    return keyIn, valueIn, nil
}
