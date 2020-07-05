// File: language.go
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
	Contains the definition of the language codes and plain names.

	This package provides the definition of language ISO names and codes.
*/
package language

import(
	"os"
	"os/exec"
	"strings"
)

// Based from ISO language codes.
const(
	AF		= "af"		// Afrikaans
	AF_ZA	= "af-ZA"	// Afrikaans (South Africa)
	AR		= "ar"		// Arabic
	AR_AE	= "ar-AE"	// Arabic (United Arab Emirates)
	AR_BH	= "ar-BH"	// Arabic (Bahrain)
	AR_DZ	= "ar-DZ"	// Arabic (Algeria)
	AR_EG	= "ar-EG"	// Arabic (Egypt)
	AR_IQ	= "ar-IQ"	// Arabic (Iraq)
	AR_JO	= "ar-JO"	// Arabic (Jordan)
	AR_KW	= "ar-KW"	// Arabic (Kuwait)
	AR_LB	= "ar-LB"	// Arabic (Lebanon)
	AR_LY	= "ar-LY"	// Arabic (Libya)
	AR_MA	= "ar-MA"	// Arabic (Morocco)
	AR_OM	= "ar-OM"	// Arabic (Oman)
	AR_QA	= "ar-QA"	// Arabic (Qatar)
	AR_SA	= "ar-SA"	// Arabic (Saudi Arabia)
	AR_SY	= "ar-SY"	// Arabic (Syria)
	AR_TN	= "ar-TN"	// Arabic (Tunisia)
	AR_YE	= "ar-YE"	// Arabic (Yemen)
	AZ		= "az"		// Azeri (Latin)
	AZ_AZ	= "az-AZ"	// Azeri (Azerbaijan)
	BE		= "be"		// Belarusian
	BE_BY	= "be-BY"	// Belarusian (Belarus)
	BG		= "bg"		// Bulgarian
	BG_BG	= "bg-BG"	// Bulgarian (Bulgaria)
	BS_BA	= "bs-BA"	// Bosnian (Bosnia and Herzegovina)
	CA		= "ca"		// Catalan
	CA_ES	= "ca-ES"	// Catalan (Spain)
	CS		= "cs"		// Czech
	CS_CZ	= "cs-CZ"	// Czech (Czech Republic)
	CY		= "cy"		// Welsh
	CY_GB	= "cy-GB"	// Welsh (United Kingdom)
	DA		= "da"		// Danish
	DA_DK	= "da-DK"	// Danish (Denmark)
	DE		= "de"		// German
	DE_AT	= "de-AT"	// German (Austria)
	DE_CH	= "de-CH"	// German (Switzerland)
	DE_DE	= "de-DE"	// German (Germany)
	DE_LI	= "de-LI"	// German (Liechtenstein)
	DE_LU	= "de-LU"	// German (Luxembourg)
	DV		= "dv"		// Divehi
	DV_MV	= "dv-MV"	// Divehi (Maldives)
	EL		= "el"		// Greek
	EL_GR	= "el-GR"	// Greek (Greece)
	EN		= "en"		// English
	EN_AU	= "en-AU"	// English (Australia)
	EN_BZ	= "en-BZ"	// English (Belize)
	EN_CA	= "en-CA"	// English (Canada)
	EN_CB	= "en-DB"	// English (Caribbean)
	EN_GB	= "en-GB"	// English (United Kingdom)
	EN_IE	= "en-IE"	// English (Ireland)
	EN_JM	= "en-JM"	// English (Jamaica)
	EN_NZ	= "en-NZ"	// English (New Zealand)
	EN_PH	= "en-PH"	// English (Philippines)
	EN_TT	= "en-TT"	// English (Trinidad and Tobago)
	EN_US	= "en-US"	// English (United States)
	EN_ZA	= "en-ZA"	// English (South Africa)
	EN_ZW	= "en-ZW"	// English (Zimbabwe)
	EO		= "eo"		// Esperanto
	ES		= "es"		// Spanish
	ES_AR	= "es-AR"	// Spanish (Argentina)
	ES_BO	= "es-BO"	// Spanish (Bolivia)
	ES_CL	= "es-CL"	// Spanish (Chile)
	ES_CO	= "es-CO"	// Spanish (Colombia)
	ES_CR	= "es-CR"	// Spanish (Costa Rica)
	ES_DO	= "es-DO"	// Spanish (Dominican Republic)
	ES_EC	= "es-EC"	// Spanish (Ecuador)
	ES_ES	= "es-ES"	// Spanish (Castilian)
	ES_GT	= "es-GT"	// Spanish (Guatemala)
	ES_HN	= "es-HN"	// Spanish (Honduras)
	ES_MX	= "es-MX"	// Spanish (Mexico)
	ES_NI	= "es-NI"	// Spanish (Nicaragua)
	ES_PA	= "es-PA"	// Spanish (Panama)
	ES_PE	= "es-PE"	// Spanish (Peru)
	ES_PR	= "es-PR"	// Spanish (Puerto Rico)
	ES_PY	= "es-PY"	// Spanish (Paraguay)
	ES_SV	= "es-SV"	// Spanish (El Salvador)
	ES_UY	= "es-UY"	// Spanish (Uruguay)
	ES_VE	= "es-VE"	// Spanish (Venezuela)
	ET		= "et"		// Estonian
	ET_EE	= "et-EE"	// Estonian (Estonia)
	EU		= "eu"		// Basque
	EU_ES	= "es-ES"	// Basque (Spain)
	FA		= "fa"		// Farsi
	FA_IR	= "fa-IR"	// Farsi (Iran)
	FI		= "fi"		// Finnish
	FI_FI	= "fi-FI"	// Finnish (Finland)
	FO		= "fo"		// Faroese
	FO_FO	= "fo-FO"	// Faroese (Faroe Islands)
	FR		= "fr"		// French
	FR_BE	= "fr-BE"	// French (Belgium)
	FR_CA	= "fr-CA"	// French (Canada)
	FR_CH	= "fr-CH"	// French (Switzerland)
	FR_FR	= "fr-FR"	// French (Franch)
	FR_LU	= "fr-LU"	// French (Luxembourg)
	FR_MC	= "fr-MC"	// French (Principality of Monaco)
	GL		= "gl"		// Galician
	GL_ES	= "gl-ES"	// Galician (Spain)
	GU		= "gu"		// Gujarati
	GU_IN	= "gu-IN"	// Gujarati (India)
	HE		= "he"		// Hebrew
	HE_IL	= "he-IL"	// Hebrew (Israel)
	HI		= "hi"		// Hindi
	HI_IN	= "hi-IN"	// Hindi (India)
	HR		= "hr"		// Croatian
	HR_BA	= "hr-BA"	// Croatian (Bosnia and Herzegovina)
	HR_HR	= "hr-HR"	// Croatian (Croatia)
	HU		= "hu"		// Hungarian
	HU_HU	= "hu-HU"	// Hungarian (Hungary)
	HY		= "hy"		// Armenian
	HY_AM	= "hy-AM"	// Armenian (Armenia)
	ID		= "id"		// Indonesian
	ID_ID	= "id-ID"	// Indonesian (Indonesia)
	IS		= "is"		// Icelandic
	IS_IS	= "is-IS"	// Icelandic
	IT		= "it"		// Italian
	IT_CH	= "it-CH"	// Italian (Switzerland)
	IT_IT	= "it-IT"	// Italian (Italy)
	JA		= "ja"		// Japanese
	JA_JP	= "ja-JP"	// Japanese (Japan)
	KA		= "ka"		// Georgian
	KA_GE	= "ka-GE"	// Georgian (Georgia)
	KK		= "kk"		// Kazakh
	KK_KZ	= "kk-KZ"	// Kazakh (Kazakhstan)
	KN		= "kn"		// Kannada
	KN_IN	= "kn-IN"	// Kannada (India)
	KO		= "ko"		// Korean
	KO_KR	= "ko-KR"	// Korean (Korea)
	KOK		= "kok"		// Konkani
	KOK_IN	= "kok-IN"	// Konkani (India)
	KY		= "ky"		// Kyrgyz
	KY_KG	= "ky-KG"	// Kyrgyz (Kyrgyzstan)
	LT		= "lt"		// Lithuanian
	LT_LT	= "lt-LT"	// Lithuanian (Lithuania)
	LV		= "lv"		// Latvian
	LV_LV	= "lv-LV"	// Latvian (Latvia)
	MI		= "mi"		// Maori
	MI_NZ	= "mi-NZ"	// Maori (New Zealand)
	MK		= "mk"		// FYRO Macedonian
	MK_MK	= "mk-MK"	// FYRO Macedonian (Former Yugoslav Republic of Macedonia)
	MN		= "mn"		// Mongolian
	MN_MN	= "mn-MN"	// Mongolian (Mongolia)
	MR		= "mr"		// Marathi
	MR_IN	= "mr-IN"	// Marathi (India)
	MS		= "ms"		// Malay
	MS_BN	= "ms-BN"	// Malay (Brunei Darussalam)
	MS_MY	= "ms-MY"	// Malay (Malaysia)
	MT		= "mt"		// Maltese
	MT_MT	= "mt-MT"	// Maltese (Malta)
	NB		= "nb"		// Norwegian (Bokm?l)
	NB_NO	= "nb-NO"	// Norwegian (Bokm?l)(Norway)
	NL		= "nl"		// Dutch
	NL_BE	= "nl-BE"	// Dutch (Belgium)
	NL_NL	= "nl-NL"	// Dutch (Netherlands)
	NN_NO	= "nn-NO"	// Norwegian (Nynorsk)(Norway)
	NS		= "ns"		// Northern Sotho
	NS_ZA	= "ns-ZA"	// Northern Sotho (South Africa)
	PA		= "pa"		// Punjabi
	PA_IN	= "pa-IN"	// Punjabi (India)
	PL		= "pl"		// Polish
	PL_PL	= "pl-PL"	// Polish (Poland)
	PS		= "ps"		// Pashto
	PS_AR	= "ps-AR"	// Pashto (Afghanistan)
	PT		= "pt"		// Portuguese
	PT_BR	= "pt-BR"	// Portuguese (Brazil)
	PT_PT	= "pt-PT"	// Portuguese (Portugal)
	QU		= "qu"		// Quechua
	QU_BO	= "qu-BO"	// Quechua (Bolivia)
	QU_EC	= "qu-EC"	// Quechua (Ecuador)
	QU_PE	= "qu-PE"	// Quechua (Peru)
	RO		= "ro"		// Romanian
	RO_RO	= "ro-RO"	// Romanian (Romania)
	RU		= "ru"		// Russian
	RU_RU	= "ru-RU"	// Russian (Russia)
	SA		= "sa"		// Sanskrit
	SA_IN	= "sa-IN"	// Sanskrit (India)
	SE		= "se"		// Sami (Northern)
	SE_FI	= "se-FI"	// Sami (Northern/Skolt/Inari)(Finland)
	SE_NO	= "se-NO"	// Sami (Northern/Lule/Southern)(Norway)
	SE_SE	= "se-SE"	// Sami (Northern/Lule/Southern)(Sweden)
	SK		= "sk"		// Slovak
	SK_SK	= "sk-SK"	// Slovak (Slovakia)
	SL		= "sl"		// Slovenian
	SL_SI	= "sl-SI"	// Slovenian (Slovenia)
	SQ		= "sq"		// Albanian
	SQ_AL	= "sq-AL"	// Albanian (Albania)
	SR_BA	= "sr-BA"	// Serbian (Latin/Cyrillic)(Bosnia and Herzegovina)
	SR_SP	= "sr-SP"	// Serbian (Latin/Cyrillic)(Serbia and Montenegro)
	SV		= "sv"		// Swedish
	SV_FI	= "sv-FI"	// Swedish (Finland)
	SV_SE	= "sv-SE"	// Swedish (Sweden)
	SW		= "sw"		// Swahili
	SW_KE	= "sw-KE"	// Swahili (Kenya)
	SYR		= "syr"		// Syriac
	SYR_SY	= "syr-SY"	// Syriac (Syria)
	TA		= "ta"		// Tamil
	TA_IN	= "ta-IN"	// Tamil (India)
	TE		= "te"		// Telugu
	TE_IN	= "te-IN"	// Telugu (India)
	TH		= "th"		// Thai
	TH_TH	= "th-TH"	// Thai (Thailand)
	TL		= "tl"		// Tagalog
	TL_PH	= "tl-PH"	// Tagalog (Republic of Philippines)
	TN		= "tn"		// Tswana
	TN_ZA	= "tn-ZA"	// Tswana (South Africa)
	TR		= "tr"		// Turkish
	TR_TR	= "tr-TR"	// Turkish (Turkey)
	TT		= "tt"		// Tatar
	TT_RU	= "tt-RU"	// Tatar (Russia)
	TS		= "ts"		// Tsonga
	UK		= "uk"		// Ukrainian
	UK_UA	= "uk-UA"	// Ukrainian (Ukraine)
	UR		= "ur"		// Urdu
	UR_PK	= "ur-PK"	// Urdu (Islamic Republic of Pakistan)
	UZ		= "uz"		// Uzbek
	UZ_UZ	= "uz-UZ"	// Uzbek (Latin/Cyrillic)(Uzbekistan)
	VI		= "vi"		// Vietnamese
	VI_VN	= "vi-VN"	// Vietnamese (Viet Nam)
	XH		= "xh"		// Xhosa
	XH_ZA	= "xh-ZA"	// Xhosa (South Africa)
	ZH		= "zh"		// Chinese
	ZH_CN	= "zh-CN"	// Chinese (China)
	ZH_HK	= "zh-HK"	// Chinese (Hong Kong)
	ZH_MO	= "zh-MO"	// Chinese (Macau)
	ZH_SG	= "zh-SG"	// Chinese (Singapore)
	ZH_TW	= "zh-TW"	// Chinese (Taiwan)
	ZU		= "zu"		// Zulu
	ZU_ZA	= "zu-ZA"	// Zulu (South Africa)
)

// Gets the current language of the system by:
//	1. Checks the LANG environment variable on non-Windows systems.
//	2. Checks the culture via PowerShell on Windows.
func GetSystemLanguage() (string, error) {
	if environmentLangVar, found := os.LookupEnv("LANG"); found {
		return strings.Split(environmentLangVar, ".")[0], nil
	}

	result, shellError := exec.Command("powershell", "Get-Culture | select -exp Name").Output()
	if shellError != nil {
		return "", shellError
	}

	return strings.Trim(string(result), "\r\n"), nil
}
