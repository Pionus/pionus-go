package schema

import (
    // "bytes"
    "io/ioutil"
    "log"
)

// func String() string {
//     buf := bytes.Buffer{}
//
//     for _, name := range AssetNames() {
//         b := MustAsset(name)
//         buf.Write(b)
//
//         if len(b) > 0 && b[len(b) - 1] != '\n' {
//             buf.WriteByte('\n')
//         }
//     }
//
//     return buf.String()
// }

func String() string {
    b, err := ioutil.ReadFile("schema/schema.graphql")

    if err != nil {
        log.Print(err)
        return ""
    }

    return string(b)
}
