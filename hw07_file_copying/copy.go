package main

import (
	"errors"
	"io"
	"os"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	fileFrom, err := os.Open(fromPath)
	if err != nil {
		return err
	}

	fileInfo, err := fileFrom.Stat()
	if err != nil {
		return err
	}

	if !fileInfo.Mode().IsRegular() {
		return ErrUnsupportedFile
	}

	fileSize := fileInfo.Size()
	if limit == 0 {
		limit = fileSize
	}

	if offset > fileSize {
		return ErrOffsetExceedsFileSize
	}

	countBytes := minInt64(fileSize-offset, limit)
	fileTo, _ := os.Create(toPath)
	if countBytes == fileSize {
		_, err := io.CopyN(fileTo, fileFrom, countBytes)
		fileTo.Close()
		return err
	}

	readBytes := 0
	arrayIndex := 0
	var currOffset, indexOffset int64 = 0, 0
	readBuf := make([]byte, 1)
	for currOffset < offset {
		_, err := fileFrom.ReadAt(readBuf, indexOffset)
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}
		//if readBuf[0] != 10 {
		currOffset++
		//}
		indexOffset++
	}

	var bar Bar
	bar.NewOption(0, countBytes)
	for int64(arrayIndex) < countBytes {
		read, err := fileFrom.ReadAt(readBuf, indexOffset+int64(readBytes))
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return err
		}
		//if readBuf[0] != 10 {
		arrayIndex++
		bar.Play(int64(arrayIndex))
		//}
		_, err = fileTo.Write(readBuf)
		if err != nil {
			return err
		}

		readBytes += read
	}
	bar.Finish()
	fileTo.Close() // что бы очистить буферы ОС
	fileFrom.Close()
	return nil
}

func minInt64(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
