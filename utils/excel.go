package utils

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

func CreateExcelTitle(ms []map[string]string) []string {
	colsName := make([]string, len(ms))
	for k, v := range ms {
		colName, _ := v["name"]
		colsName[k] = colName
	}
	return colsName
}

func CreateExcelCols(ms []map[string]string) []string {
	colsName := make([]string, len(ms))
	for k, v := range ms {
		colName, _ := v["key"]
		colsName[k] = colName
	}
	return colsName
}

func CreateExcelFile(titleCols []string, items [][]string) (string, error) {
	buf := new(bytes.Buffer)
	buf.Write([]byte{0xEF, 0xBB, 0xBF})
	r2 := csv.NewWriter(buf)
	r2.Write(titleCols)
	r2.WriteAll(items)
	r2.Flush()
	fileName := fmt.Sprintf("./excelfiles/excel_%d.csv", time.Now().Unix())
	f, err := os.Create(fileName)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	defer f.Close()
	f.Write(buf.Bytes())
	return fileName, nil
}
