package lf

import (
	"encoding/binary"
	"io"
	"math"
	"time"
)

// TimeMs returns the time in milliseconds since epoch.
func TimeMs() uint64 { return uint64(time.Now().UnixNano()) / uint64(1000000) }

// TimeMsToTime converts a time in milliseconds since epoch to a Go native time.Time structure.
func TimeMsToTime(ms uint64) time.Time { return time.Unix(int64(ms/1000), int64((ms%1000)*1000000)) }

// byteAndArrayReader wraps io.Reader to also make it support io.ByteReader because you can't read a varint from a Reader because derp
type byteAndArrayReader struct{ r io.Reader }

func (mr byteAndArrayReader) Read(p []byte) (int, error) { return mr.r.Read(p) }

func (mr byteAndArrayReader) ReadByte() (byte, error) {
	var tmp [1]byte
	_, err := io.ReadFull(mr.r, tmp[:])
	return tmp[0], err
}

func writeUVarint(out io.Writer, v uint64) (int, error) {
	var tmp [10]byte
	l := binary.PutUvarint(tmp[:], v)
	return out.Write(tmp[0:l])
}

// IntegerSqrtRounded computes the rounded integer square root of a 32-bit unsigned int.
func IntegerSqrtRounded(op uint32) (res uint32) {
	// Just doing this is faster on most platforms and if your float sqrt or round are bad enough for this to be
	// inconsistent with the integer version your platform has personal problems.
	res = uint32(math.Round(math.Sqrt(float64(op))))
	/*
		// Translated from C at https://stackoverflow.com/questions/1100090/looking-for-an-efficient-integer-square-root-algorithm-for-arm-thumb2
		one := uint32(1 << 30)
		for one > op {
			one >>= 2
		}
		for one != 0 {
			if op >= (res + one) {
				op = op - (res + one)
				res = res + 2*one
			}
			res >>= 1
			one >>= 2
		}
		if op > res { // rounding
			res++
		}
	*/
	return
}