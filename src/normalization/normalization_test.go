package normalization

import (
	"testing"
)

func TestValidateTimestamp(t *testing.T) {
	type validateTimestampResults struct {
		arg string 
		expectedTimestamp string
		expectedErr bool
	}
 
	testCases := []validateTimestampResults{
		{
			arg: "4/1/11 11:00:00 AM",
			expectedTimestamp: "2011-04-01T14:00:00-04:00", 
			expectedErr: false,
		},
		{
			arg: "11/11/11 11:11:11 AM",
			expectedTimestamp: "2011-11-11T14:11:11-05:00", 
			expectedErr: false,
		},
		{
			arg: "",
			expectedTimestamp: "", 
			expectedErr: true,
		},
		{
			arg: "random string ",
			expectedTimestamp: "", 
			expectedErr: true,
		},
	}
    for _, test := range testCases{

        actualTimestamp, actualErr := validateTimestamp(test.arg)
		if actualTimestamp != test.expectedTimestamp {
            t.Errorf("Actual timestamp %s not equal to expected %s", actualTimestamp, test.expectedTimestamp)
		}

		if (actualErr != nil) != test.expectedErr {
			t.Errorf("Actual %v not equal to expected %t", actualErr, test.expectedErr)
        }
    }

}

func TestFormatZipcode(t *testing.T) {
	type zipCode struct {
		arg string 
		err bool
		expectedZipCode string
	}
	testCases := []zipCode{
		{
			arg: "12345",
			expectedZipCode: "12345", 
			err: false,
		},
		{
			arg: "",
			expectedZipCode: "00000", 
			err: false,
		},
		{
			arg: "9021",
			expectedZipCode: "09021", 
			err: false,
		},
		{
			arg: "invalid zip code test case ",
			expectedZipCode: "", 
			err: true,
		},
	}
    for _, test := range testCases{

        actualZipCode, actualErr := formatZipCode(test.arg)
		if actualZipCode != test.expectedZipCode {
            t.Errorf("Got zipcode %s want zipcode %s", actualZipCode, test.expectedZipCode)
		}
		
		if (actualErr != nil) != test.err {
			t.Errorf("Got err %v. Expected an err: %t", actualErr, test.err)
        }
    }

}

func TestNameToUpperCase(t *testing.T) {
	type name struct {
		arg string 
		expectedName string
	}
	testCases := []name{
		{
			arg: "12345",
			expectedName: "12345", 
		},
		{
			arg: "john smith",
			expectedName: "JOHN SMITH", 
		},
		{
			arg: "haN SoLO",
			expectedName: "HAN SOLO", 
		},
		{
			arg: "Superman übertan",
			expectedName: "SUPERMAN ÜBERTAN", 
		},
	}
    for _, test := range testCases{

        actualName  := nameToUpperCase(test.arg)
		if actualName != test.expectedName {
            t.Errorf("Got name %s want %s", actualName, test.expectedName)
		}

    }

}