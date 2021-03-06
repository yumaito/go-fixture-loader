package loader

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/pkg/errors"
)

func (fx FixtureLoader) getDataFromCSV(file, format string) (Data, error) {
	f, err := os.Open(file)
	if err != nil {
		err = errors.Wrapf(err, "file: %s open error", file)
		return Data{}, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	if format == "tsv" {
		reader.Comma = '\t'
	}

	columns, err := reader.Read()
	if err != nil {
		err = errors.Wrapf(err, "file: %s read error", file)
		return Data{}, err
	}

	data := Data{columns: columns}
	for {
		row, err := reader.Read()

		if err != nil {
			if err == io.EOF {
				break
			}
			return data, err
		}

		rows := make(map[string]string, len(row))
		for i, value := range row {
			rows[data.columns[i]] = value
		}

		data.rows = append(data.rows, rows)
	}

	return data, nil
}
