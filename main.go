package main

import "fmt"

type Customer struct {
	ID            int
	Name          string
	Email         string
	Phone         string
	WebsiteVisits []string
	AppUsage      []string
	Newsletter    bool
}

// Data collection: Simulate data collection from various sources
func collectData() []Customer {
	return []Customer{
		{ID: 1, Name: "John Doe", Email: "john.doe@example.com", Phone: "111111111", WebsiteVisits: []string{"Visited homepage", "Read politics article"}, AppUsage: []string{"Read sports article"}, Newsletter: true},
		{ID: 2, Name: "Jane Smith", Email: "jane.smith@example.com", Phone: "888888888", WebsiteVisits: []string{"Visited homepage", "Read tech article"}, AppUsage: []string{"Read tech article", "Clicked notification"}, Newsletter: false},
	}
}

// Unify data: Combine data to form a complete user profile
func unifyData(customers []Customer) map[int]Customer {
	userProfiles := make(map[int]Customer)
	for _, customer := range customers {
		if existing, exists := userProfiles[customer.ID]; exists {
			existing.WebsiteVisits = append(existing.WebsiteVisits, customer.WebsiteVisits...)
			existing.AppUsage = append(existing.AppUsage, customer.AppUsage...)
			userProfiles[customer.ID] = existing
		} else {
			userProfiles[customer.ID] = customer
		}
	}
	return userProfiles
}

// Segment users: For simplicity, segment based on interests and engagement
func segmentUsers(profiles map[int]Customer) map[string][]Customer {
	segments := make(map[string][]Customer)
	for _, profile := range profiles {
		if profile.Newsletter {
			segments["Newsletter Subscribers"] = append(segments["Newsletter Subscribers"], profile)
		}
		for _, visit := range profile.WebsiteVisits {
			if visit == "Read politics article" {
				segments["Politics Readers"] = append(segments["Politics Readers"], profile)
			}
			if visit == "Read tech article" {
				segments["Tech Readers"] = append(segments["Tech Readers"], profile)
			}
		}
		for _, usage := range profile.AppUsage {
			if usage == "Read sports article" {
				segments["Sports Readers"] = append(segments["Sports Readers"], profile)
			}
		}
	}
	return segments
}

// Output personalized messages based on segments
func sendPersonalizedContent(segments map[string][]Customer) {
	for segment, users := range segments {
		fmt.Printf("Segment: %s\n", segment)
		for _, user := range users {
			fmt.Printf("Sending personalized content to %s at %s\n", user.Name, user.Email)
		}
	}
}

func main() {
	// Step 1: Collect user data
	customers := collectData()

	// Step 2: Unify user data
	userProfiles := unifyData(customers)

	// Step 3: Segment users
	segments := segmentUsers(userProfiles)

	// Step 4: Send personalized content
	sendPersonalizedContent(segments)
}
