package main

import (
	"errors"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3" //nolint:depguard
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrInvalidParameter      = errors.New("invalid parameter")
)

const (
	chunk = int64(1024)
)

func copySource(source, destination *os.File, limit int64) error {
	var err error
	var chunkSize int64
	bar := pb.Start64(limit)

	for i := int64(0); i < limit; i += chunk {
		if i+chunk > limit {
			chunkSize = limit % chunk
		} else {
			chunkSize = chunk
		}
		_, err = io.CopyN(destination, source, chunkSize)
		bar.Add64(chunkSize)
		if err != nil {
			return err
		}
	}
	bar.Finish()
	return nil
}

func Copy(fromPath, toPath string, offset, limit int64) error {
	if offset < 0 || limit < 0 {
		return ErrInvalidParameter
	}
	source, err := os.OpenFile(fromPath, os.O_RDONLY, 0o644)
	if err != nil {
		return err
	}
	defer source.Close()
	sourceStat, err := source.Stat()
	if err != nil {
		return err
	}
	if sourceStat.IsDir() {
		return ErrUnsupportedFile
	}
	sourceSize := sourceStat.Size()
	if offset >= sourceSize {
		return ErrOffsetExceedsFileSize
	}
	destination, err := os.OpenFile(toPath, os.O_CREATE|os.O_WRONLY, sourceStat.Mode())
	if err != nil {
		// Не удалось создать файл-копию:
		return err
	}
	defer destination.Close()

	_, err = source.Seek(offset, 0)
	if err != nil {
		return err
	}
	if limit > sourceSize-offset || limit == 0 {
		limit = sourceSize - offset
	}
	return copySource(source, destination, limit)
}
