package content_test

import (
	"testing"

	"github.com/mkrivickiy/crm-management/core/rest/content"
)

func TestValuesEncode(t *testing.T) {
	content := content.Values{}
	encodedString := content.Encode()
	if encodedString == "" {
		t.Error("Empty result of the Encode function")
	}

	if encodedString != "friend=Jess&friend=Sarah&friend=Zoe&name=Ava" {
		t.Error("Invalid result of the Encode function")
	}

	if encodedString == "friend=Jess&friend=Sarah&friend=Zoe&name=Ava" {
		t.Log("Successful result of the Encode function")
	}

}
