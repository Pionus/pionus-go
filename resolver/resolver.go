package resolver

import (
    // "log"
    "context"
    // "github.com/kataras/iris"
    "github.com/pionus/pionus-go/services"
    // "golang.org/x/net/context"
)

type Resolver struct {

}

// func GetRootResolver() (*Resolver, error) {
//     return &Resolver{}, nil
// }


func (r *Resolver) Article(ctx context.Context, args struct {
    ID string
}) (*ArticleResolver, error) {
    article, err := services.GetArticleByID(args.ID)

    if err != nil {
        return &ArticleResolver{}, err
    }

    return &ArticleResolver{ article }, nil
}

func (r *Resolver) List(ctx context.Context) (*[]*ArticleResolver, error) {
    articles, err := services.GetArticleList()

    if err != nil {
        panic(err)
    }

    var list []*ArticleResolver

    for _, article := range *articles {
        list = append(list, &ArticleResolver{ article })
    }

    return &list, nil
}
