# home-loan-calculator
 
Use POST in postman with the following url https://sustained-opus-332900.ts.r.appspot.com/calculate-loan

The following input parameters were used in postman to check if the app is running
{
  "loan_amount": 30000,
  "loan_type": "interest",
  "payment_frequency": "monthly",
  "interest_rate": 0.03,
  "loan_term": 30
}

This gives us the following output in postman
{
    "monthly_repayments": 75,
    "total_interest_payable": 27000,
    "amount_owing": [
        {
            "interest": 27000,
            "principal": 30000,
            "total": 57000
        },
        {
            "year": 1,
            "interest": 26925,
            "principal": 30000,
            "total": 56925
        },
        {
            "year": 2,
            "interest": 26850,
            "principal": 30000,
            "total": 56850
        },
        {
            "year": 3,
            "interest": 26775,
            "principal": 30000,
            "total": 56775
        },
        {
            "year": 4,
            "interest": 26700,
            "principal": 30000,
            "total": 56700
        },
        {
            "year": 5,
            "interest": 26625,
            "principal": 30000,
            "total": 56625
        },
        {
            "year": 6,
            "interest": 26550,
            "principal": 30000,
            "total": 56550
        },
        {
            "year": 7,
            "interest": 26475,
            "principal": 30000,
            "total": 56475
        },
        {
            "year": 8,
            "interest": 26400,
            "principal": 30000,
            "total": 56400
        },
        {
            "year": 9,
            "interest": 26325,
            "principal": 30000,
            "total": 56325
        },
        {
            "year": 10,
            "interest": 26250,
            "principal": 30000,
            "total": 56250
        },
        {
            "year": 11,
            "interest": 26175,
            "principal": 30000,
            "total": 56175
        },
        {
            "year": 12,
            "interest": 26100,
            "principal": 30000,
            "total": 56100
        },
        {
            "year": 13,
            "interest": 26025,
            "principal": 30000,
            "total": 56025
        },
        {
            "year": 14,
            "interest": 25950,
            "principal": 30000,
            "total": 55950
        },
        {
            "year": 15,
            "interest": 25875,
            "principal": 30000,
            "total": 55875
        },
        {
            "year": 16,
            "interest": 25800,
            "principal": 30000,
            "total": 55800
        },
        {
            "year": 17,
            "interest": 25725,
            "principal": 30000,
            "total": 55725
        },
        {
            "year": 18,
            "interest": 25650,
            "principal": 30000,
            "total": 55650
        },
        {
            "year": 19,
            "interest": 25575,
            "principal": 30000,
            "total": 55575
        },
        {
            "year": 20,
            "interest": 25500,
            "principal": 30000,
            "total": 55500
        },
        {
            "year": 21,
            "interest": 25425,
            "principal": 30000,
            "total": 55425
        },
        {
            "year": 22,
            "interest": 25350,
            "principal": 30000,
            "total": 55350
        },
        {
            "year": 23,
            "interest": 25275,
            "principal": 30000,
            "total": 55275
        },
        {
            "year": 24,
            "interest": 25200,
            "principal": 30000,
            "total": 55200
        },
        {
            "year": 25,
            "interest": 25125,
            "principal": 30000,
            "total": 55125
        },
        {
            "year": 26,
            "interest": 25050,
            "principal": 30000,
            "total": 55050
        },
        {
            "year": 27,
            "interest": 24975,
            "principal": 30000,
            "total": 54975
        },
        {
            "year": 28,
            "interest": 24900,
            "principal": 30000,
            "total": 54900
        },
        {
            "year": 29,
            "interest": 24825,
            "principal": 30000,
            "total": 54825
        }
    ]
}
