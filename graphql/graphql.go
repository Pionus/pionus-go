package graphql

import (
    "github.com/pionus/arry"
    graphql "github.com/graph-gophers/graphql-go"

    "./resolver"
    "./schema"
)


type GraphqlParam struct {
    Query string `json:"query"`
    OperationName string `json:"operationName,omitempty"`
    Variables map[string]interface{} `json:"variables,omitempty"`
}



func GetController() arry.Handler {
    s := graphql.MustParseSchema(schema.String(), resolver.New())

    return func(ctx arry.Context) {
        var param GraphqlParam

        ctx.Decode(&param)

        result := s.Exec(
            ctx.Request().Context(),
            param.Query,
            param.OperationName,
            param.Variables,
        )

        if len(result.Errors) != 0 {
            ctx.Reply(502)
            return
        }

        ctx.JSON(200, result.Data)
    }

}
