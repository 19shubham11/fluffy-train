package mysql

import (
	"database/sql"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/19shubham11/snippetbox/pkg/models"
)

var db *sql.DB

func TestMain(m *testing.M) {
	var teardown func()
	db, teardown = newTestDB()
	defer teardown()
	code := m.Run()
	os.Exit(code)
}

func TestInsert(t *testing.T) {
	snippet := SnippetModel{db}
	t.Run("Insert into the db", func(t *testing.T) {
		// insert once
		_, err := snippet.Insert("Title1", "some content 1", "7")
		if err != nil {
			t.Fatal(err)
		}
		// insert again
		_, err = snippet.Insert("Title2", "some content 2", "11")
		if err != nil {
			t.Fatal(err)
		}
	})
}

func TestGet(t *testing.T) {
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
			ID:      id,
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
		got, err := snippet.Get(11111)

		if !reflect.DeepEqual(err, models.ErrNoRecord) {
			t.Errorf("Expected error %v got %v", err, got)
		}
	})
}
