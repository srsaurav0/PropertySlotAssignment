package main

import (
	"fmt"
	"sort"
	"strings"
)

// Main function
func main() {
	input := map[string]float64{
		"11": 30.0,
		"12": 30.0,
		"24": 40.0,
	}
	input = validate_input(input)
	output := assign_feed(input, 9, "12-11-24")
	fmt.Println("Input percentage map: ", input)
	fmt.Println("Output feed map: ", output)
}

// Check for invalid inputs and modify the input
func validate_input(input map[string]float64) map[string]float64 {
	sum := 0.0
	mod_input := make(map[string]float64)

	// Remove any input if percentage is 0
	for feed, percentage := range input {
		if percentage != 0 {
			mod_input[feed] = percentage
			sum += percentage
		}
	}

	// Convert to 100% if total percentage is not 100%
	if sum != 100.00 {
		for feed, percentage := range mod_input {
			mod_input[feed] = percentage / sum * 100.00
		}
	}
	return mod_input
}

// Function to assign feed
func assign_feed(input map[string]float64, limit int, priority_feed string) map[string]int {
	percentage_float := make(map[string]float64) // Float value assignment
	percentage_int := make(map[string]int)       // Final output map
	sum := 0

	// Assign the float, int and calculate sum
	for feed, percentage := range input {
		percentage_float[feed] = percentage * float64(limit) / 100                      //Initial assignment
		percentage_int[feed] = int(percentage_float[feed])                              // Convert float to int
		percentage_float[feed] = percentage_float[feed] - float64(percentage_int[feed]) // Take only the remainder part
		sum += percentage_int[feed]
	}
	// fmt.Println("Previous percentage float map: ", percentage_float)
	// fmt.Println("Previous percentage int map: ", percentage_int)

	// fmt.Println("Sum is: ", sum)

	// If sum == limit, then return the output (percentage_int)
	if sum == limit {
		return percentage_int
	} else {
		priority_feed_map := make(map[string]int)  // Map to store the feed value and priority. Lower value is higher priority
		feeds := strings.Split(priority_feed, "-") // Split the string values and get feed

		// Assign priority to feeds (Lower value means greater priority)
		for i := range feeds {
			priority_feed_map[feeds[i]] = i
		}

		for sum < limit {
			type percentage_float_struct struct { // Struct to sort the percentage value
				Key   string
				Value float64
			}

			var percentage_sorted []percentage_float_struct // Variable to sort the map values by float value
			// Copy the values initially
			for k, v := range percentage_float {
				percentage_sorted = append(percentage_sorted, percentage_float_struct{k, v})
			}

			// Sort the values in the structure by float value
			sort.Slice(percentage_sorted, func(i, j int) bool {
				return percentage_sorted[i].Value > percentage_sorted[j].Value
			})

			// Boolean map to check if a feed is incremented
			feed_mark := map[string]bool{}
			for i := range percentage_sorted {
				feed_mark[percentage_sorted[i].Key] = false
			}

			// Initial assignment
			feed_to_update := "" // Variable to mark which feed to increment
			i := 0
			// Traverse the sorted variable and increment the appropriate feed
			for i <= len(percentage_sorted) {
				// Check if the last element
				if i == len(percentage_sorted)-1 {
					// Check if already incremented
					if !feed_mark[percentage_sorted[i].Key] {
						feed_to_update = percentage_sorted[i].Key
						feed_mark[percentage_sorted[i].Key] = true
					}
					// Check if the percentage are same and apply priority
				} else if percentage_sorted[i].Value == percentage_sorted[i+1].Value && !feed_mark[percentage_sorted[i].Key] {
					// Initially select the first one to increment
					if feed_to_update == "" {
						feed_to_update = percentage_sorted[i].Key
						feed_mark[percentage_sorted[i].Key] = true
					}
					// Check if the second one has higher priority (lower value) and increment it
					if priority_feed_map[percentage_sorted[i+1].Key] < priority_feed_map[feed_to_update] && !feed_mark[percentage_sorted[i+1].Key] {
						feed_to_update = percentage_sorted[i+1].Key  // Select to update
						feed_mark[percentage_sorted[i+1].Key] = true // Make the feed visited
						feed_mark[percentage_sorted[i].Key] = false  // Make the previous feed unvisited
						i--                                          // Custom decrement to check the previous feed again
					}
				} else {
					// Different percentage, so increment the one with higher percentage
					if feed_to_update == "" && !feed_mark[percentage_sorted[i].Key] {
						feed_to_update = percentage_sorted[i].Key
						feed_mark[percentage_sorted[i].Key] = true
					}
				}
				// Check for changes and then update the feed
				if feed_to_update != "" {
					percentage_int[feed_to_update]++
					sum++
				}
				feed_to_update = "" // Initialize the feed variable again
				i++
				// If limit reached, break
				if sum == limit {
					break
				}
			}
		}
		return percentage_int // Return the updated map
	}

}
