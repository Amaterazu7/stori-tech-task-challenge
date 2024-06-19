package models

import (
	"github.com/google/uuid"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type TestCase[T string | float64 | Transaction] struct {
	testName      string
	parameter     T
	expectedValue float64
}

func TestProcessorResult_AddBalance(t *testing.T) {
	processorResult := ProcessorResult{
		TotalBalance: 0,
	}

	totalBalancesTestCases := []TestCase[float64]{
		{
			testName:      "Given a negative number, should be [ -7 ]",
			parameter:     -7,
			expectedValue: -7,
		},
		{
			testName:      "Give a positive number, should be [ 200 ]",
			parameter:     +207,
			expectedValue: 200,
		},
		{
			testName:      "Give a zero number, shouldn't change the value [ 200 ]",
			parameter:     0,
			expectedValue: 200,
		},
	}

	for _, testCases := range totalBalancesTestCases {
		t.Run(testCases.testName, func(t *testing.T) {
			processorResult.AddBalance(testCases.parameter)
			assert.Equal(t, testCases.expectedValue, processorResult.TotalBalance)
		})
	}
}

func TestProcessorResult_CalculateAverageDebit(t *testing.T) {
	var times float64 = 2
	processorResult := ProcessorResult{
		AverageDebitAmount: 0,
	}

	averageDebitAmountTestCases := []TestCase[float64]{
		{
			testName:      "Given a negative float number, should be the correct value [ -77.07 ]",
			parameter:     -154.14,
			expectedValue: -77.07,
		},
		{
			testName:      "Give a big negative float number, should be the correct value [ -35.27 ]",
			parameter:     -70.534323234234,
			expectedValue: -35.27,
		},
	}

	for _, testCases := range averageDebitAmountTestCases {
		t.Run(testCases.testName, func(t *testing.T) {
			processorResult.CalculateAverageDebit(testCases.parameter, times)
			assert.Equal(t, testCases.expectedValue, processorResult.AverageDebitAmount)
		})
	}
}

func TestProcessorResult_CalculateAverageCredit(t *testing.T) {
	var times float64 = 2
	processorResult := ProcessorResult{
		AverageCreditAmount: 0,
	}

	averageCreditAmountTestCases := []TestCase[float64]{
		{
			testName:      "Given a positive float number, should be the correct value [ 77.07 ]",
			parameter:     +154.14,
			expectedValue: 77.07,
		},
		{
			testName:      "Give a big positive float number, should be the correct value [ 35.27 ]",
			parameter:     +70.534323234234,
			expectedValue: 35.27,
		},
	}

	for _, testCases := range averageCreditAmountTestCases {
		t.Run(testCases.testName, func(t *testing.T) {
			processorResult.CalculateAverageCredit(testCases.parameter, times)
			assert.Equal(t, testCases.expectedValue, processorResult.AverageCreditAmount)
		})
	}
}

func TestProcessorResult_FillAvgValues(t *testing.T) {
	processorResult := ProcessorResult{}

	debitTxA := Transaction{}
	debitTxA.Build(uuid.New(), uuid.New(), -35.2, DEBIT, time.Now())
	creditTx := Transaction{}
	creditTx.Build(uuid.New(), uuid.New(), 135.2, CREDIT, time.Now())
	debitTxB := Transaction{}
	debitTxB.Build(uuid.New(), uuid.New(), -64.2, DEBIT, time.Now())

	debitAmount, debitCount, creditAmount, creditCount := 0.0, 0.0, 0.0, 0.0

	averageCreditAmountTestCases := []TestCase[Transaction]{
		{
			testName:      "Given a negative float number, the debitAmount should increase the value [ -99.4 ]",
			parameter:     debitTxA,
			expectedValue: -35.2,
		},
		{
			testName:      "Give a positive float number, the debitAmount value shouldn't change [ -35.2 ]",
			parameter:     creditTx,
			expectedValue: -35.2,
		},
		{
			testName:      "Give a negative float number, the debitAmount should increase the value [ -99.4 ]",
			parameter:     debitTxB,
			expectedValue: -99.4,
		},
	}

	for _, testCases := range averageCreditAmountTestCases {
		t.Run(testCases.testName, func(t *testing.T) {
			processorResult.FillAvgValues(&testCases.parameter, &debitAmount, &debitCount, &creditAmount, &creditCount)

			assert.Equal(t, testCases.expectedValue, debitAmount)
		})
	}

	assert.Equal(t, -99.4, debitAmount)
	assert.Equal(t, 2.0, debitCount)
	assert.Equal(t, +135.2, creditAmount)
	assert.Equal(t, 1.0, creditCount)
}

func TestProcessorResult_FillMap(t *testing.T) {
	processorResult := ProcessorResult{
		MonthMap: make(map[string]int),
	}

	debitTxA := Transaction{}
	debitTxA.Build(uuid.New(), uuid.New(), -35.2, DEBIT, time.Date(2024, time.July, 26, 12, 19, 89, 0, time.UTC))
	creditTxA := Transaction{}
	creditTxA.Build(uuid.New(), uuid.New(), 135.2, CREDIT, time.Date(2024, time.August, 26, 12, 19, 89, 0, time.UTC))
	debitTxB := Transaction{}
	debitTxB.Build(uuid.New(), uuid.New(), -64.2, DEBIT, time.Date(2024, time.July, 26, 12, 19, 89, 0, time.UTC))
	creditTxB := Transaction{}
	creditTxB.Build(uuid.New(), uuid.New(), -64.2, DEBIT, time.Date(2024, time.June, 26, 12, 19, 89, 0, time.UTC))

	averageCreditAmountTestCases := []TestCase[Transaction]{
		{
			testName:      "Given a July Month, the map value should be [ 1 ]",
			parameter:     debitTxA,
			expectedValue: 1,
		},
		{
			testName:      "Given a August Month, the map value should be [ 1 ]",
			parameter:     creditTxA,
			expectedValue: 1,
		},
		{
			testName:      "Given a July Month, the map value should be [ 2 ]",
			parameter:     debitTxB,
			expectedValue: 2,
		},
		{
			testName:      "Given a June Month, the map value should be [ 1 ]",
			parameter:     creditTxB,
			expectedValue: 1,
		},
	}

	for _, testCases := range averageCreditAmountTestCases {
		t.Run(testCases.testName, func(t *testing.T) {
			processorResult.FillMap(&testCases.parameter)

			assert.Equal(t, int(testCases.expectedValue), processorResult.MonthMap[testCases.parameter.CreatedAt.Month().String()])
		})
	}
}
