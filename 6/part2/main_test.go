package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingle(t *testing.T) {
	input := []string{"abc"}
	result := groupToCommonAnswers(input)
	assert.Equal(t, len(result), 3)
}
func TestThreeDiff(t *testing.T) {
	input := []string{"a", "b", "c"}
	result := groupToCommonAnswers(input)
	assert.Equal(t, len(result), 0)
}

func TestTwoOneCommon(t *testing.T) {
	input := []string{"ab", "ac"}
	result := groupToCommonAnswers(input)
	assert.Equal(t, len(result), 1)
}

func TestFourOneAnswer(t *testing.T) {
	input := []string{"a", "a", "a", "a"}
	result := groupToCommonAnswers(input)
	assert.Equal(t, len(result), 1)
}

func TestOneSingleAnswer(t *testing.T) {
	input := []string{"b"}
	result := groupToCommonAnswers(input)
	assert.Equal(t, len(result), 1)
}
