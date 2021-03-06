package executer

import (
	"net/http"
	"strings"
	"unsafe"
)

type jsonOutput struct {
	Template         string                 `json:"template"`
	Type             string                 `json:"type"`
	Matched          string                 `json:"matched"`
	MatcherName      string                 `json:"matcher_name,omitempty"`
	ExtractedResults []string               `json:"extracted_results,omitempty"`
	Name             string                 `json:"name"`
	Severity         string                 `json:"severity"`
	Author           string                 `json:"author"`
	Description      string                 `json:"description"`
	Request          string                 `json:"request,omitempty"`
	Response         string                 `json:"response,omitempty"`
	Meta             map[string]interface{} `json:"meta,omitempty"`
}

// unsafeToString converts byte slice to string with zero allocations
func unsafeToString(bs []byte) string {
	return *(*string)(unsafe.Pointer(&bs))
}

// headersToString converts http headers to string
func headersToString(headers http.Header) string {
	builder := &strings.Builder{}

	for header, values := range headers {
		builder.WriteString(header)
		builder.WriteString(": ")

		for i, value := range values {
			builder.WriteString(value)

			if i != len(values)-1 {
				builder.WriteRune('\n')
				builder.WriteString(header)
				builder.WriteString(": ")
			}
		}

		builder.WriteRune('\n')
	}

	return builder.String()
}
