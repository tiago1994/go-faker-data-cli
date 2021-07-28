package generate

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/manveru/faker"
)

func GenerateData(typeFile string, numberOfLines int32) string {
	fullData := buildData(typeFile, numberOfLines)
	return fullData
}

func buildData(typeFile string, numberOfLines int32) string {
	var data string

	switch typeFile {
	case "txt":
		data = generateTxtData(numberOfLines)
	case "csv":
		data = generateCsv(numberOfLines)
	}

	return data
}

func generateTxtData(numberOfLines int32) string {
	var lines string
	var fileName string

	for i := 0; i < int(numberOfLines); i++ {
		fake, err := faker.New("en")
		if err != nil {
			panic(err)
		}

		lines = lines + fake.Name() + " " + fake.Email() + " " + fake.PhoneNumber() + " " + fake.CellPhoneNumber() + " " + fake.City() + " " + fake.Country() + " " + fake.CompanyName() + " " + fake.CompanyBs() + "\n"
	}

	now := time.Now()
	sec := now.Unix()

	fileName = "exported/txt/new_file_" + strconv.Itoa(int(sec)) + ".txt"
	saveFile(lines, fileName)
	return fileName
}

func generateCsv(numberOfLines int32) string {
	var lines string
	var fileName string

	lines = "name,email,phone,cell_phone,city,country,company,company_bs\n"

	for i := 0; i < int(numberOfLines); i++ {
		fake, err := faker.New("en")
		if err != nil {
			panic(err)
		}

		lines = lines + fake.Name() + "," + fake.Email() + "," + fake.PhoneNumber() + "," + fake.CellPhoneNumber() + "," + fake.City() + "," + fake.Country() + "," + fake.CompanyName() + "," + fake.CompanyBs() + "\n"
	}

	now := time.Now()
	sec := now.Unix()

	fileName = "exported/csv/new_file_" + strconv.Itoa(int(sec)) + ".csv"
	saveFile(lines, fileName)
	return fileName
}

func saveFile(data string, fileName string) {
	err := ioutil.WriteFile(fileName, []byte(data), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
}
