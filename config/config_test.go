package config

import (
  "testing"
  "github.com/stretchr/testify/assert"
)

func configurationMergeTest(t *testing.T) {
  c1 := Configuration{token: "a", channels: []string{"#test"}}
  c2 := Configuration{token: "b", channels: []string{"#testing", "#testing2"}}

  m := c1.Merge(c2)
  assert.Equal(m.token, "b", "merged config didn't allow overriding of the token")
}
