package resolver

import (
    "context"

    "github.com/pionus/pionus-go/services"
)

type Resolver struct {}


func New() *Resolver {
    return &Resolver{}
}


func (r *Resolver) List(ctx context.Context) (*[]*ArticleResolver, error) {
    list, _ := services.GetArticleList()

    var l []*ArticleResolver
    for _, article := range *list {
        l = append(l, &ArticleResolver{ article })
    }

    return &l, nil
}

func (r *Resolver) Article(ctx context.Context, args struct { ID string }) (*ArticleResolver, error) {
    article, err := services.GetArticleByID(args.ID)

    if err != nil {
        return &ArticleResolver{}, err
    }

    return &ArticleResolver{ article }, nil
}
