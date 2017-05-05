package genericresource

import (
	"testing"

	"github.com/docker/swarmkit/api"
	"github.com/stretchr/testify/assert"
)

func TestParseDiscrete(t *testing.T) {
	res := &api.Resources{}

	err := Parse("apple=3", res)
	assert.NoError(t, err)
	assert.Equal(t, len(res.Generic), 1)

	apples := GetResource("apple", res.Generic)
	assert.Equal(t, len(apples), 1)
	assert.Equal(t, apples[0].GetDiscrete().Value, int64(3))
}

func TestParseStr(t *testing.T) {
	res := &api.Resources{}

	err := Parse("orange={red,green,blue}", res)
	assert.NoError(t, err)
	assert.Equal(t, len(res.Generic), 3)

	oranges := GetResource("orange", res.Generic)
	assert.Equal(t, len(oranges), 3)
	for _, k := range []string{"red", "green", "blue"} {
		assert.True(t, HasResource(NewString("orange", k), oranges))
	}
}

func TestParseDiscreteAndStr(t *testing.T) {
	res := &api.Resources{}

	err := Parse("orange={red,green,blue};apple=3", res)
	assert.NoError(t, err)
	assert.Equal(t, len(res.Generic), 4)

	oranges := GetResource("orange", res.Generic)
	assert.Equal(t, len(oranges), 3)
	for _, k := range []string{"red", "green", "blue"} {
		assert.True(t, HasResource(NewString("orange", k), oranges))
	}

	apples := GetResource("apple", res.Generic)
	assert.Equal(t, len(apples), 1)
	assert.Equal(t, apples[0].GetDiscrete().Value, int64(3))
}
