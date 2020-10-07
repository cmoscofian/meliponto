package context

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/cmoscofian/meliponto/src/cli/util/constant"
	"github.com/cmoscofian/meliponto/src/shared/domain/entities"
)

var dirname string = path.Join("/", "usr", "local", "etc")
var filename string = "meliponto.json"

// ConfigPath is the path to which the config file
// will be writted to. It is defined at compile
// time.
var ConfigPath string = path.Join(dirname, filename)

// New returns a pointer to a valid configuration
// based upon a valid config file, if unable or corrupited
// it will generate a new config file based on a default and
// exit with status 1.
func New() *entities.Context {
	content, err := ioutil.ReadFile(ConfigPath)
	if err != nil {
		fmt.Println(constant.CorruptedConfigFileError)
		GenerateConfig()
		os.Exit(1)
	}

	context := new(entities.Context)

	err = json.Unmarshal(content, context)
	if err != nil {
		fmt.Println(constant.ParseConfigError)
		GenerateConfig()
		os.Exit(1)
	}

	return context
}

// GenerateConfig creates a brand new config file from scratch
// based on the default values set on createDefaultConfig.
// [Warning]: It will override any existing config files.
func GenerateConfig() {
	c := createDefaultConfig()

	file, err := os.Create(ConfigPath)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "\t")
	if err := encoder.Encode(c); err != nil {
		log.Fatalln(err)
	}

	fmt.Fprintln(os.Stderr, constant.NewConfigFileGeneratedError)
}

func createDefaultConfig() *entities.Context {
	return &entities.Context{
		CompanyID: "a382748",
		Default: &entities.DefaultField{
			Hours: &entities.DefaultFieldConfig{
				WorkStart:  "09:00",
				LunchStart: "12:00",
				LunchEnd:   "13:00",
				WorkEnd:    "18:48",
			},
			Messages: &entities.DefaultFieldConfig{
				WorkStart:  "Início de jornada",
				LunchStart: "Saída para almoço",
				LunchEnd:   "Retorno do almoço",
				WorkEnd:    "Final de jornada",
			},
		},
		Gard: &entities.GardField{
			Messages: &entities.GardFieldMessages{
				Default: "Guardia",
			},
			Hours: &entities.GardFieldHours{
				Begin:    []*entities.GardFieldHoursRange{{Start: "18:48", End: "24:00"}},
				Finish:   []*entities.GardFieldHoursRange{{Start: "05:00", End: "09:00"}},
				Holiday:  []*entities.GardFieldHoursRange{{Start: "05:00", End: "18:00"}},
				Saturday: []*entities.GardFieldHoursRange{{Start: "05:00", End: "18:00"}},
				Sunday:   []*entities.GardFieldHoursRange{{Start: "07:00", End: "23:00"}},
				Weekday:  []*entities.GardFieldHoursRange{{Start: "00:00", End: "02:00"}, {Start: "05:00", End: "09:00"}, {Start: "18:48", End: "24:00"}},
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
