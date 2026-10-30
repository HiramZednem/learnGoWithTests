package iteration

import "testing"

func TestIsWordInside(t *testing.T) {
	sentence := "What a wonderfull day"

	t.Run("word is inside sentence", func(t *testing.T) {
		expected := "yes";
		got := IsWordInside(sentence, "wonderfull")

		if expected != got {
			t.Errorf("Expected %q but got %q", expected, got);
		}
	})

	t.Run("word is not in the sentence", func(t *testing.T) {
		expected := "no";
		got := IsWordInside(sentence, "good")

		if expected != got {
			t.Errorf("Expected %q but got %q", expected, got);
		}
	})

}

func TestCountChars(t *testing.T) {
	word := "Cheese"
	char := "e";
	expected := 3
	got := CountChars(word, char)

	if expected != got {
		t.Errorf("Word: %q expected to have %d %q but got %d", word, expected, char, got);
	}
}

func TestCutPrefix(t *testing.T) {
	s := "fdrr_file_11-30-2026"

	t.Run("Cut Segment Succesfully", func(t *testing.T){
		suffix := "fdrr_file_"
		expected := "11-30-2026"
		got := CutPrefix(s, suffix)

		if expected != got {
			t.Errorf("Expected %q but got %q", expected, got);
		}
	})

	t.Run("Cut Segment Doesn't Exist", func(t *testing.T){
		suffix := "als_file_"
		expected := ""
		got := CutPrefix(s, suffix)

		if expected != got {
			t.Errorf("Expected %q but got %q", expected, got);
		}
	})
}

// split and sep
func TestSeparateString(t *testing.T) {
	sentence := "Thrusting out all your enemies before you, as the LORD said."

	t.Run("Cut by whitespaces", func(t *testing.T){
		result := SeparateString(sentence, "");
		pieces := len(result)
		expectedPieces := 11

		if expectedPieces != pieces {
			t.Errorf("Expected to have %d pieces, but got %d", expectedPieces, pieces)
		}
	})

	t.Run("Cut by comma", func(t *testing.T){
		result := SeparateString(sentence, ",");
		pieces := len(result)
		expectedPieces := 2

		if expectedPieces != pieces {
			t.Errorf("Expected to have %d pieces, but got %d", expectedPieces, pieces)
		}
	})
}