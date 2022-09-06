package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/MaximChernomorov/challenge-test/internal/repository"
	"github.com/MaximChernomorov/challenge-test/internal/repository/model"
	importerPkg "github.com/MaximChernomorov/challenge-test/pkg/importer"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/types/pgeo"
)

const defaultFilePath = "data_dump.csv"

var importCmd = &cobra.Command{
	Use:  "import",
	Args: cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		importer()
	},
}

var filePath string

func init() {
	importCmd.Flags().StringVarP(
		&filePath,
		"file-path",
		"p",
		defaultFilePath,
		"--file-path=data_dump.csv",
	)
	rootCmd.AddCommand(importCmd)
}

func importer() {
	var err error
	ctx := context.Background()

	repo, err := repository.NewPostgresRepo(viper.GetString("psqURL"))
	cobra.CheckErr(err)
	defer func() {
		err = repo.Close()
		if err != nil {
			cobra.CheckErr(err)
		}
	}()
	start := time.Now()
	sourceFile, err := os.Open(filePath)
	cobra.CheckErr(err)

	rows := &importerPkg.CSVRows{}
	csvImporter := importerPkg.GetCSVImporter()
	err = csvImporter.Import(sourceFile, rows)
	cobra.CheckErr(err)

	geoSlice := getGeoSliceByCSVRows(rows)

	err = repo.AddGeolocationSlice(ctx, geoSlice)
	cobra.CheckErr(err)

	fmt.Println("rows accepted", geoSlice.GetLength())
	fmt.Println("rows discarded", rows.GetDiscardedCnt())
	fmt.Println("time elapsed", time.Since(start))
}

func getGeoSliceByCSVRows(rows *importerPkg.CSVRows) repository.GeolocationSlice {
	geoSlice := make(model.GeolocationSlice, 0)
	for _, row := range rows.GetRows() {
		geoSlice = append(geoSlice, &model.Geolocation{
			IPAddress:    row.IPAddress,
			CountryCode:  null.StringFrom(row.CountryCode),
			Country:      null.StringFrom(row.Country),
			City:         null.StringFrom(row.City),
			Coordinates:  pgeo.NewPoint(row.Latitude, row.Longitude),
			MysteryValue: null.StringFrom(row.MysteryValue),
		})
	}
	return repository.GeolocationSlice{GeolocationSlice: geoSlice}
}
