package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/72636c/hyperspaced"
)

var (
	transformNs = map[string]func(str string, n int) string{
		"lower":  hyperspaced.LowerN,
		"sponge": hyperspaced.SpongeN,
		"upper":  hyperspaced.UpperN,
	}

	spacedVariables = []*spacedVariable{
		{
			name:         "n",
			description:  "number of spaces between each character",
			valueDefault: "1",
			valueType:    "int",
			parse: func(arg *string, config *spacedConfig) error {
				n := &config.n

				if arg == nil || *arg == "" {
					*n = 1
					return nil
				}

				value, err := strconv.Atoi(*arg)
				if err != nil {
					return err
				}

				*n = value

				return nil
			},
		},
		{
			name:         "case",
			description:  caseDescription(),
			valueDefault: `""`,
			valueType:    "string",
			parse: func(arg *string, config *spacedConfig) error {
				transformN := &config.transformN

				if arg == nil || *arg == "" {
					*transformN = hyperspaced.N
					return nil
				}

				value, ok := transformNs[*arg]
				if !ok {
					return fmt.Errorf("unrecognised case '%s'", *arg)
				}

				*transformN = value

				return nil
			},
		},
	}
)

type spacedConfig struct {
	n          int
	transformN func(str string, n int) string
}

func parseSpacedConfig(args []string) (*spacedConfig, error) {
	config := new(spacedConfig)

	for index, variable := range spacedVariables {
		var arg *string
		if index < len(args)-1 {
			arg = &args[index+1]
		}

		err := variable.parse(arg, config)
		if err != nil {
			return nil, err
		}
	}

	return config, nil
}

type spacedVariable struct {
	name         string
	description  string
	valueDefault string
	valueType    string
	parse        func(arg *string, config *spacedConfig) error
}

func (variable *spacedVariable) Parameter() string {
	return fmt.Sprintf(" [%s]", variable.name)
}

func (variable *spacedVariable) Usage() string {
	return fmt.Sprintf(
		"  %s %s\n    %s (default %s)\n",
		variable.name,
		variable.valueType,
		variable.description,
		variable.valueDefault,
	)
}

func caseDescription() string {
	cases := make([]string, len(transformNs))
	index := 0

	for key := range transformNs {
		cases[index] = key
		index++
	}

	return strings.Join(cases, "|")
}

func spacedUsage() string {
	var cmd, vars strings.Builder

	cmd.WriteString("spaced")

	for _, variable := range spacedVariables {
		cmd.WriteString(variable.Parameter())
		vars.WriteString(variable.Usage())
	}

	return fmt.Sprintf("%s\n%s", cmd.String(), vars.String())
}
