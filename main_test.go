package main

import (
	"testing"
)

func TestIsASCII(t *testing.T) {
	// Test cases where the input is ASCII
	asciiInputs := []string{
		"Hello, World!",
		"12345",
		"Testing ASCII",
		"Special Characters: !@#$%^&*()_+",
	}

	for _, input := range asciiInputs {
		if !isASCII(input) {
			t.Errorf("Expected isASCII(%q) to be true, but it returned false", input)
		}
	}

	// Test cases where the input is not ASCII
	nonASCIIInputs := []string{
		"こんにちは",
		"你好",
		"Résumé",
	}

	for _, input := range nonASCIIInputs {
		if isASCII(input) {
			t.Errorf("Expected isASCII(%q) to be false, but it returned true", input)
		}
	}
}
func TestReturnAsciiCodeInt(t *testing.T) {
	input := "Hello"
	expected := []int{40, 69, 76, 76, 79}
	result := returnAsciiCodeInt(input)

	if len(result) != len(expected) {
		t.Errorf("Expected len(result) to be %d, but it was %d", len(expected), len(result))
	}

	for i := range result {
		if result[i] != expected[i] {
			t.Errorf("Expected result[%d] to be %d, but it was %d", i, expected[i], result[i])
		}
	}
}

func TestLenOfAsciiArtWithSpace(t *testing.T) {
	// Test case with ASCII art lines of varying lengths
	asciiArt := []string{
		"1234567890",
		" / _ \\  ",
		"| (_) |  ",
		" \\___/   ",
		"   |     ",
		"  / \\   ",
		" / _ \\  ",
	}
	expected := 10 // Maximum horizontal length

	result := lenOfAsciiArtWithSpace(asciiArt)

	if result != expected {
		t.Errorf("Expected lenOfAsciiArtWithSpace() to return %d (maximum horizontal length), but it returned %d", expected, result)
	}

	// Test case with ASCII art lines of equal length
	equalLengthArt := []string{
		"12345",
		"ABCDE",
		"!@#$%",
	}
	expectedEqual := 5 // Maximum horizontal length

	resultEqual := lenOfAsciiArtWithSpace(equalLengthArt)

	if resultEqual != expectedEqual {
		t.Errorf("Expected lenOfAsciiArtWithSpace() to return %d (maximum horizontal length), but it returned %d", expectedEqual, resultEqual)
	}
}
