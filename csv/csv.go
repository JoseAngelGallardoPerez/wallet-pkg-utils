package csv

import (
	"bytes"
	"encoding/csv"
)

const defaultFileName = "file.csv"

// File represents csv file with base operations
type File struct {
	buffer *bytes.Buffer
	writer *csv.Writer
	err    error

	Name string
}

// NewFile returns initialized File
func NewFile() *File {
	buffer := &bytes.Buffer{}
	writer := csv.NewWriter(buffer)
	return &File{buffer, writer, nil, defaultFileName}
}

// WriteRow writes row to buffer and saves error.
// Do nothing if File has error in previous writing
func (f *File) WriteRow(records []string) {
	if f.err == nil {
		f.err = f.writer.Write(records)
	}
}

// Bytes returns bytres for response or saved error. Also retursn error if it happens during Flush for csv writer
func (f *File) Bytes() ([]byte, error) {
	if f.err != nil {
		return nil, f.err
	}

	f.writer.Flush()
	if f.err = f.writer.Error(); f.err != nil {
		return nil, f.err
	}

	return f.buffer.Bytes(), nil
}
