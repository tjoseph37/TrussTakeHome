package normalization

import(
	"time"
	"fmt"
	"strings"
	"log"
	"unicode/utf8"
	"github.com/tjoseph37/src/models"
	"strconv"
)
 
func NormalizeRecords(records []*models.Records) []*models.Records{
	normalizedRecords := []*models.Records{}

	for index, record := range records{
		if timestamp, err := validateTimestamp(record.Timestamp); err!= nil{
			log.Printf("Skipping record %d. Invalid timestamp error: %v ", index, err)
			continue
		} else {
			record.Timestamp = timestamp
		}

		if zip, err := formatZipCode(record.Zipcode); err!= nil{
			log.Printf("Skipping record %d. Invalid zipcode error: %v ", index, err)
			continue
		} else {
			record.Zipcode = zip
		}

		record.FullName = nameToUpperCase(record.FullName)
		// Add Foo and Bar to total
		foo, err := convertToMiliseconds(record.FooDuration)

		if err!= nil{
			log.Printf("Skipping record %d. Invalid FooDuration error: %v ", index, err)
			continue
		} else {
			record.FooDuration = strconv.FormatInt(foo, 10)
		}

		bar, err := convertToMiliseconds(record.BarDuration)
		if err!= nil{
			log.Printf("Skipping record %d. Invalid FooDuration error: %v ", index, err)
			continue
		} else {
			record.FooDuration = strconv.FormatInt(bar, 10)
		}

		// add if validated
		total := bar + foo
		record.TotalDuration = strconv.FormatInt(total, 10)
		normalizedRecords = append(normalizedRecords, record)
	}

	return normalizedRecords
}

// validate datetime string is RFC3339 compliant and PST -> EST
func validateTimestamp(datetime string) (string, error) {
	// TODO: add custom time stamp validation
	loc, err := time.LoadLocation("America/Los_Angeles")
    if err != nil {
		return "", fmt.Errorf("Error while parsing the date time %v",  err)
    }

	formattedDate, err := time.ParseInLocation("1/2/06 3:4:05 PM", datetime, loc)
	if err != nil {
		return "", fmt.Errorf("Error while parsing the date time %v",  err)
	}

	// set to EST
	loc, err = time.LoadLocation("America/New_York")
	if err != nil {
		return "", fmt.Errorf("Error while parsing the date time %v", err)
	}

	formattedDate = formattedDate.In(loc)

	return formattedDate.Format(time.RFC3339), nil
}

func nameToUpperCase(name string) string {
	return strings.ToUpper(name)
}

func formatZipCode(zipcode string) (string, error) {
	if len(zipcode) < 5 {
		return fmt.Sprintf("%05s", zipcode), nil
	} else if len(zipcode)>5 {
		return "", fmt.Errorf("Error zipcode longer than 5 digits: %s", zipcode)
	}
	return zipcode, nil 
}

func validateUnicode(unicode string) bool {
	return utf8.ValidString(unicode)
}

// TODO: convert foo and bar to secs
func convertToMiliseconds(timestamp string) (int64, error) {
  format := "3:4:05.99"
  formattedTime, err := time.Parse(format, timestamp)

  if err != nil {
	return -1, fmt.Errorf("Error while parsing the date time %v",  err)
  }

  return formattedTime.Unix(), nil
}