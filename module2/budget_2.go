package module2

import (
	"errors"
	"time"
	//log "github.com/sirupsen/logrus"
)

// START Initial code
const Environment = "dev"

// Budget stores Budget information
type Budget struct {
	Max   float32
	Items []Item
}

// Item stores Item information
type Item struct {
	Description string
	Price       float32
}

var report map[time.Month]*Budget

// InitializeReport creates an empty map
// to store each budget
func InitializeReport() {
	report = make(map[time.Month]*Budget)
}

func init() {
	// Create a new instance of the logger. You can have any number of instances.
	InitializeReport()
}

// CurrentCost returns how much we've added
// to the current budget
func (b Budget) CurrentCost() float32 {
	var sum float32
	for _, item := range b.Items {
		sum += item.Price
	}
	return sum
}

var errDoesNotFitBudget = errors.New("Item does not fit the budget")

var errReportIsFull = errors.New("Report is full")

var errDuplicateEntry = errors.New("Cannot add duplicate entry")

// END Initial code

// START Project code

// AddItem adds an item to the current budget
func (b *Budget) AddItem(description string, price float32) error {

	newItem := Item{Description: description, Price: price}
	//log.WithFields(log.Fields{
	//	"description": description,
	//	"price": price,
	//	"max": b.Max,
	//	"current_cost": b.CurrentCost(),
	//}).Info("Adding item")
	if b.CurrentCost() + price > b.Max {
		return errDoesNotFitBudget
	}
	b.Items = append(b.Items, newItem)
	return nil
}

// RemoveItem removes a given item from the current budget
func (b *Budget) RemoveItem(description string) {
	//log.WithFields(log.Fields{
	//	"description": description,
	//}).Info("Remove Item")
	for i := range b.Items {
		if b.Items[i].Description == description {
			b.Items = append(b.Items[:i], b.Items[i+1:]...)
			break
		}
	}
}

// CreateBudget creates a new budget with a specified max
func CreateBudget(month time.Month, max float32) (*Budget, error) {
	var newBudget *Budget

	if len(report) >= 12 {
		log.Info(len(report))
		return nil, errReportIsFull
	}

	if _, hasEntry := report[month]; hasEntry {
		//log.Info(hasEntry)
		return nil, errDuplicateEntry
	}

	newBudget = &Budget{
		Max:   max,
		Items: nil,
	}

	//log.WithFields(log.Fields{
	//	"new_budget_addr": &newBudget,
	//	"new_budget": newBudget,
	//}).Info("Create Budget")

	report[month] = newBudget
	//log.WithFields(log.Fields{
	//	"new_budget": newBudget,
	//	"report": report,
	//	"month_report": report[month],
	//}).Info("Budget Added To Report")

	return newBudget, nil
}

// GetBudget returns budget for given month
func GetBudget(month time.Month) *Budget {

	if budget, ok := report[month]; ok != false {
		//log.WithFields(log.Fields{
		//	"budget": budget,
		//}).Info("Getting budget")
		return budget
	}

	return nil
}

// END Project code
