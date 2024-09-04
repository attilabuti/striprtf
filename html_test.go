package striprtf

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestExtractHtml(t *testing.T) {
	file, err := os.Open("./testdata/encoding_1251.rtf")
	if err != nil {
		t.Fatalf("Failed to open test file: %v", err)
	}
	defer file.Close()

	r, err := ExtractHtml(file)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	html, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedTexts := []string{
		"<div dir=\"ltr\">你好，世界",
		"שלום, עולם!<br clear=\"all\"><div><br>",
		"<div style=\"font-size:12.8px\">This is your third encoding",
	}

	for _, expectedText := range expectedTexts {
		if !bytes.Contains(html, []byte(expectedText)) {
			t.Errorf("Expected text not found: %s", expectedText)
		}
	}
}

func TestNonUnicodeMail(t *testing.T) {
	file, err := os.Open("./testdata/non_unicode_mail.rtf")
	if err != nil {
		t.Fatalf("Failed to open test file: %v", err)
	}
	defer file.Close()

	r, err := ExtractHtml(file)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	html, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedTexts := []string{
		`<html xmlns:v="urn:schemas-microsoft-com:vml" xmlns:o="urn:schemas-microsoft-com:office:office" xmlns:w="urn:schemas-microsoft-com:office:word" xmlns:m="http://schemas.microsoft.com/office/2004/12/omml" xmlns="http://www.w3.org/TR/REC-html40">`,
		`<w:LsdException Locked="false" Priority="9" SemiHidden="true" UnhideWhenUsed="true" QFormat="true" Name="heading 9"/>`,
		`<div class=WordSection1><p class=MsoNormal><span lang=EN-US style='font-size:10.0pt;mso-bidi-font-size:11.0pt'>Non Unicode mail body!<o:p></o:p></span>`,
	}

	for _, expectedText := range expectedTexts {
		if !bytes.Contains(html, []byte(expectedText)) {
			t.Errorf("Expected text not found: %s", expectedText)
		}
	}
}
