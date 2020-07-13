package adyen

func NewFingerprintString() string {
    fp := NewDF()
    fp.random()
    return fp.String()
}
