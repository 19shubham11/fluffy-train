package mysql

import (
	"github.com/19shubham11/snippetbox/pkg/models"
	"reflect"
	"testing"
	"time"
)

func TestInsert(t *testing.T) {
	db, teardown := newTestDB(t)
	defer teardown()
	snippet := SnippetModel{db}
	t.Run("Insert into the db", func(t *testing.T) {
		// insert once
		id1, err := snippet.Insert("some title", "some content", "7")
		if err != nil {
			t.Fatal(err)
		}

		want1 := 1
		if id1 != want1 {
			t.Errorf("Got %d, want %d", id1, want1)
		}

		// insert again
		id2, err := snippet.Insert("some title", "some content", "7")
		if err != nil {
			t.Fatal(err)
		}

		want2 := 2
		if id2 != want2 {
			t.Errorf("Got %d, want %d", id2, want2)
		}
	})
}

func TestGet(t *testing.T) {
	db, teardown := newTestDB(t)
	defer teardown()
	snippet := SnippetModel{db}

	t.Run("Return an existing recrod for a valid id", func(t *testing.T) {
		id, err := snippet.Insert("Title1", "some content", "7")
		if err != nil {
			t.Fatal(err)
		}

		gotModel, err := snippet.Get(id)
		if err != nil {
			t.Fatal(err)
		}

		utcFormatWithoutMilliseconds := "2006-01-02 15:04:05 +0000 UTC"

		expectedModel := models.Snippet{
			ID:      1,
			Title:   "Title1",
			Content: "some content",
			Created: time.Now().UTC(),
			Expires: time.Now().UTC(),
		}

		// properly assert these, time is causing troubles :shrug:
		if gotModel.ID != expectedModel.ID {
			t.Errorf("Got %v want %v", gotModel, expectedModel)
		}
		if gotModel.Title != expectedModel.Title {
			t.Errorf("Got %v want %v", gotModel, expectedModel)
		}
		if gotModel.Created.Format(utcFormatWithoutMilliseconds) != expectedModel.Created.Format(utcFormatWithoutMilliseconds) {
			t.Errorf("Got %v want %v", gotModel, expectedModel)
		}

	})

	t.Run("Return error for an invalid id", func(t *testing.T) {
		got, err := snippet.Get(11)

		if !reflect.DeepEqual(err, models.ErrNoRecord) {
			t.Errorf("Expected error %v got %v", err, got)
		}
	})
}
