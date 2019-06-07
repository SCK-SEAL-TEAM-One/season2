package vending_test

import (
	vm "./vendingmachine"
	"testing"
)

func TestTotalBalanceIs_1_ShouldChangeCoin_O(t *testing.T) {
	expectedResult := ", O"
	v := vm.NewVendingMachine()
	v.TotalCoins = 1

	actualResult := v.ChangeCoins()

	if actualResult != expectedResult {
		t.Errorf("%v but got %v", expectedResult, actualResult)
	}
}

func TestTotalBalanceIs_2_ShouldChangeCoin_TW(t *testing.T) {
	expectedResult := ", TW"
	v := vm.NewVendingMachine()
	v.TotalCoins = 2

	actualResult := v.ChangeCoins()

	if actualResult != expectedResult {
		t.Errorf("%v but got %v", expectedResult, actualResult)
	}
}

func TestTotalBalanceIs_5_ShouldChangeCoin_F(t *testing.T) {
	expectedResult := ", F"
	v := vm.NewVendingMachine()
	v.TotalCoins = 5

	actualResult := v.ChangeCoins()

	if actualResult != expectedResult {
		t.Errorf("%v but got %v", expectedResult, actualResult)
	}
}

func TestTotalBalanceIs_10_ShouldChangeCoin_T(t *testing.T) {
	expectedResult := ", T"
	v := vm.NewVendingMachine()
	v.TotalCoins = 10

	actualResult := v.ChangeCoins()

	if actualResult != expectedResult {
		t.Errorf("%v but got %v", expectedResult, actualResult)
	}
}

func TestTotalBalanceIs_7_ShouldChangeCoin_F_TW(t *testing.T) {
	expectedResult := ", F, TW"
	v := vm.NewVendingMachine()
	v.TotalCoins = 7

	actualResult := v.ChangeCoins()

	if actualResult != expectedResult {
		t.Errorf("%v but got %v", expectedResult, actualResult)
	}
}

func TestTotalBalanceIs_4_ShouldChangeCoin_TW_TW(t *testing.T) {
	expectedResult := ", TW, TW"
	v := vm.NewVendingMachine()
	v.TotalCoins = 4

	actualResult := v.ChangeCoins()

	if actualResult != expectedResult {
		t.Errorf("%v but got %v", expectedResult, actualResult)
	}
}

func TestTotalBalanceIs_25_ShouldChangeCoin_T_T_F(t *testing.T) {
	expectedResult := ", T, T, F"
	v := vm.NewVendingMachine()
	v.TotalCoins = 25

	actualResult := v.ChangeCoins()

	if actualResult != expectedResult {
		t.Errorf("%v but got %v", expectedResult, actualResult)
	}
}
