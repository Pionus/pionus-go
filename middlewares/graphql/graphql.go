package graphql

import (
    // "time"
    // "context"
    "github.com/kataras/iris"
    graphql "github.com/graph-gophers/graphql-go"
)

type Options struct {
    Schema string
    Resolver interface{}
}

type Graphql struct {
    Schema *graphql.Schema
}

func New(options Options) iris.Handler {
    schema := graphql.MustParseSchema(options.Schema, options.Resolver)
    g := &Graphql{
        Schema: schema,
    }

    return g.Serve
}

type Params struct {
    Query         string                 `json:"query"`
    OperationName string                 `json:"operationName"`
    Variables     map[string]interface{} `json:"variables"`
}

func (g *Graphql) Serve(ctx iris.Context) {
    var params Params

    err := ctx.ReadJSON(&params)

    if err != nil {
        ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString(err.Error())
        ctx.Next()
        return
    }

    result := g.Schema.Exec(ctx.Request().Context(), params.Query, params.OperationName, params.Variables)

    if len(result.Errors) != 0 {
        ctx.StatusCode(iris.StatusBadRequest)
		ctx.WriteString("parse error")
        ctx.Next()
        return
    }

    ctx.Text(string(result.Data))

    ctx.Next()
}
