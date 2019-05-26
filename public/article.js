marked.setOptions({
    kaTex: katex
});

function graphql(options) {
    let {query = ''} = options

    return fetch('/graphql/', {
        method: 'POST',
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ query: query}),
    }).then(res => res.json())
}

$el = document.querySelector('#article')


if(article_id) {
    graphql({
        query: `{ article(id: "${article_id}") { id content } }`
    }).then(data => {
        let content = marked(data.article.content)
        console.log(content)
        $el.innerHTML = content
    })
} else {
    // get list
    graphql({
        query: '{ list {id content} }',
    }).then(data => {
        console.log(data)
    })
    console.log("no article")
}
