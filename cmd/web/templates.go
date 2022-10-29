package main

import "snippetbox.KazbekTutinbay.net/internal/models"

type templateData struct {
	Snippet  *models.Snippet
	Snippets []*models.Snippet
}
