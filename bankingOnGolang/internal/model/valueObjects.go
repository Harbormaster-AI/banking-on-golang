
package model


import "github.com/shopspring/decimal"

import (
)
import (
)
import (
)
import (
)
import (
)
import (
)
import (
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

type AccountNumber struct {
     Value            string

// parent associations as their child

}

type IBAN struct {
     Value            string

// parent associations as their child

}

type BIC struct {
     Value            string

// parent associations as their child

}

type CardPAN struct {
     Value            string

// parent associations as their child

}

type Percentage struct {
     Value            decimal.Decimal

// parent associations as their child

}

