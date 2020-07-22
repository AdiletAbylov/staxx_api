package helpers

import "strconv"

// ProgressPrinter counts the number of bytes written to it. It implements to the io.Writer interface
// and we can pass this into io.TeeReader() which will report progress on each write cycle.
// Uses given func(bytesWrited uint64, bytesTotal uint64) to print progress. Please, keep it mind, that sometimes Total value is unknown.
type ProgressPrinter struct {
	Writed            uint64
	Total             uint64
	ProgressPrintFunc func(bytesWrited uint64, bytesTotal uint64)
}

// NewProgressPrinter returns ProgressPrinter instance with given printer function.
func NewProgressPrinter(progressPrintFunc func(bytesWrited uint64, bytesTotal uint64)) ProgressPrinter {
	return ProgressPrinter{
		ProgressPrintFunc: progressPrintFunc,
	}
}

// SetTotalLengthFromString sets Total value from string
func (pp *ProgressPrinter) SetTotalLengthFromString(s string) {
	if n, err := strconv.Atoi(s); err == nil {
		pp.Total = uint64(n)
	} else {
		pp.Total = 0
	}
}

func (pp *ProgressPrinter) Write(p []byte) (int, error) {
	n := len(p)
	pp.Writed += uint64(n)
	if pp.ProgressPrintFunc != nil {
		pp.ProgressPrintFunc(pp.Writed, pp.Total)
	}

	return n, nil
}
