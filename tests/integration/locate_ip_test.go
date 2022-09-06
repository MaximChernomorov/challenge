package integration

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

type locateIPResponse struct {
	Error       string `json:"error,omitempty"`
	Country     string `json:"country,omitempty"`
	City        string `json:"city,omitempty"`
	Coordinates struct {
		Longitude float64 `json:"longitude,omitempty"`
		Latitude  float64 `json:"latitude,omitempty"`
	} `json:"coordinates,omitempty"`
}

type locateIPRequest struct {
	IPAddress string `json:"ip_address"`
}

// TestLocateIP test relies on data_dump.csv data provided with challenge
func TestLocateIP(t *testing.T) {
	type args struct {
		ip string
	}
	tests := []struct {
		name     string
		args     args
		wantErr  bool
		expected *locateIPResponse
		status   int
	}{
		{
			name: "success Nepal",
			args: args{
				ip: "200.106.141.15",
			},
			wantErr: false,
			expected: &locateIPResponse{
				Country: "Nepal",
				City:    "DuBuquemouth",
				Coordinates: struct {
					Longitude float64 `json:"longitude,omitempty"`
					Latitude  float64 `json:"latitude,omitempty"`
				}{
					Longitude: -84.87503094689836,
					Latitude:  7.206435933364332,
				},
			},
			status: http.StatusOK,
		},
		{
			name: "success Bermuda",
			args: args{
				ip: "33.173.188.44",
			},
			wantErr: false,
			expected: &locateIPResponse{
				Country: "Bermuda",
				City:    "Lake Alexis",
				Coordinates: struct {
					Longitude float64 `json:"longitude,omitempty"`
					Latitude  float64 `json:"latitude,omitempty"`
				}{
					Longitude: 34.26394614136889,
					Latitude:  160.13364176192965,
				},
			},
			status: http.StatusOK,
		},
		{
			name: "invalid IP",
			args: args{
				ip: "33.173.188",
			},
			wantErr: false,
			expected: &locateIPResponse{
				Error: "invalid IP address",
			},
			status: http.StatusBadRequest,
		},
		{
			name: "not found",
			args: args{
				ip: "33.173.188.3",
			},
			wantErr: false,
			expected: &locateIPResponse{
				Error: "location not found",
			},
			status: http.StatusNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, statusCode, err := locateIP(t, tt.args.ip)
			require.Equal(t, tt.wantErr, err != nil)
			require.Equal(t, tt.status, statusCode)
			require.Equal(t, tt.expected, &response)
		})
	}
}

func locateIP(t *testing.T, ip string) (locateIPResponse, int, error) {
	request := locateIPRequest{IPAddress: ip}
	jsonData, err := json.Marshal(request)
	require.Nil(t, err)
	req, err := http.NewRequest("POST", "http://localhost:3011/api/ip/locate", bytes.NewBuffer(jsonData))
	require.Nil(t, err)

	client := &http.Client{}
	resp, err := client.Do(req)
	defer func(Body io.ReadCloser) { Body.Close() }(resp.Body)
	require.Nil(t, err)

	response := locateIPResponse{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	require.Nil(t, err)

	return response, resp.StatusCode, nil
}
