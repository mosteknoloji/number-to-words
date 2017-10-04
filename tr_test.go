package ntw

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func ExampleIntegerToTurkish() {
	fmt.Println(IntegerToTurkish(42))
	// Output: kırk iki
}

func TestIntegerToTurkish(t *testing.T) {
	Convey("Testing IntegerToTurkish()", t, FailureContinues, func() {
		testing := map[int]string{
			0:         "sıfır",
			1:         "bir",
			9:         "dokuz",
			10:        "on",
			11:        "on bir",
			19:        "on dokuz",
			20:        "yirmi",
			21:        "yirmi bir",
			80:        "seksen",
			90:        "doksan",
			99:        "doksan dokuz",
			100:       "yüz",
			101:       "yüz bir",
			111:       "yüz on bir",
			120:       "yüz yirmi",
			121:       "yüz yirmi bir",
			900:       "dokuz yüz",
			909:       "dokuz yüz dokuz",
			919:       "dokuz yüz on dokuz",
			990:       "dokuz yüz doksan",
			999:       "dokuz yüz doksan dokuz",
			1000:      "bin",
			2000:      "iki bin",
			4000:      "dört bin",
			5000:      "beş bin",
			11000:     "on bir bin",
			21000:     "yirmi bir bin",
			999000:    "dokuz yüz doksan dokuz bin",
			999999:    "dokuz yüz doksan dokuz bin dokuz yüz doksan dokuz",
			1000000:   "bir milyon",
			2000000:   "iki milyon",
			4000000:   "dört milyon",
			5000000:   "beş milyon",
			100100100: "yüz milyon yüz bin yüz",
			500500500: "beş yüz milyon beş yüz bin beş yüz",
			606606606: "altı yüz altı milyon altı yüz altı bin altı yüz altı",
			/*
				999000000:     "nine hundred ninety-nine million",
				999000999:     "nine hundred ninety-nine million nine hundred ninety-nine",
				999999000:     "nine hundred ninety-nine million nine hundred ninety-nine thousand",
				999999999:     "nine hundred ninety-nine million nine hundred ninety-nine thousand nine hundred ninety-nine",
				1174315110:    "one billion one hundred seventy-four million three hundred fifteen thousand one hundred ten",
				1174315119:    "one billion one hundred seventy-four million three hundred fifteen thousand one hundred nineteen",
				15174315119:   "fifteen billion one hundred seventy-four million three hundred fifteen thousand one hundred nineteen",
				35174315119:   "thirty-five billion one hundred seventy-four million three hundred fifteen thousand one hundred nineteen",
				935174315119:  "nine hundred thirty-five billion one hundred seventy-four million three hundred fifteen thousand one hundred nineteen",
				2935174315119: "two trillion nine hundred thirty-five billion one hundred seventy-four million three hundred fifteen thousand one hundred nineteen",
			*/
		}
		for input, expectedOutput := range testing {
			So(IntegerToTurkish(input), ShouldEqual, expectedOutput)
		}

		// testing negative values
		So(IntegerToTurkish(-1), ShouldEqual, "eksi bir")
	})
}
