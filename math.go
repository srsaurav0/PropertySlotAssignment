package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	input := map[string]float64{
		"11": 0,
		"12": 40,
		"24": 60,
	}
	input = validate_input(input)
	output := assign_feed(input, 7, "12-11-24")
	fmt.Println("Input percentage map: ", input)
	fmt.Println("Output feed map: ", output)
}

func validate_input(input map[string]float64) map[string]float64 {
	sum := 0.0
	mod_input := make(map[string]float64)
	for feed, percentage := range input {
		if percentage != 0 {
			mod_input[feed] = percentage
			sum += percentage
		}
	}
	if sum != 100.00 {
		for feed, percentage := range mod_input {
			mod_input[feed] = percentage / sum * 100.00
		}
	}
	return mod_input
}

func assign_feed(input map[string]float64, limit int, priority_feed string) map[string]int {
	percentage_float := make(map[string]float64) // Float value assignment
	percentage_int := make(map[string]int)       // Final output map
	sum := 0
	for feed, percentage := range input {
		percentage_float[feed] = percentage * float64(limit) / 100
		percentage_int[feed] = int(percentage_float[feed])
		percentage_float[feed] = percentage_float[feed] - float64(percentage_int[feed])
		sum += percentage_int[feed]
	}
	// fmt.Println("Previous percentage float map: ", percentage_float)
	// fmt.Println("Previous percentage int map: ", percentage_int)

	// fmt.Println("Sum is: ", sum)
	if sum == limit {
		return percentage_int
	} else if sum < limit {
		priority_feed_map := make(map[string]int)  // Map to store the feed value and priority. Lower value is higher priority
		feeds := strings.Split(priority_feed, "-") // Split the string values and get feed
		// fmt.Println(feeds)

		for i := range feeds {
			priority_feed_map[feeds[i]] = i
		}
		// fmt.Println("Priority feed map: ", priority_feed_map)

		for sum < limit {
			type percentage_float_struct struct { // Struct to sort the percentage value
				Key   string
				Value float64
			}

			var percentage_sorted []percentage_float_struct
			for k, v := range percentage_float { // Variable of type struct
				percentage_sorted = append(percentage_sorted, percentage_float_struct{k, v})
			}

			sort.Slice(percentage_sorted, func(i, j int) bool { // Sort function
				return percentage_sorted[i].Value > percentage_sorted[j].Value
			})

			// fmt.Println(percentage_sorted)

			feed_mark := map[string]bool{}
			for i := range percentage_sorted {
				feed_mark[percentage_sorted[i].Key] = false
			}

			// fmt.Println("feed_mark: ", feed_mark)

			feed_to_update := ""
			i := 0
			for i <= len(percentage_sorted) {
				// fmt.Println("Length: ", len(percentage_sorted))
				// fmt.Println("i: ", i)
				if i == len(percentage_sorted)-1 {
					if feed_mark[percentage_sorted[i].Key] == false {
						feed_to_update = percentage_sorted[i].Key
						// fmt.Println("feed_to_update: ", feed_to_update)
						feed_mark[percentage_sorted[i].Key] = true
						// fmt.Println("1")
					}
				} else if percentage_sorted[i].Value == percentage_sorted[i+1].Value && feed_mark[percentage_sorted[i].Key] == false {
					if feed_to_update == "" {
						feed_to_update = percentage_sorted[i].Key
						feed_mark[percentage_sorted[i].Key] = true
						// fmt.Println("2")
					}
					if priority_feed_map[percentage_sorted[i+1].Key] < priority_feed_map[feed_to_update] && feed_mark[percentage_sorted[i+1].Key] == false {
						feed_to_update = percentage_sorted[i+1].Key
						feed_mark[percentage_sorted[i+1].Key] = true
						feed_mark[percentage_sorted[i].Key] = false
						i--
						// fmt.Println("3")
						// fmt.Println("i mod: ", i)
					}
				} else {
					if feed_to_update == "" && feed_mark[percentage_sorted[i].Key] == false {
						feed_to_update = percentage_sorted[i].Key
						feed_mark[percentage_sorted[i].Key] = true
						// fmt.Println("4")
					}
					// break
				}
				if feed_to_update != "" {
					percentage_int[feed_to_update]++
					sum++
				}
				// fmt.Println("Feed to update: ", feed_to_update)
				feed_to_update = ""
				i++
				if sum == limit {
					break
				}
			}
		}
		return percentage_int

	} else {
		return percentage_int
	}

}
