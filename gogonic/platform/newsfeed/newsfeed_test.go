package newsfeed

import "testing"

func TestAdd(t *testing.T) {
	feed := New()

	feed.Add(Item{})

	if len(feed.Items) == 0 || len(feed.Items) >= 2 {
		t.Errorf("Item either not added or added more than 1 times !\n Expected Behaviour not met!")
	}
}

func TestGetAll(t *testing.T) {

	feed := New()

	feed.Add(Item{"hello", "there"})

	results := feed.GetAll()

	if len(results) != 1 {
		t.Errorf("Items not added!")
	}
}
