marked.setOptions({
    // highlight: function(code, lang) {
    //   if (typeof lang === 'undefined') {
    //     return hljs.highlightAuto(code).value;
    //   } else if (lang === 'nohighlight') {
    //     return code;
    //   } else {
    //     return hljs.highlight(lang, code).value;
    //   }
    // },
    kaTex: katex
});

$el = document.querySelector('#article')


if(article_id) {
    fetch(`/md/${article_id}.md`, {
        credentials: 'include'
    }).then(response => {
        return response.text()
    }).then(text => {
        $el.innerHTML = marked(text)
    })
} else {
    console.log("no article")
}
