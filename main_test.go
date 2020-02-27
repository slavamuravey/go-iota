// This file contains tests that demonstrate how iota works.
package iota

import (
	"testing"
)

type iotaValue uint

type testSet struct {
	actual   iotaValue
	expected iotaValue
}

func TestIotaEverywhere(t *testing.T) {
	const (
		iota1 iotaValue = iota
		iota2           = iota
		iota3           = iota
	)

	testSets := []testSet{
		{iota1, 0},
		{iota2, 1},
		{iota3, 2},
	}

	runTestSets(t, testSets)
}

func TestIotaOnce(t *testing.T) {
	const (
		iota1 iotaValue = iota
		iota2
		iota3
	)

	testSets := []testSet{
		{iota1, 0},
		{iota2, 1},
		{iota3, 2},
	}

	runTestSets(t, testSets)
}

func TestIotaSometime(t *testing.T) {
	const (
		iota1 iotaValue = iota
		iota2
		iota3 iotaValue = iota
	)

	testSets := []testSet{
		{iota1, 0},
		{iota2, 1},
		{iota3, 2},
	}

	runTestSets(t, testSets)
}

func TestIotaModified(t *testing.T) {
	const (
		iota1 iotaValue = 777 + iota
		iota2
		iota3
	)

	testSets := []testSet{
		{iota1, 777},
		{iota2, 778},
		{iota3, 779},
	}

	runTestSets(t, testSets)
}

func TestIotaGap(t *testing.T) {
	const (
		iota1 iotaValue = iota
		_
		iota3
	)

	testSets := []testSet{
		{iota1, 0},
		{iota3, 2},
	}

	runTestSets(t, testSets)
}

func TestIotaConcreteValue(t *testing.T) {
	const (
		iota1           = 555
		iota2 iotaValue = iota
		iota3           = 444
		iota4 iotaValue = iota
		iota5
	)

	testSets := []testSet{
		{iota1, 555},
		{iota2, 1},
		{iota3, 444},
		{iota4, 3},
		{iota5, 4},
	}

	runTestSets(t, testSets)
}

func runTestSets(t *testing.T, testSets []testSet) {
	for _, ts := range testSets {
		if ts.actual == ts.expected {
			continue
		}
		errorf(t, ts.actual, ts.expected)
	}
}

func errorf(t *testing.T, actual, expected interface{}) {
	t.Errorf("got %d; want %d", actual, expected)
}
