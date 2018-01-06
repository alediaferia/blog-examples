package main

import (
	"flag"
	"fmt"
	"github.com/alediaferia/prefixmap"
	"math"
	"os"
	"strings"
)

var Countries = []string{
	"Afghanistan",
	"Albania",
	"Algeria",
	"American Samoa",
	"Andorra",
	"Angola",
	"Anguilla",
	"Antarctica",
	"Antigua and Barbuda",
	"Argentina",
	"Armenia",
	"Aruba",
	"Australia",
	"Austria",
	"Azerbaijan",
	"Bahamas",
	"Bahrain",
	"Bangladesh",
	"Barbados",
	"Belarus",
	"Belgium",
	"Belize",
	"Benin",
	"Bermuda",
	"Bhutan",
	"Bolivia",
	"Bosnia and Herzegovina",
	"Botswana",
	"Bouvet Island",
	"Brazil",
	"British Antarctic Territory",
	"British Indian Ocean Territory",
	"British Virgin Islands",
	"Brunei",
	"Bulgaria",
	"Burkina Faso",
	"Burundi",
	"Cambodia",
	"Cameroon",
	"Canada",
	"Canton and Enderbury Islands",
	"Cape Verde",
	"Cayman Islands",
	"Central African Republic",
	"Chad",
	"Chile",
	"China",
	"Christmas Island",
	"Cocos [Keeling] Islands",
	"Colombia",
	"Comoros",
	"Congo - Brazzaville",
	"Congo - Kinshasa",
	"Cook Islands",
	"Costa Rica",
	"Croatia",
	"Cuba",
	"Cyprus",
	"Czech Republic",
	"Côte d’Ivoire",
	"Denmark",
	"Djibouti",
	"Dominica",
	"Dominican Republic",
	"Dronning Maud Land",
	"East Germany",
	"Ecuador",
	"Egypt",
	"El Salvador",
	"Equatorial Guinea",
	"Eritrea",
	"Estonia",
	"Ethiopia",
	"Falkland Islands",
	"Faroe Islands",
	"Fiji",
	"Finland",
	"France",
	"French Guiana",
	"French Polynesia",
	"French Southern Territories",
	"French Southern and Antarctic Territories",
	"Gabon",
	"Gambia",
	"Georgia",
	"Germany",
	"Ghana",
	"Gibraltar",
	"Greece",
	"Greenland",
	"Grenada",
	"Guadeloupe",
	"Guam",
	"Guatemala",
	"Guernsey",
	"Guinea",
	"Guinea-Bissau",
	"Guyana",
	"Haiti",
	"Heard Island and McDonald Islands",
	"Honduras",
	"Hong Kong SAR China",
	"Hungary",
	"Iceland",
	"India",
	"Indonesia",
	"Iran",
	"Iraq",
	"Ireland",
	"Isle of Man",
	"Israel",
	"Italy",
	"Jamaica",
	"Japan",
	"Jersey",
	"Johnston Island",
	"Jordan",
	"Kazakhstan",
	"Kenya",
	"Kiribati",
	"Kuwait",
	"Kyrgyzstan",
	"Laos",
	"Latvia",
	"Lebanon",
	"Lesotho",
	"Liberia",
	"Libya",
	"Liechtenstein",
	"Lithuania",
	"Luxembourg",
	"Macau SAR China",
	"Macedonia",
	"Madagascar",
	"Malawi",
	"Malaysia",
	"Maldives",
	"Mali",
	"Malta",
	"Marshall Islands",
	"Martinique",
	"Mauritania",
	"Mauritius",
	"Mayotte",
	"Metropolitan France",
	"Mexico",
	"Micronesia",
	"Midway Islands",
	"Moldova",
	"Monaco",
	"Mongolia",
	"Montenegro",
	"Montserrat",
	"Morocco",
	"Mozambique",
	"Myanmar [Burma]",
	"Namibia",
	"Nauru",
	"Nepal",
	"Netherlands",
	"Netherlands Antilles",
	"Neutral Zone",
	"New Caledonia",
	"New Zealand",
	"Nicaragua",
	"Niger",
	"Nigeria",
	"Niue",
	"Norfolk Island",
	"North Korea",
	"North Vietnam",
	"Northern Mariana Islands",
	"Norway",
	"Oman",
	"Pacific Islands Trust Territory",
	"Pakistan",
	"Palau",
	"Palestinian Territories",
	"Panama",
	"Panama Canal Zone",
	"Papua New Guinea",
	"Paraguay",
	"People's Democratic Republic of Yemen",
	"Peru",
	"Philippines",
	"Pitcairn Islands",
	"Poland",
	"Portugal",
	"Puerto Rico",
	"Qatar",
	"Romania",
	"Russia",
	"Rwanda",
	"Réunion",
	"Saint Barthélemy",
	"Saint Helena",
	"Saint Kitts and Nevis",
	"Saint Lucia",
	"Saint Martin",
	"Saint Pierre and Miquelon",
	"Saint Vincent and the Grenadines",
	"Samoa",
	"San Marino",
	"Saudi Arabia",
	"Senegal",
	"Serbia",
	"Serbia and Montenegro",
	"Seychelles",
	"Sierra Leone",
	"Singapore",
	"Slovakia",
	"Slovenia",
	"Solomon Islands",
	"Somalia",
	"South Africa",
	"South Georgia and the South Sandwich Islands",
	"South Korea",
	"Spain",
	"Sri Lanka",
	"Sudan",
	"Suriname",
	"Svalbard and Jan Mayen",
	"Swaziland",
	"Sweden",
	"Switzerland",
	"Syria",
	"São Tomé and Príncipe",
	"Taiwan",
	"Tajikistan",
	"Tanzania",
	"Thailand",
	"Timor-Leste",
	"Togo",
	"Tokelau",
	"Tonga",
	"Trinidad and Tobago",
	"Tunisia",
	"Turkey",
	"Turkmenistan",
	"Turks and Caicos Islands",
	"Tuvalu",
	"U.S. Minor Outlying Islands",
	"U.S. Miscellaneous Pacific Islands",
	"U.S. Virgin Islands",
	"Uganda",
	"Ukraine",
	"Union of Soviet Socialist Republics",
	"United Arab Emirates",
	"United Kingdom",
	"United States",
	"Unknown or Invalid Region",
	"Uruguay",
	"Uzbekistan",
	"Vanuatu",
	"Vatican City",
	"Venezuela",
	"Vietnam",
	"Wake Island",
	"Wallis and Futuna",
	"Western Sahara",
	"Yemen",
	"Zambia",
	"Zimbabwe",
	"Åland Islands",
}

