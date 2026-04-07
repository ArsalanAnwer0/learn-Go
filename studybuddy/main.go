package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Topic struct {
	Name       string
	Understood bool
	Notes      string
}

type StudyBuddy struct {
	Topics   []Topic
	TopicMap map[string]*Topic
}

func NewStudyBuddy() *StudyBuddy {
	return &StudyBuddy{
		Topics:   []Topic{},
		TopicMap: make(map[string]*Topic),
	}
}

func listTopics(topics []Topic) {
	for _, topic := range topics {
		fmt.Printf("- %s - Understood: %t\n   Notes: %s\n", topic.Name, topic.Understood, topic.Notes)
	}
}

func (sb *StudyBuddy) AddTopic(name string, notes string) {
	sb.Topics = append(sb.Topics, Topic{Name: name, Understood: false, Notes: notes})
	sb.TopicMap[name] = &sb.Topics[len(sb.Topics)-1]
}

func (sb *StudyBuddy) DeleteTopic(name string) error {
	// check if topic exists in map
	_, exists := sb.TopicMap[name]
	if !exists {
		return fmt.Errorf("Topic '%s' not found", name)
	}
	// if exists then delete from map and slice
	delete(sb.TopicMap, name)
	for i, topic := range sb.Topics {
		if topic.Name == name {
			sb.Topics = append(sb.Topics[:i], sb.Topics[i+1:]...)
			break
		}
	}
	return nil
}

func (sb *StudyBuddy) MarkUnderstood(name string) error {
	topic, exists := sb.TopicMap[name]
	if !exists {
		return fmt.Errorf("Topic '%s' not found", name)
	}
	topic.Understood = true
	return nil // mark topic as understood in map and slice and return none
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
	// loop through topics and if topic is understood then add it to understood slice
	for _, topic := range topics {
		if topic.Understood {
			understood = append(understood, topic)
		}
	}
	return understood
}

func menu(sb *StudyBuddy) {
	fmt.Println("Main Menu")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n1. List Topics")
		fmt.Println("2. List Pending Topics")
		fmt.Println("3. List Understood Topics")
		fmt.Println("4. Add Topic")
		fmt.Println("5. Mark Topic as Understood")
		fmt.Println("6. Exit")
		fmt.Println("Enter choice: ")
		var choice int
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		choice, _ = strconv.Atoi(input)
		switch choice {
		case 1:
			// listTopics(topics)
			listTopics(sb.Topics)
		case 2:
			// listPendingTopics(topics)
			listPendingTopics(sb.Topics)

		case 3:
			// listUnderstoodTopics(topics)
			listUnderstoodTopics(sb.Topics)
		case 4:
			fmt.Println("Enter Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Println("Enter Notes: ")
			notes, _ := reader.ReadString('\n')
			notes = strings.TrimSpace(notes)
			sb.AddTopic(name, notes)

		case 5:
			// mark topic as understood
			fmt.Println("Enter Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)
			err := sb.MarkUnderstood(name)
			if err != nil {
				fmt.Println("Error: ", err)
			}
		case 6:
			fmt.Println("Enter Name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)
			err := sb.DeleteTopic(name)
			if err != nil {
				fmt.Println("Error: ", err)
			}
		case 7:
			fmt.Println("Exiting Study Buddy. Happy Studying!")
			return

		default:
			fmt.Println("Invalid choice, please try again.")
			continue
		}
	}
}

func main() {

	sb := NewStudyBuddy()

	fmt.Println("Welcome to Study Buddy!")
	// creating a slice
	sb.AddTopic("Go Basics", "Need to review slices and maps.")
	sb.AddTopic("Concurrency in Go", "Focus on goroutines and channels.")
	sb.AddTopic("Error Handling", "Practice with custom error types.")

	menu(sb)
}
