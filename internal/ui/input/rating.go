package ui_input

import (
	"strings"

	"scouts-risk/internal/model"
)

func ratingItems(r model.Rating) (items []model.RatingItem) {
  for i := r.Min; i <= r.Max; i++ {
    items = append(items, model.RatingItem{Index: i, R: r})
  }
  return
}

func radioClass(item model.RatingItem) string {
	classes := []string{
	"mask mask-star-2 bg-yellow-400",
	}
	if (item.R.Class != "") {
	classes = append(classes, item.R.Class)
	}
	if item.Index+1==item.R.Max {
	classes = append(classes, "checked")
	}
	return strings.Join(classes, " ")
}