var similarity float64
var datasource *prefixmap.PrefixMap

type Match struct {
	Value      string
	Similarity float64
}

func (match *Match) Print() {
	fmt.Printf("match: \t%s\tsimilarity: %.2f\t\n", match.Value, match.Similarity)
}

func init() {
	flag.Float64Var(&similarity, "similarity", 0.3, "the similarity target to use when searching")

	datasource = prefixmap.New()
}

func main() {
	flag.Parse()

	input := flag.Arg(0)
	if input == "" {
		fmt.Println("Please, specify an input string")
		os.Exit(1)
	}

	// here we populate the datasource
	for _, country := range Countries {
		parts := strings.Split(strings.ToLower(country), " ")
		for _, part := range parts {
			datasource.Insert(part, country)
		}
	}

	values := datasource.GetByPrefix(strings.ToLower(input))
	results := make([]*Match, 0, len(values))
	for _, v := range values {
		value := v.(string)
		s := ComputeSimilarity(len(value), len(input), LevenshteinDistance(value, input))
		if s >= similarity {
			m := &Match{value, s}
			results = append(results, m)
		}
	}

	fmt.Printf("Result for target similarity: %.2f\n", similarity)
	PrintMatches(results)
}

func PrintMatches(matches []*Match) {
	for _, m := range matches {
		m.Print()
	}
}

func ComputeSimilarity(w1Len, w2Len, ld int) float64 {
	maxLen := math.Max(float64(w1Len), float64(w2Len))

	return 1.0 - float64(ld)/float64(maxLen)
}

func LevenshteinDistance(source, destination string) int {
	vec1 := make([]int, len(destination)+1)
	vec2 := make([]int, len(destination)+1)

	w1 := []rune(source)
	w2 := []rune(destination)

	// initializing vec1
	for i := 0; i < len(vec1); i++ {
		vec1[i] = i
	}

	// initializing the matrix
	for i := 0; i < len(w1); i++ {
		vec2[0] = i + 1

		for j := 0; j < len(w2); j++ {
			cost := 1
			if w1[i] == w2[j] {
				cost = 0
			}
			min := minimum(vec2[j]+1,
				vec1[j+1]+1,
				vec1[j]+cost)
			vec2[j+1] = min
		}

		for j := 0; j < len(vec1); j++ {
			vec1[j] = vec2[j]
		}
	}

	return vec2[len(w2)]
}

func minimum(value0 int, values ...int) int {
	min := value0
	for _, v := range values {
		if v < min {
			min = v
		}
	}
	return min
}
