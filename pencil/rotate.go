package pencil

import (
	"io"
	"time"

	rotate "github.com/lestrrat-go/file-rotatelogs"
)

func NewRotateTime(filePath string, rotationSize, rotationTime, maxAge int64) io.Writer {
	writer, _ := rotate.New(
		filePath,
		rotate.WithMaxAge(time.Duration(maxAge)*time.Second),
		rotate.WithRotationSize(rotationSize),
		rotate.WithRotationTime(time.Duration(rotationTime)*time.Second),
	)
	// Success
	return writer
}

func NewRotateCount(filePath string, rotationSize, rotationTime int64, rotationCount uint) io.Writer {
	writer, _ := rotate.New(
		filePath,
		rotate.WithRotationSize(rotationSize),
		rotate.WithRotationTime(time.Duration(rotationTime)*time.Second),
		rotate.WithRotationCount(rotationCount),
	)
	// Success
	return writer
}
