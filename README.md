
# Gocale - Golang Simple Localization Library

Provides simple library for basic loading, creating, and saving localization datas and the definition of language ISO names and codes.

## Installing Gocale

Gocale can be installed to your locale machine by typing the following into the terminal:

```cmd
go get -v github.com/coffeeneyt/gocale
```

## Features

Here are some Gocale features that are programmer-friendly:

- Creating localization structure.
- Checking if local is available.
- Loads localization data from a local file.
- Multiple data of localization can be loaded once.
- Saving and translating a locale structure via JSON.
- Localizing strings with a template.

## Quick Guide

Localization files should be in JSON format and is located at the ```locales``` folder on the same directory of the application. For instance,

```
/hello_world
    /hello_world.exe
    /locales
        en-PH.json
        en-US.json
        ...
```

The JSON format of the localization files should only have three keys; ```Language```, ```LanguageCode```, and the ```Dictionary``` (which is hash pair type). Below is an example localization file written in pure JSON format:

```json
{
    "Language":     "Tagalog",
    "LanguageCode": "tl-PH",
    "Dictionary":   {
        "greeting": "kamusta",
        "friend":   "kaibigan"
    }
}
```

## Examples

Here are some examples useful for beginners:

### Loading Localization and Localizing with Templates

```go
package main

import(
    "fmt"
    "os"

    "github.com/coffeeneyt/gocale"
    "github.com/coffeeneyt/gocale/language"
)

func main() {
    // Load a locale structure of the language tl-PH.
    locale, loadError := gocale.LoadLocale(language.TL_PH)
    if loadError != nil {
        fmt.Println(loadError)
        os.Exit(1)
    }

    // Localize the string with a template, then print!
    fmt.Println(locale.Localize("{.greeting.}, {.friend.}!"))
}
```

### Creating a New Localization Language Template

```go
package main

import(
    "fmt"
    "os"

    "github.com/coffeeneyt/gocale"
    "github.com/coffeeneyt/gocale/language"
)

func main() {
    // Create new locale structure.
    var locale gocale.Locale = gocale.CreateLocale("Spanish", language.ES)

    // Insert dictionary entries.
    locale.Dictionary["greeting"] = "hola"
    locale.Dictionary["friend"] = "amigo"

    // Print the localized string later.
    defer fmt.Println(locale.Localize("{.greeting.}, {.friend.}!"))

    // Optionally save the locale to a local file.
    if saveError := locale.SaveToFile("locales/es.json"); saveError != nil {
        fmt.Println(saveError)
        os.Exit(1)
    }
}
```

### Detect the System Language and Check If Available

```go
package main

import(
    "fmt"
    "os"

    "github.com/coffeeneyt/gocale"
    "github.com/coffeeneyt/gocale/language"
)

func main() {
    // Detect the system's defined language.
    language, detectionError := language.GetSystemLanguage()
    if detectionError != nil {
        fmt.Println(detectionError)
        os.Exit(1)
    }

    // Print the detected language.
    fmt.Print("Your system's language is '" + language + "'. ")

    // Check is the locale is available and/or installed.
    if gocale.IsLocaleAvailable(language) {
        fmt.Println("The locale is available!")
    } else {
        fmt.Println("The locale isn't available!")
    }
}
```

## License

MIT Public License

Copyright 2020 Jonathan Eldy Isip-Baldivicio

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
