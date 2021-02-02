package base

import "github.com/cockroachdb/errors"

// ErrNotFound means that a get or delete call did not find the requested key.
var ErrNotFound = errors.New("not found")

// ErrCorruption is a marker to indicate that data in a file (WAL, MANIFEST,
// sstable) isn't in the expected format.
var ErrCorruption = errors.New("corruption")

// MarkCorruptionError marks given error as a corruption error.
func MarkCorruptionError(err error) error {
	if errors.Is(err, ErrCorruption) {
		return err
	}
	return errors.Mark(err, ErrCorruption)
}

// CorruptionErrorf formats according to a format specifier and returns
// the string as an error value that is marked as a corruption error.
func CorruptionErrorf(format string, args ...interface{}) error {
	return errors.Mark(errors.Newf(format, args...), ErrCorruption)
}
