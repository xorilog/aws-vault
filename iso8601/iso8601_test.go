package iso8601

import (
	"encoding/json"
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	input, _ := time.Parse(time.RFC3339, "2009-02-04T21:00:57-08:00")
	want := "2009-02-05T05:00:57Z"
	result := Format(input)
	if result != want {
		t.Errorf("expected %s for %q got %s", want, input, result)
	}
}

func TestParseRoundtrip(t *testing.T) {
	input := "2009-02-05T05:00:57Z"
	s, _ := Parse(input)
	result := s.String()
	if result != input {
		t.Errorf("expected %s got %s", input, result)
	}
}

// the Java SDK cannot parse this string, see https://github.com/99designs/aws-vault/issues/655
func TestParseForJavaSDK(t *testing.T) {
	input := "2020-09-10T18:16:52+02:00"
	_, err := Parse(input)
	if err == nil {
		t.Error("expected an error")
	}
}

func TestJsonRoundtrip(t *testing.T) {
	now := New(time.Now())

	data, err := json.Marshal(now)
	if err != nil {
		t.Fatal(err)
	}

	_, err = time.Parse(jsonFormat, string(data))
	if err != nil {
		t.Fatal(err)
	}

	var now2 Time
	err = json.Unmarshal(data, &now2)
	if err != nil {
		t.Fatal(err)
	}

	if now.String() != now2.String() {
		t.Fatalf("String format for %s does not equal expected %s", now2, now)
	}
}
