package gomime

import "testing"

func TestHeaders(t *testing.T) {
	if HeaderContentType != "Content-Type" {
		t.Error("Invalid HeaderContentType constant value.")
	}

	if HeaderUserAgent != "User-Agent" {
		t.Error("Invalid HeaderUserAgent constant value.")
	}

	if ContentTypeJSON != "application/json" {
		t.Error("Invalid ContentTypeJSON constant value.")
	}

	if ContentTypeXML != "application/xml" {
		t.Error("Invlid ContentTypeXML constant value.")
	}

	if ContentTypeOctetStream != "application/octet-stream" {
		t.Error("Invalid ContentTypeOctetStream constant value.")
	}
}
