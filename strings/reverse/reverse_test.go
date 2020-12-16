package reverse

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestReverse_NonEmpty(t *testing.T) {
	output := reverse("asdfghj")
	assert.Equal(t, "jhgfdsa", output)
}

func TestReverse_Empty(t *testing.T) {
	output := reverse("")
	assert.Equal(t, "", output)
}