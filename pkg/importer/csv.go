package importer

import (
	csv2 "encoding/csv"
	"io"
	"net"

	"github.com/friendsofgo/errors"
	"github.com/jszwec/csvutil"
)

type CSVRow struct {
	IPAddress    string  `csv:"ip_address"`
	CountryCode  string  `csv:"country_code"`
	Country      string  `csv:"country"`
	City         string  `csv:"city"`
	Latitude     float64 `csv:"latitude"`
	Longitude    float64 `csv:"longitude"`
	MysteryValue string  `csv:"mystery_value"`
}

type CSVRows struct {
	rows               []CSVRow
	rowsDiscardedCount int
}

type CSVImporter struct{}

func GetCSVImporter() Importer {
	return &CSVImporter{}
}

// Import imports csv data from source to provided rows
func (CSVImporter *CSVImporter) Import(source io.Reader, rows ImportedRows) error {
	csvReader := csv2.NewReader(source)
	decoder, err := csvutil.NewDecoder(csvReader)
	if err != nil {
		return errors.Wrap(err, "create decoder")
	}

	uniquenessMap := make(map[string]struct{})
	for {
		row := CSVRow{}
		if err := decoder.Decode(&row); err == io.EOF {
			break
		} else if err != nil {
			rows.IncrementDiscardedCnt()
			continue
		}
		if !row.IsValid() {
			rows.IncrementDiscardedCnt()
			continue
		}
		if _, exists := uniquenessMap[row.IPAddress]; exists {
			rows.IncrementDiscardedCnt()
			continue
		}
		uniquenessMap[row.IPAddress] = struct{}{}
		err = rows.addRow(row)
		if err != nil {
			return errors.Wrap(err, "failed to add row")
		}
	}

	return nil
}

func (csvRows *CSVRows) addRow(row interface{}) error {
	csvRow, isCorrectType := row.(CSVRow)
	if !isCorrectType {
		return errors.New("incorrect csv row type")
	}
	csvRows.rows = append(csvRows.rows, csvRow)

	return nil
}

// GetRows returns underlying rows collection
func (csvRows *CSVRows) GetRows() []CSVRow {
	return csvRows.rows
}

func (csvRows *CSVRows) GetDiscardedCnt() int {
	return csvRows.rowsDiscardedCount
}

func (csvRows *CSVRows) IncrementDiscardedCnt() {
	csvRows.rowsDiscardedCount++
}

// IsValid checks if row satisfies all restrictions
func (csvRow *CSVRow) IsValid() bool {
	if net.ParseIP(csvRow.IPAddress) == nil {
		return false
	}
	if len(csvRow.City) == 0 ||
		len(csvRow.Country) == 0 ||
		len(csvRow.CountryCode) == 0 {
		return false
	}

	if csvRow.Longitude < -180 || csvRow.Longitude > 180 || csvRow.Latitude < -90 || csvRow.Latitude > 90 {
		return false
	}

	return true
}
