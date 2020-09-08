package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterPlayer(t *testing.T) {
	cases := []struct {
		name     string
		filters  []string
		player   string
		expected bool
	}{
		{
			name: "Player in filters",
			filters: []string{
				"filterPlayer",
			},
			player:   "filterPlayer.player1",
			expected: true,
		},
		{
			name: "Player not in filters",
			filters: []string{
				"filterPlayer",
			},
			player:   "nofilterPlayer.player1",
			expected: false,
		},
		{
			name:     "No filters",
			filters:  []string{},
			player:   "filterPlayer.player1",
			expected: false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := filterPlayer(tc.player, tc.filters)

			assert.Equal(t, got, tc.expected)
		})
	}
}

func TestFilterStrings(t *testing.T) {
	cases := []struct {
		name     string
		players  []string
		filters  []string
		expected []string
	}{
		{
			name: "Filter one",
			players: []string{
				"filter1.player1",
				"filter2.player2",
			},
			filters: []string{
				"filter1",
			},
			expected: []string{
				"filter2.player2",
			},
		},
		{
			name: "Filter multiple",
			players: []string{
				"filter1.player1",
				"filter2.player2",
			},
			filters: []string{
				"filter1",
				"filter2",
			},
			expected: []string{},
		},
		{
			name:     "All empty",
			players:  []string{},
			filters:  []string{},
			expected: []string{},
		},
		{
			name: "Filters empty",
			players: []string{
				"filter1.player1",
				"filter2.player2",
			},
			filters: []string{},
			expected: []string{
				"filter1.player1",
				"filter2.player2",
			},
		},
		{
			name:    "Players empty",
			players: []string{},
			filters: []string{
				"filter1",
				"filter2",
			},
			expected: []string{},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := FilterStrings(tc.players, tc.filters)

			assert.Equal(t, got, tc.expected)
		})
	}
}
