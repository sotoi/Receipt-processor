package rules

import (
	"fetch/models"
	"math"
	"strconv"
	"strings"
	"time"
)

// DefaultRules is a slice of Rule that contains the default rules
var DefaultRules = []Rule{
	&RetailerNameRule{},
	&RoundDollarRule{},
	&MultipleOfQuarterRule{},
	&TwoItemsRule{},
	&ItemDescriptionRule{},
	&OddDayRule{},
	&Between2And4Rule{},
}

//
type Rule interface {
	Apply(receipt *models.Receipt) int
}

// One point for every alphanumeric character in the retailer name.

type RetailerNameRule struct{}

func (r *RetailerNameRule) Apply(receipt *models.Receipt) int {
	count := 0
	for _, c := range receipt.Retailer {
		if c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c >= '0' && c <= '9' {
			count++
		}
	}
	return count
}

// 50 points if the total is a round dollar amount with no cents.

type RoundDollarRule struct{}

func (r *RoundDollarRule) Apply(receipt *models.Receipt) int {
	num, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		return 0
	}
	if num == float64(int(num)) {
		return 50
	}
	return 0
}

// 25 points if the total is a multiple of 0.25.

type MultipleOfQuarterRule struct{}

func (r *MultipleOfQuarterRule) Apply(receipt *models.Receipt) int {
	num, err := strconv.ParseFloat(receipt.Total, 64)
	if err != nil {
		return 0
	}
	if math.Mod(num, 0.25) == 0 {
		return 25
	}
	return 0
}

// 5 points for every two items on the receipt.

type TwoItemsRule struct{}

func (r *TwoItemsRule) Apply(receipt *models.Receipt) int {
	return len(receipt.Items) / 2 * 5
}

// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer. The result is the number of points earned.

type ItemDescriptionRule struct{}

func (r *ItemDescriptionRule) Apply(receipt *models.Receipt) int {
	points := 0
	for _, item := range receipt.Items {
		trimmedDescription := strings.TrimSpace(item.ShortDescription)
		if len(trimmedDescription) % 3 != 0 {
			continue
		}
		price, err := strconv.ParseFloat(item.Price, 64)
		if err != nil {
			continue
		}
		points += int(math.Ceil(price * 0.2))
	}
	return points
}
// 6 points if the day in the purchase date is odd.

type OddDayRule struct{}

func (r *OddDayRule) Apply(receipt *models.Receipt) int {
	purchaseDate, err := time.Parse(time.DateOnly, receipt.PurchaseDate)
	if err != nil {
		return 0
	}
	if purchaseDate.Day() % 2 == 1 {
		return 6
	}
	return 0
}


// 10 points if the time of purchase is after 2:00pm and before 4:00pm.

type Between2And4Rule struct{}

func (r *Between2And4Rule) Apply(receipt *models.Receipt) int {
	purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime)
	if err != nil {
		return 0
	}
	if purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
		return 10
	}
	return 0
}

