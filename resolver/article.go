package resolver

import (
    "github.com/pionus/pionus-go/models"

    graphql "github.com/graph-gophers/graphql-go"
)

type ArticleResolver struct {
    article *models.Article
}

func (r *ArticleResolver) ID() graphql.ID {
    return graphql.ID(r.article.ID)
}

func (r *ArticleResolver) Content() string {
    return r.article.Content
}
