package argument

import (
	"fmt"
	"net/url"
	"testing"
)

func TestTool(t *testing.T) {
	query, _ := url.ParseQuery("")
	fmt.Println(query)
}
