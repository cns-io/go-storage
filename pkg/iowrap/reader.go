package iowrap

import (
	"io"
)

//go:generate mockgen -package iowrap -destination mock_test.go io Reader,Closer,ReaderAt,Seeker

// LimitReadCloser will return a limited hasCall closer.
func LimitReadCloser(r io.ReadCloser, n int64) *LimitedReadCloser {
	return &LimitedReadCloser{r, io.LimitReader(r, n)}
}

// LimitedReadCloser hasCall from underlying r and provide Close as well.
type LimitedReadCloser struct {
	r  io.ReadCloser
	lr io.Reader
}

// Read is copied from io.LimitedReader's Read.
func (l *LimitedReadCloser) Read(p []byte) (n int, err error) {
	return l.lr.Read(p)
}

// Close will close underlying reader.
func (l *LimitedReadCloser) Close() error {
	return l.r.Close()
}

// SectionReadCloser will return a sectioned hasCall closer.
func SectionReadCloser(r interface {
	io.Closer
	io.ReaderAt
}, off, n int64) *SectionedReadCloser {
	return &SectionedReadCloser{r, io.NewSectionReader(r, off, n)}
}

// SectionedReadCloser hasCall from underlying r and provide Close as well.
type SectionedReadCloser struct {
	r interface {
		io.Closer
		io.ReaderAt
	}
	sr *io.SectionReader
}

// Read is copied from io.SectionReader's Read.
func (s *SectionedReadCloser) Read(p []byte) (n int, err error) {
	return s.sr.Read(p)
}

// Close will close underlying reader.
func (s *SectionedReadCloser) Close() error {
	return s.r.Close()
}

// ReadSeekCloser wraps a io.Reader returning a SeekCloseableReader. Allows the
// SDK to accept an io.Reader that is not also an io.Seeker for unsigned
// streaming payload API operations.
//
// A ReadSeekCloser wrapping an nonseekable io.Reader used in an API
// operation's input will prevent that operation being retried in the case of
// network errors, and cause operation requests to fail if the operation
// requires payload signing.
//
// NOTES: Idea borrows from AWS Go SDK.
func ReadSeekCloser(r io.Reader) SeekCloseableReader {
	return SeekCloseableReader{r}
}

// SeekCloseableReader represents a reader that can also delegate io.Seeker and
// io.Closer interfaces to the underlying object if they are available.
type SeekCloseableReader struct {
	r io.Reader
}

// Read reads from the reader up to size of p. The number of bytes read, and
// error if it occurred will be returned.
//
// If the reader is not an io.Reader zero bytes read, and nil error will be
// returned.
//
// Performs the same functionality as io.Reader Read
func (r SeekCloseableReader) Read(p []byte) (int, error) {
	return r.r.Read(p)
}

// Seek sets the offset for the next Read to offset, interpreted according to
// whence: 0 means relative to the origin of the file, 1 means relative to the
// current offset, and 2 means relative to the end. Seek returns the new offset
// and an error, if any.
//
// If the SeekCloseableReader is not an io.Seeker nothing will be done.
func (r SeekCloseableReader) Seek(offset int64, whence int) (int64, error) {
	switch t := r.r.(type) {
	case io.Seeker:
		return t.Seek(offset, whence)
	}
	return int64(0), nil
}

// Close closes the SeekCloseableReader.
//
// If the SeekCloseableReader is not an io.Closer nothing will be done.
func (r SeekCloseableReader) Close() error {
	switch t := r.r.(type) {
	case io.Closer:
		return t.Close()
	}
	return nil
}

func CallbackReader(r io.Reader, fn func([]byte)) *CallbackifyReader {
	return &CallbackifyReader{
		r:  r,
		fn: fn,
	}
}

type CallbackifyReader struct {
	r  io.Reader
	fn func([]byte)
}

func (r CallbackifyReader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	r.fn(p[:n])
	return n, err
}

func CallbackReadCloser(r io.ReadCloser, fn func([]byte)) *CallbackifyReadCloser {
	return &CallbackifyReadCloser{
		r:  r,
		fn: fn,
	}
}

type CallbackifyReadCloser struct {
	r  io.ReadCloser
	fn func([]byte)
}

func (r CallbackifyReadCloser) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	r.fn(p[:n])
	return n, err
}

func (r CallbackifyReadCloser) Close() error {
	return r.r.Close()
}
