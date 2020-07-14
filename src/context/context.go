package context

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/cmoscofian/meliponto/src/util/constants"
)

var dirname string = path.Join("/", "usr", "local", "etc")
var filename string = "meliponto.json"
var defaultPath string = path.Join(dirname, filename)

func Create() *Configuration {
	content, err := ioutil.ReadFile(defaultPath)
	if err != nil {
		fmt.Println(constants.CorruptedConfigFileError)
		Generate()
		os.Exit(1)
	}

	config := new(Configuration)

	err = json.Unmarshal(content, config)
	if err != nil {
		fmt.Println(constants.ParseConfigError)
		Generate()
		os.Exit(1)
	}

	return config
}

func (c *Configuration) SetCompanyID(companyID string) error {
	c.CompanyID = companyID
	return c.writeToDisk()
}

func (c *Configuration) SetUserID(userID string) error {
	c.UserID = userID
	return c.writeToDisk()
}

func (c *Configuration) writeToDisk() error {
	bs, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(defaultPath, bs, os.ModePerm); err != nil {
		return err
	}

	return nil
}

func Generate() {
	c := createDefaultConfig()

	file, err := os.Create(defaultPath)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t")
	if err := encoder.Encode(c); err != nil {
		log.Fatalln(err)
	}

	fmt.Fprintln(os.Stderr, constants.NewConfigFileGeneratedError)
}

func createDefaultConfig() Configuration {
	return Configuration{
		CompanyID: "a382748",
		Default: DefaultField{
			Hours: DefaultFieldConfig{
				WorkStart:  "09:00",
				LunchStart: "12:00",
				LunchEnd:   "13:00",
				WorkEnd:    "18:48",
			},
			Messages: DefaultFieldConfig{
				WorkStart:  "Início de jornada",
				LunchStart: "Saída para almoço",
				LunchEnd:   "Retorno do almoço",
				WorkEnd:    "Final de jornada",
			},
		},
		Gard: GardField{
			Messages: GardFieldMessages{
				Default: "Guardia",
			},
			Hours: GardFieldHours{
				Begin:    []GardFieldHoursRange{{Start: "18:48", End: "24:00"}},
				Finish:   []GardFieldHoursRange{{Start: "05:00", End: "09:00"}},
				Holiday:  []GardFieldHoursRange{{Start: "05:00", End: "18:00"}},
				Saturday: []GardFieldHoursRange{{Start: "05:00", End: "18:00"}},
				Sunday:   []GardFieldHoursRange{{Start: "07:00", End: "23:00"}},
				Weekday:  []GardFieldHoursRange{{Start: "00:00", End: "02:00"}, {Start: "05:00", End: "09:00"}, {Start: "18:48", End: "24:00"}},
			},
		},
		Holidays: []string{
			"01-01-20",
			"25-01-20",
			"24-02-20",
			"25-02-20",
			"10-04-20",
			"21-04-20",
			"01-05-20",
			"11-06-20",
			"09-07-20",
			"07-09-20",
			"12-10-20",
			"02-11-20",
			"15-11-20",
			"20-11-20",
			"25-12-20",
		},
	}
}
