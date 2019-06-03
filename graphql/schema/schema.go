package schema

import (
    "io/ioutil"
)


func String() string {
    b, err := ioutil.ReadFile("graphql/schema/schema.graphql")

    if err != nil {
        panic(err)
    }

    return string(b)
}
