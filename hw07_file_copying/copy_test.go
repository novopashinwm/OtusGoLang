package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	tests := []struct {
		name        string
		offset      int64
		limit       int64
		fileIn      string
		fileCheck   string
		expectedErr error
	}{
		{
			name:        "error offset file size",
			offset:      10000,
			limit:       0,
			fileIn:      "testdata/input.txt",
			fileCheck:   "",
			expectedErr: ErrOffsetExceedsFileSize,
		},
		{
			name:        "error - not supported file",
			offset:      0,
			limit:       0,
			fileIn:      "testdata",
			fileCheck:   "",
			expectedErr: ErrUnsupportedFile,
		},
		{
			name:        "offset 0 limit 0",
			offset:      0,
			limit:       0,
			fileIn:      "testdata/input.txt",
			fileCheck:   "out_offset0_limit0.txt",
			expectedErr: nil,
		},
		{
			name:        "offset 0 limit 10",
			offset:      0,
			limit:       10,
			fileIn:      "testdata/input.txt",
			fileCheck:   "out_offset0_limit10.txt",
			expectedErr: nil,
		},
		{
			name:        "offset 0 limit 1000",
			offset:      0,
			limit:       1000,
			fileIn:      "testdata/input.txt",
			fileCheck:   "out_offset0_limit1000.txt",
			expectedErr: nil,
		},
		{
			name:        "offset 0 limit 10000",
			offset:      0,
			limit:       10000,
			fileIn:      "testdata/input.txt",
			fileCheck:   "out_offset0_limit10000.txt",
			expectedErr: nil,
		},

		{
			name:        "offset 100 limit 1000",
			offset:      100,
			limit:       1000,
			fileIn:      "testdata/input.txt",
			fileCheck:   "out_offset100_limit1000.txt",
			expectedErr: nil,
		},

		{
			name:        "offset 6000 limit 1000",
			offset:      6000,
			limit:       1000,
			fileIn:      "testdata/input.txt",
			fileCheck:   "out_offset6000_limit1000.txt",
			expectedErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("case %s", test.name), func(t *testing.T) {
			test := test
			t.Parallel()
			err := Copy(test.fileIn, "out.txt", test.offset, test.limit)
			if err != nil {
				require.Truef(t, errors.Is(err, test.expectedErr), "actual err - %v", err)
			} else {
				check := deepCompare("testdata/"+test.fileCheck, "out.txt")
				require.Truef(t, check, "actual err - %v", err)
			}
		})
	}
}

func deepCompare(file1, file2 string) bool {
	sf, err := os.Open(file1)
	if err != nil {
		log.Fatal(err)
	}

	df, err := os.Open(file2)
	if err != nil {
		log.Fatal(err)
	}

	sscan := bufio.NewScanner(sf)
	dscan := bufio.NewScanner(df)

	for sscan.Scan() {
		dscan.Scan()
		if !bytes.Equal(sscan.Bytes(), dscan.Bytes()) {
			sf.Close()
			df.Close()
			return false
		}
	}
	sf.Close()
	df.Close()

	return true
}
