package swagger

import "testing"

func TestMonthlyRepayments(t *testing.T) {
	var calcLoanBod CalculateloanBody
	calcLoanBod.LoanAmount = 350000
	calcLoanBod.LoanType = "interest"
	calcLoanBod.PaymentFrequency = "monthly"
	calcLoanBod.InterestRate = 0.03
	calcLoanBod.LoanTerm = 1

	got := CalcInterestOnly(calcLoanBod)
	var want LoanRepayments
	want.MonthlyRepayments = 875
	want.TotalInterestPayable = 10500

	if got.MonthlyRepayments != want.MonthlyRepayments {
        t.Errorf("got %q want %q", got, want)
    }

	if got.TotalInterestPayable != want.TotalInterestPayable {
        t.Errorf("got %q want %q", got, want)
    }

	var loanrepamown LoanRepaymentsAmountOwing
	loanrepamown.Year = 0
	loanrepamown.Interest = 10500
	loanrepamown.Principal = 350000
	loanrepamown.Total = 360500

	want.AmountOwing = append(want.AmountOwing, loanrepamown)

	if len(got.AmountOwing) != len(want.AmountOwing) {
        t.Errorf("got %q want %q", got, want)
    }
}