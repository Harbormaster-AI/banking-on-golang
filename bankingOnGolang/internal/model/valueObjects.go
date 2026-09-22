
package model

import (
    "github.com/shopspring/decimal"
)

 type Money struct {
     Amount            decimal.Decimal
    Currency            string

// parent associations as their child

}

 type Address struct {
     Street            string
    City            string
    State            string
    PostalCode            string
    Country            string

// parent associations as their child

}

 Name is value
type AccountNumber     string

 Name is value
type IBAN     string

 Name is value
type BIC     string

 Name is value
type CardPAN     string

 Name is value
type Percentage     decimal.Decimal

