package main

// import (
// 	"fmt"
// 	"github.com/daveshanley/vacuum/model"
// 	"github.com/daveshanley/vacuum/motor"
// 	"github.com/daveshanley/vacuum/rulesets"
// 	"io/ioutil"
// )

// func main() {

// 	// read in an OpenAPI Spec to a byte array
// 	specBytes, err := ioutil.ReadFile("api.json")
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	// build and store built-in vacuum default RuleSets.
// 	defaultRS := rulesets.BuildDefaultRuleSets()

// 	// generate the 'recommended' RuleSet
// 	recommendedRS := defaultRS.GenerateOpenAPIRecommendedRuleSet()

// 	// apply the rules in the ruleset to the specification
// 	lintingResults := motor.ApplyRulesToRuleSet(
// 		&motor.RuleSetExecution{
// 			RuleSet: recommendedRS,
// 			Spec:    specBytes,
// 		})

// 	// create a new model.RuleResultSet from the results.
// 	// structure allows categorization, sorting and searching
// 	// in a simple and consistent way.
// 	resultSet := model.NewRuleResultSet(lintingResults.Results)

// 	// sort results by line number (so they are not all jumbled)
// 	resultSet.SortResultsByLineNumber()

// 	//.. do something interesting with the results
// 	// print only the results from the 'schemas' category
// 	schemasResults := resultSet.GetRuleResultsForCategory("all")

// 	// for every rule that is violated, it contains a list of violations.
// 	// so first iterate through the schemas sesults
// 	for _, ruleResult := range schemasResults.RuleResults {

// 		// print out which rule was violated
// 		fmt.Printf("Rule: %s\n", ruleResult.Rule.Id)

// 		// iterate over each violation of this rule
// 		for _, violation := range ruleResult.Results {

// 			// print out the start line, column, violation message.
// 			fmt.Printf(" - [%d:%d] %s\n", violation.StartNode.Line,
// 				violation.StartNode.Column, violation.Message)
// 		}
// 	}
// }
