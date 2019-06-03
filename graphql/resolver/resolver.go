package resolver

import (
    "context"

    "../../services"
)

type Resolver struct {}


func New() *Resolver {
    return &Resolver{}
}


func (r *Resolver) List(ctx context.Context) (*[]*ArticleResolver, error) {
    l := make([]*ArticleResolver, 1)
    return &l, nil
}

func (r *Resolver) Article(ctx context.Context, args struct { ID string }) (*ArticleResolver, error) {
    article, err := services.GetArticleByID(args.ID)

    if err != nil {
        return &ArticleResolver{}, err
    }

    return &ArticleResolver{ article }, nil
}
