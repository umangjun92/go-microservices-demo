package github

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRepoRequestAsJson(t *testing.T) {
	request := &CreateRepoRequest{
		Name:        "golang-intro",
		Description: "string",
		Homepage:    "string",
		Private:     true,
		HasIssues:   true,
		HasProjects: true,
		HasWiki:     true,
	}
	bytes, err := json.Marshal(request)

	assert.Nil(t, err)
	assert.NotNil(t, bytes)

	fmt.Println(string(bytes))

	var target CreateRepoRequest

	err = json.Unmarshal(bytes, &target)

	assert.NotNil(t, err)

	fmt.Println("sdd", target)

}
