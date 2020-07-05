// File: gocale.go
//
// The MIT Public License
//
// Copyright 2020 Jonathan Eldy Isip-Baldivicio
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software
// and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense,
// and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all copies or substantial
// portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT
// LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN
// NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
// SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

/*
	Provides simple library for effective localization of Go applications.

	This package provides simple library for basic loading, creating, and saving localization datas.
*/
package gocale

import(
	"encoding/json"
	"io/ioutil"
	"os"
)

// Represents the locale structure, and contains two locale fundamental datas:
// language name and word entries (in dictionary form).
type Locale struct {
	// Plain language name
	Language		string

	// ISO language code of the locale.
	LanguageCode	string

	// Provides the dictionary for key and value of translation entries.
	Dictionary		map[string] string
}

// Localizing strings with templates. The key should be put between "{." and "}."
// formatter to be localized. Key should be found in the dictionary of the locale,
// otherwise, an empty string will be added instead.
//
//	exampleLocale := locale.CreateLocale("example-language")
//	exampleLocale.Dictionary["ask_the_user"] = "How are you"
//
//	fmt.Println(exampleLocale.Localize("{.ask_the_user.}?"))
//
// The key "ask_the_user" will be replace with the corresponding translation
// from the dictionary of the locale exampleLocale.
func (locale Locale) Localize(template string) string {
	var result, pos = "", 0

	for pos < len(template) {
		if pos + 2 < len(template) && template[pos: pos + 2] == "{." {
			pos = pos + 2
			key := ""

			for char := template[pos]; pos < len(template); pos = pos + 1 {
				char = template[pos]
				if pos + 2 < len(template) && template[pos:pos + 2] == ".}" {
					pos = pos + 2
					break
				}

				key = key + string(char)
			}

			result = result + locale.Translate(key)
		} else {
			result = result + string(template[pos])
			pos = pos + 1
		}
	}

	return result
}

// Save locale to the specified local file name in JSON format.
func (locale Locale) SaveToFile(fileName string) error {
	jsonData, marshalError := locale.ToJSON()
	if marshalError != nil {
		return marshalError
	}

	return ioutil.WriteFile(fileName, jsonData, 0644)
}

// Converts the locale structure to JSON. The JSON format will be in minified
// version as provided by the package "encoding/json".
func (locale Locale) ToJSON() ([]byte, error) {
	return json.Marshal(locale)
}

// Returns the corresponding translation of the parameter 'global' from the locale's
// dictionary. If the parameter input was not found, an empty string will be returned.
func (locale Locale) Translate(global string) string {
	return locale.Dictionary[global]
}

// Creates and initializes a new locale structure with an empty language dictionary.
func CreateLocale(language, code string) Locale {
	return Locale{Language: language, LanguageCode: code, Dictionary: make(map[string] string)}
}

// Checks if the locale is available by language.
func IsLocaleAvailable(code string) bool {
	if _, stat := os.Stat("locales/" + code + ".json"); stat == nil {
		return true
	}

	return false
}

// Parses the input string as locale structure in JSON format.
// The input should be in format:
//
//	{
//		"Language":	<language>,
//		"LanguageCode":	<language code>,
//		"Dictionary":	{
//			<key>: <value>...
//		}
//	}
//
// Language name should be in plain form, while the language code should be the ISO code.
func LoadFromString(data string) (Locale, error) {
	var locale Locale
	if marshalError := json.Unmarshal([]byte(data), &locale); marshalError != nil {
		return locale, marshalError
	}

	return locale, nil
}

// Loads a localization file from the subfolder 'locales' of the current directory.
func LoadLocale(code string) (Locale, error) {
	var locale Locale
	contents, readError := ioutil.ReadFile("locales/" + code + ".json")

	if readError != nil {
		return locale, readError
	}

	if unmarshalError := json.Unmarshal(contents, &locale); unmarshalError != nil {
		return locale, unmarshalError
	}

	return locale, nil
}

// Loads multiple locale structures from a single file.
func LoadLocales(fileName string) ([]Locale, error) {
	var locales []Locale
	contents, readError := ioutil.ReadFile(fileName)

	if readError != nil {
		return locales, readError
	}

	if unmarshalError := json.Unmarshal(contents, &locales); unmarshalError != nil {
		return locales, unmarshalError
	}

	return locales, nil
}
