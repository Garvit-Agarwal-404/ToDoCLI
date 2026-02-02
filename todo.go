package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

// struct declare kiya hai
type Todo struct {
	Title        string
	Compeleted   bool
	CreatedAt    time.Time
	CompeletedAt *time.Time
}

type Todos []Todo // apan ne ek Todo type  ka slice banaya bhai typed as Typos

// ek add method banaya hai to add a task
func (T *Todos) add(title string) {
	todo := Todo{
		Title:        title,
		Compeleted:   false,
		CompeletedAt: nil,
		CreatedAt:    time.Now(),
	}

	*T = append(*T, todo)

}

// to check index for operations like remove, edit,or toggle to see if  it is valid
func (T *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*T) {
		err := errors.New("invalid index")
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (T *Todos) delete(index int) error {
	t := *T // to use *T without needing to dereference again and again .. but if you want to make changes in T, you have to use *T only ..
	if err := T.validateIndex(index); err != nil {
		return err

	}
	*T = append(t[:index], t[index+1:]...)
	return nil
}

func (T *Todos) toggle(index int) error {
	t := *T
	if err := t.validateIndex(index); err != nil {
		return err
	}
	isCompleted := t[index].Compeleted
	if !isCompleted {
		completiontime := time.Now()
		t[index].CompeletedAt = &completiontime
	}
	t[index].Compeleted = !isCompleted
	return nil
}

func (T *Todos) edit(index int, title string) error {
	t := *T
	if err := t.validateIndex(index); err != nil {
		return err
	}

	t[index].Title = title
	return nil
}

func (todos *Todos) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")

	for index, t := range *todos {
		completed := "❌"
		completedAt := ""

		if t.Compeleted {
			completed = "✅"
			if t.CompeletedAt != nil {
				completedAt = t.CompeletedAt.Format(time.RFC1123)

			}

		}

		table.AddRow(strconv.Itoa(index), t.Title, completed, t.CreatedAt.Format(time.RFC1123), completedAt)

	}

	table.Render()
}
