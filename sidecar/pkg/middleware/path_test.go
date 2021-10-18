package middleware

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidator(t *testing.T) {
	var cases = []struct {
		path      string
		expection bool
	}{
		{"/hello", true},
		{"company", true},
		{"tenant/sj3co3s4", false},
		{"company/sd45f768", true},
		{"account/acc74850", true},
		{"company/account", true},
		{"acc734340", true},
		{"account/acc234234/user", true},
		{"account/blocked", false},
		{"tenant/account/blocked", true},
		{"tenant/account/acc23849", false},
	}

	for _, tc := range cases {
		require.Equal(t, tc.expection, ValidatePath(tc.path), fmt.Sprintf("Test is failing!: %s", tc.path))
	}
}

func TestRegexQuote(t *testing.T) {
	var cases = []struct {
		path      string
		expection string
	}{
		{"/company/", "^company$"},
		{"/company/{id}", "^company/(sd|sj|acc)[a-z0-9]+$"},
		{"/company/account", "^company/account$"},
		{"/account", "^account$"},
		{"/account/{id}", "^account/(sd|sj|acc)[a-z0-9]+$"},
		{"/{id}", "^(sd|sj|acc)[a-z0-9]+$"},
		{"/account/{id}/user", "^account/(sd|sj|acc)[a-z0-9]+/user$"},
		{"/tenant/account/blocked", "^tenant/account/blocked$"},
	}

	for _, tc := range cases {
		tc.path = strings.TrimPrefix(tc.path, "/")
		tc.path = strings.TrimSuffix(tc.path, "/")
		require.Equal(t, tc.expection, generateRegexQuery(tc.path), fmt.Sprintf("Test is failing!: %s", tc.path))
	}
}
