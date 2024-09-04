package striprtf

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestExtractText(t *testing.T) {
	file, err := os.Open("./testdata/rtf.rtf")
	if err != nil {
		t.Fatalf("Failed to open test file: %v", err)
	}
	defer file.Close()

	r, err := ExtractText(file)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	text, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedText := []byte(`Bold REPLACED Underline Size 14

Times New Roman 11. Adipiscing elit. Integer quis eros at tortor underline1. Quisque viverra tellus id mauris blue1 luctus. Fusce in interdum ipsum. Cum sociis natoque penatibus et italic1 dis parturient montes, nascetur ridiculus mus. Special characters: tést1 añu {\test2}.

Arial 12. Ac leo justo, vitae rutrum elit. Nam purus odio, bold1. Etiam vitae red1. Aenean molestie, quis blue2 leo placerat in. REPLACED2 malesuada eleifend nunc vitae cursus. Praesent dapibus aliquet sem, ac pharetra ipsum tempus id.ést.
`)

	if !bytes.Equal(text, expectedText) {
		t.Errorf("Expected text not found: %s", expectedText)
	}
}

func TestEncoding1251(t *testing.T) {
	file, err := os.Open("./testdata/encoding_1251.rtf")
	if err != nil {
		t.Fatalf("Failed to open test file: %v", err)
	}
	defer file.Close()

	r, err := ExtractText(file)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	text, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedTexts := []string{
		"你好，世界",
		"שלום, עולם!\n",
		"\nThis is your third encoding\n",
	}

	for _, expectedText := range expectedTexts {
		if !bytes.Contains(text, []byte(expectedText)) {
			t.Errorf("Expected text not found: %s", expectedText)
		}
	}
}

func TestEncoding932(t *testing.T) {
	file, err := os.Open("./testdata/encoding_932.rtf")
	if err != nil {
		t.Fatalf("Failed to open test file: %v", err)
	}
	defer file.Close()

	r, err := ExtractText(file)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	text, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedText := []byte("こんにちは． 世界？\n")
	if !bytes.Equal(text, expectedText) {
		t.Errorf("Expected text not found: %s", expectedText)
	}
}

func TestNonUnicode(t *testing.T) {
	file, err := os.Open("./testdata/non_unicode_mail.rtf")
	if err != nil {
		t.Fatalf("Failed to open test file: %v", err)
	}
	defer file.Close()

	r, err := ExtractText(file)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	text, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedText := []byte("Non Unicode mail body!\n")
	if !bytes.Equal(text, expectedText) {
		t.Errorf("Expected text not found: %s", expectedText)
	}
}

func TestLarge(t *testing.T) {
	file, err := os.Open("./testdata/large.rtf")
	if err != nil {
		t.Fatalf("Failed to open test file: %v", err)
	}
	defer file.Close()

	r, err := ExtractText(file)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	text, err := io.ReadAll(r)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedTexts := []string{
		"Lorem ipsum (Heading 1)",
		"Dolor sit amet (Heading 2)",
		"Lorem ipsum dolor sit amet, http://www.google.com elit. Integer quis eros at tortor pharetra laoreet.",
		"Quisque viverra tellus (Heading 3)",
		"Cras pharetra, velit vel malesuada lobortis, nibh neque luctus lorem, pretium rutrum diam ligula quis risus. Sed pretium orci nec metus rutrum porttitor. Fusce at libero quis orci consequat pulvinar.",
	}

	for _, expectedText := range expectedTexts {
		if !bytes.Contains(text, []byte(expectedText)) {
			t.Errorf("Expected text not found: %s", expectedText)
		}
	}
}
