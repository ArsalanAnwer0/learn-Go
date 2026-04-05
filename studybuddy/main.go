package main

import "fmt"

type Topic struct {
	Name       string
	Understood bool
	Notes      string
}

func listTopics(topics []Topic) {
	for _, topic := range topics {
		fmt.Printf("- %s - Understood: %t\n   Notes: %s\n", topic.Name, topic.Understood, topic.Notes)
	}
}

func addTopic(topics []Topic, name string, notes string) []Topic {
	topics = append(topics, Topic{Name: name, Understood: false, Notes: notes})
	return topics
}

func markUnderstood(topics []Topic, name string) []Topic {
	// if name is name in topic then like mark it as understood
	for i, topic := range topics {
		if topic.Name == name {
			topics[i].Understood = true
			break
		}
	}
	return topics
}

func listPendingTopics(topics []Topic) {
	pending := pendingTopics(topics)
	if len(pending) == 0 {
		fmt.Println("All topics are understood! Great job!")
		return
	} else {
		fmt.Println("Pending Topics:")
		for _, topic := range pending {
			fmt.Printf("- %s\n   Notes: %s\n", topic.Name, topic.Notes)
		}
	}
}

func pendingTopics(topics []Topic) []Topic {
	// create a new slice of topics that are not understood
	var pending []Topic
	// loop through topics and if topic is not understood then add it to pending slice
	for _, topic := range topics {
		if !topic.Understood {
			pending = append(pending, topic)
		}
	}
	return pending
}

func listUnderstoodTopics(topics []Topic) {
	understood := understoodTopics(topics)
	if len(understood) == 0 {
		fmt.Println("No topics are understood yet.")
		return
	} else {
		fmt.Println("Understood Topics:")
		for _, topic := range understood {
			fmt.Printf("- %s\n   Notes: %s\n", topic.Name, topic.Notes)
		}
	}
}

func understoodTopics(topics []Topic) []Topic {
	var understood []Topic
	// loop through topics and if topic is not understood then add it to pending slice
	for _, topic := range topics {
		if topic.Understood {
			understood = append(understood, topic)
		}
	}
	return understood
}

func menu(topics *[]Topic) {
	fmt.Println("Main Menu")
	for {
		fmt.Println("\n1. List Topics")
		fmt.Println("2. List Pending Topics")
		fmt.Println("3. List Understood Topics")
		fmt.Println("4. Add Topic")
		fmt.Println("5. Mark Topic as Understood")
		fmt.Println("6. Exit")
		fmt.Println("Enter choice: ")
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			// listTopics(topics)
			listTopics(*topics)
		case 2:
			// listPendingTopics(topics)
			listPendingTopics(*topics)

		case 3:
			// listUnderstoodTopics(topics)
			listUnderstoodTopics(*topics)
		case 4:
			// addTopic(topics)
			fmt.Println("Enter Name: ")
			var name string
			fmt.Scanln(&name)
			fmt.Println("Enter Notes: ")
			var notes string
			fmt.Scanln(&notes)
			*topics = addTopic(*topics, name, notes)

		case 5:
			// mark topic as understood
			fmt.Println("Enter Name: ")
			var name string
			fmt.Scanln(&name)
			*topics = markUnderstood(*topics, name)
		case 6:
			fmt.Println("Exiting Study Buddy. Happy Studying!")
			return

		default:
			fmt.Println("Invalid choice, please try again.")
			continue
		}
	}
}

func main() {
	fmt.Println("Welcome to Study Buddy!")
	// creating a slice
	topics := []Topic{
		{Name: "Go Basics", Understood: false, Notes: "Need to review slices and maps."},
		{Name: "Concurrency in Go", Understood: false, Notes: "Focus on goroutines and channels."},
		{Name: "Error Handling", Understood: false, Notes: "Practice with custom error types."},
	}
	menu(&topics)
}
