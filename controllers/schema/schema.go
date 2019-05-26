package schema

import (
    "io/ioutil"
)


func String() string {
    b, err := ioutil.ReadFile("controllers/schema/schema.graphql")

    if err != nil {
        panic(err)
    }

    return string(b)
}
