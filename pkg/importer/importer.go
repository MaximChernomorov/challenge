package importer

import (
	"io"
)

type ImportedRows interface {
	// GetDiscardedCnt returns count of discarded rows during import
	GetDiscardedCnt() int
	// IncrementDiscardedCnt increases internal counter of discarded rows during import by one
	IncrementDiscardedCnt()

	// addRow appends row to underlying rows collection
	addRow(interface{}) error
}

type Importer interface {
	// Import imports data from source to provided rows
	Import(source io.Reader, rows ImportedRows) error
}
