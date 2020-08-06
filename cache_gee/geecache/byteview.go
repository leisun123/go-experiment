package geecache

// A ByteView holds an immutable view of bytes.
type ByteView struct {
	b []byte
}

// Len returns the view's length
func (v ByteView) Len() int {
	return len(v.b)
}

// ByteSlice returns a copy of the data as a byte slice.
// choose []byte to accept any type of data, eg image, string and etc.
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

// String returns the data as a string, making a copy if necessary.
func (v ByteView) String() string {
	return string(v.b)
}

// b is read-only, use ByteSlice to get a copy
// Prevent cached values from being modified by external programs
func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
