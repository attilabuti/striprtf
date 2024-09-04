package striprtf

import (
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
)

// https://latex2rtf.sourceforge.net/rtfspec_6.html
var (
	charmaps = map[string]*charmap.Charmap{
		"437": charmap.CodePage437, // United States IBM
		"708": charmap.ISO8859_6,   // Arabic (ASMO 708)
		//"709":  nil, 				// Arabic (ASMO 449+, BCON V4)
		//"710":  nil, 				// Arabic (transparent Arabic)
		//"711":  nil, 				// Arabic (Nafitha Enhanced)
		//"720":  nil, 				// Arabic (transparent ASMO)
		"819": charmap.ISO8859_1,   // Windows 3.1 (United States and Western Europe)
		"850": charmap.CodePage850, // IBM multilingual
		"852": charmap.CodePage852, // Eastern European
		"860": charmap.CodePage860, // Portuguese
		"862": charmap.CodePage862, // Hebrew
		"863": charmap.CodePage863, // French Canadian
		//"864":  nil, 				// Arabic
		"865": charmap.CodePage865, // Norwegian
		"866": charmap.CodePage866, // Soviet Union
		"874": charmap.Windows874,  // Thai
		// 932 						// Japanese
		// 936 						// Simplified Chinese
		// 949 					    // Korean
		// 950 					    // Traditional Chinese
		"1250": charmap.Windows1250, // Windows 3.1 (Eastern European)
		"1251": charmap.Windows1251, // Windows 3.1 (Cyrillic)
		"1252": charmap.Windows1252, // Western European
		"1253": charmap.Windows1253, // Greek
		"1254": charmap.Windows1254, // Turkish
		"1255": charmap.Windows1255, // Hebrew
		"1256": charmap.Windows1256, // Arabic
		"1257": charmap.Windows1257, // Baltic
		"1258": charmap.Windows1258, // Vietnamese
		//"1361": nil, 				 // Johab
	}

	encodings = map[string]encoding.Encoding{
		"932": japanese.ShiftJIS,       // Japanese
		"936": simplifiedchinese.GBK,   // Simplified Chinese
		"949": korean.EUCKR,            // Korean
		"950": traditionalchinese.Big5, // Traditional Chinese
	}
)
