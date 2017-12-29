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

fetch("/md/20171209.md", {
    credentials: 'include'
}).then(response => {
    return response.text()
}).then(text => {
    $el.innerHTML = marked(text)
})
