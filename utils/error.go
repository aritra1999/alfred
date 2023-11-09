package utils

import (
	"errors"
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ExtractTableName(errorMessage string) string {
	re := regexp.MustCompile(`"(\w+)_key"`)
	match := re.FindStringSubmatch(errorMessage)
	if len(match) > 1 {
		return strings.Split(match[1], "_")[0]
	}
	return ""
}

func ProcessTableName(tableName string) string {
	table := tableName[:len(tableName)-1]
	return cases.Title(language.Und).String(table)
}

func ParseError(err error) error {
	if matched, _ := regexp.MatchString("SQLSTATE 23505", err.Error()); matched {
		tableName := ExtractTableName(err.Error())
		return errors.New(ProcessTableName(tableName) + " already exists")
	}

	return err
}
