package importer

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

const csvHeader = "ip_address,country_code,country,city,latitude,longitude,mystery_value"

func TestCSVImporter_Import(t *testing.T) {
	type args struct {
		fileContent string
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		expected *CSVRows
	}{
		{
			name: "all success",
			args: args{
				fileContent: csvHeader + "\n" +
					"192.184.51.218,RU,Morocco,Willburgh,76.7892707471672,-8.617777079132821,2815330924\n" +
					"160.168.85.54,BO,Cuba,Mohamedview,-66.20896958745531,81.62948730878543,8879434387\n" +
					"214.165.161.44,PR,Republic of Korea,West Erika,48.92021642445653,14.900399560492929,3829378711\n" +
					"187.197.68.39,MM,Zimbabwe,North Anamouth,48.82685320435576,2.9300655090904684,8762174736",
			},
			wantErr: false,
			expected: &CSVRows{
				rows: []CSVRow{
					{
						IPAddress:    "192.184.51.218",
						CountryCode:  "RU",
						Country:      "Morocco",
						City:         "Willburgh",
						Latitude:     76.7892707471672,
						Longitude:    -8.617777079132821,
						MysteryValue: "2815330924",
					},
					{
						IPAddress:    "160.168.85.54",
						CountryCode:  "BO",
						Country:      "Cuba",
						City:         "Mohamedview",
						Latitude:     -66.20896958745531,
						Longitude:    81.62948730878543,
						MysteryValue: "8879434387",
					},
					{
						IPAddress:    "214.165.161.44",
						CountryCode:  "PR",
						Country:      "Republic of Korea",
						City:         "West Erika",
						Latitude:     48.92021642445653,
						Longitude:    14.900399560492929,
						MysteryValue: "3829378711",
					},
					{
						IPAddress:    "187.197.68.39",
						CountryCode:  "MM",
						Country:      "Zimbabwe",
						City:         "North Anamouth",
						Latitude:     48.82685320435576,
						Longitude:    2.9300655090904684,
						MysteryValue: "8762174736",
					},
				},
				rowsDiscardedCount: 0,
			},
		},
		{
			name: "2 failed, cannot unmarshal csv",
			args: args{
				fileContent: csvHeader + "\n" +
					"192.184.51.218,RU,Morocco,Willburgh,76.7892707471672,-8.617777079132821,2815330924\n" +
					"51.23.171.108,SK,Slovakia (Slovak Republic),Mrazview,,,0\n" +
					"114.242.56.225,SX,Brunei Darussalam,Lake Jamirville,,,0",
			},
			wantErr: false,
			expected: &CSVRows{
				rows: []CSVRow{
					{
						IPAddress:    "192.184.51.218",
						CountryCode:  "RU",
						Country:      "Morocco",
						City:         "Willburgh",
						Latitude:     76.7892707471672,
						Longitude:    -8.617777079132821,
						MysteryValue: "2815330924",
					},
				},
				rowsDiscardedCount: 2,
			},
		},
		{
			name: "incorrect file contents",
			args: args{
				fileContent: "",
			},
			wantErr:  true,
			expected: &CSVRows{},
		},
		{
			name: "invalid ip address",
			args: args{
				fileContent: csvHeader + "\n" +
					"192.184.51.2181,RU,Morocco,Willburgh,76.7892707471672,-8.617777079132821,2815330924\n" +
					"192.184.51.12sa,RU,Morocco,Willburgh,76.7892707471672,-8.617777079132821,2815330924\n",
			},
			wantErr: false,
			expected: &CSVRows{
				rowsDiscardedCount: 2,
			},
		},
		{
			name: "empty city name",
			args: args{
				fileContent: csvHeader + "\n" +
					"192.184.51.2181,RU,Morocco,,76.7892707471672,-8.617777079132821,2815330924\n",
			},
			wantErr: false,
			expected: &CSVRows{
				rowsDiscardedCount: 1,
			},
		},
		{
			name: "empty country code",
			args: args{
				fileContent: csvHeader + "\n" +
					"192.184.51.2181,,Morocco,Willburgh,76.7892707471672,-8.617777079132821,2815330924\n",
			},
			wantErr: false,
			expected: &CSVRows{
				rowsDiscardedCount: 1,
			},
		},
		{
			name: "empty country name",
			args: args{
				fileContent: csvHeader + "\n" +
					"192.184.51.2181,RU,,Willburgh,76.7892707471672,-8.617777079132821,2815330924\n",
			},
			wantErr: false,
			expected: &CSVRows{
				rowsDiscardedCount: 1,
			},
		},
		{
			name: "wrong coordinates",
			args: args{
				fileContent: csvHeader + "\n" +
					"192.184.51.2181,RU,Morocco,Willburgh,91,-8.617777079132821,2815330924\n" +
					"192.184.51.2181,RU,Morocco,Willburgh,-91,-8.617777079132821,2815330924\n" +
					"192.184.51.2181,RU,Morocco,Willburgh,76.7892707471672,181,2815330924\n" +
					"192.184.51.2181,RU,Morocco,Willburgh,76.7892707471672,-180.34,2815330924\n",
			},
			wantErr: false,
			expected: &CSVRows{
				rowsDiscardedCount: 4,
			},
		},
		{
			name: "ip uniqueness",
			args: args{
				fileContent: csvHeader + "\n" +
					"192.184.51.218,RU,Morocco,Willburgh,76.7892707471672,-8.617777079132821,2815330924\n" +
					"192.184.51.218,RU,Moroccosdfsdf,Willburgh,76.7892707471672,-8.617777079132821,2815330924\n",
			},
			wantErr: false,
			expected: &CSVRows{
				rows: []CSVRow{
					{
						IPAddress:    "192.184.51.218",
						CountryCode:  "RU",
						Country:      "Morocco",
						City:         "Willburgh",
						Latitude:     76.7892707471672,
						Longitude:    -8.617777079132821,
						MysteryValue: "2815330924",
					},
				},
				rowsDiscardedCount: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CSVImporter := &CSVImporter{}
			rows := &CSVRows{}
			buf := &bytes.Buffer{}
			buf.WriteString(tt.args.fileContent)

			err := CSVImporter.Import(buf, rows)
			require.Equal(t, tt.wantErr, err != nil)
			require.Equal(t, tt.expected, rows)
			require.Equal(t, tt.expected.rows, rows.GetRows())
			require.Equal(t, tt.expected.rowsDiscardedCount, rows.GetDiscardedCnt())
		})
	}
}
