package constant

type Timing struct {
	Prefix string
	Suffix string
}

const (
	CPUS = 4
)

func (t *Timing) GetPrefix() string {
	return t.Prefix
}
func (t *Timing) GetSuffix() string {
	return t.Suffix
}
func (t *Timing) SetPrefix(p string) {
	t.Prefix = p
}
func (t *Timing) SetSuffix(s string) {
	t.Suffix = s
}
