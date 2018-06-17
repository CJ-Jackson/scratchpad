package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// Step 1: The Create the Error Collector and Validation Error type

type ErrorCollector []error

func (e ErrorCollector) Error() string {
	str := []string{}
	for _, err := range e {
		str = append(str, fmt.Sprint(err))
	}
	return strings.Join(str, "\n")
}

func (e ErrorCollector) FilterErrors() error {
	errorCollection := ErrorCollector{}
	for _, err := range e {
		if nil != err {
			errorCollection = append(errorCollection, err)
		}
	}

	if len(errorCollection) == 0 {
		return nil
	}

	return errorCollection
}

type ValidationError string

func (e ValidationError) Error() string {
	return fmt.Sprint(string(e))
}

// Step 2: Create a Validation Rule Type and Function

type StringValidationRule func(value *string, hasError bool) error

func StringValidator(value string, rules ...StringValidationRule) (string, error) {
	value = strings.TrimSpace(value)
	valuePtr := &value
	hasError := false

	errorCollection := ErrorCollector{}
	for _, rule := range rules {
		if err := rule(valuePtr, hasError); nil != err {
			hasError = true
			errorCollection = append(errorCollection, err)
		}
	}

	return value, errorCollection.FilterErrors()
}

// Step 3:  Lets create a couple of rules. The Goal of the rule is not to satisfy the if statement.

func StringMaxRune(max int) StringValidationRule {
	return func(value *string, hasError bool) error {
		if utf8.RuneCountInString(*value) > max {
			return ValidationError(fmt.Sprintf("Value should not be more than '%d'", max))
		}

		return nil
	}
}

func StringMinRune(min int) StringValidationRule {
	return func(value *string, hasError bool) error {
		if utf8.RuneCountInString(*value) < min {
			return ValidationError(fmt.Sprintf("Value should be more than '%d'", min))
		}

		return nil
	}
}

// Step: 4 Now lets run the validator.

func main() {
	rules := []StringValidationRule{
		StringMinRune(2),
		StringMaxRune(4),
	}

	values := []string{"a", "cat", "dog", "horse"}
	for _, value := range values {
		value, err := StringValidator(value, rules...)
		fmt.Println("Value:", value)
		fmt.Println("Valid:", nil == err)
		fmt.Println("Error:", err)
	}
}
