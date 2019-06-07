package vending_test

import (
	vm "./vendingmachine"
	"testing"
)

func TestBuy_SD_WithTotalBalance_18_BathShouldReturn_SD(t *testing.T) {
	v := vm.NewVendingMachine()
	expectedResult := "SD"
	v.TotalCoins = 18

	actualResult := v.BuyDrink("SD")

	if actualResult != expectedResult {
		t.Errorf("%v but got %v", expectedResult, actualResult)
	}
}

func TestBuy_CC_WithTotalBalance_12_BathShouldReturn_CC(t *testing.T) {
	expectedResult := "CC"
	v := vm.NewVendingMachine()
	v.TotalCoins = 12

	actualResult := v.BuyDrink("CC")

	if actualResult != expectedResult {
		t.Errorf("%v but got %v", expectedResult, actualResult)
	}

	if v.ShowTotalBalance() != 0 {
		t.Errorf("Balance is not 0 but got %v", v.ShowTotalBalance())
	}
}

func TestBuy_XYZ_WithTotalBalance_12_BathShouldReturn_NoItem(t *testing.T) {
	expectedResult := "No Item"
	v := vm.NewVendingMachine()
	v.TotalCoins = 12

	actualResult := v.BuyDrink("XYZ")

	if actualResult != expectedResult {
		t.Errorf("%v but got %v", expectedResult, actualResult)
	}
}

func TestBuy_CC_WithTotalBalance_10_BathShouldReturn_AddMoreMoney(t *testing.T) {
	expectedResult := "Add more money"
	v := vm.NewVendingMachine()
	v.TotalCoins = 10

	actualResult := v.BuyDrink("CC")

	if actualResult != expectedResult {
		t.Errorf("%v but got %v", expectedResult, actualResult)
	}
}

func TestBuy_CC_WithTotalBalance_15_BathShouldReturn_CC_TW_O(t *testing.T) {
	expectedResult := "CC, TW, O"
	v := vm.NewVendingMachine()
	v.TotalCoins = 15

	actualResult := v.BuyDrink("CC")

	if actualResult != expectedResult {
		t.Errorf("%v but got %v", expectedResult, actualResult)
	}
}
