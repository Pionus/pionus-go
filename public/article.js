// marked.setOptions({
//     // highlight: function(code, lang) {
//     //   if (typeof lang === 'undefined') {
//     //     return hljs.highlightAuto(code).value;
//     //   } else if (lang === 'nohighlight') {
//     //     return code;
//     //   } else {
//     //     return hljs.highlight(lang, code).value;
//     //   }
//     // },
//     kaTex: katex
// });

const INLINE_DELIM = '$'
const BLOCK_DELIM = '$$'

function katex_inline_rule(state, silent) {
    if(silent) return false

    let delim_pos = state.pos + INLINE_DELIM.length
    let openDelim = state.src.slice(state.pos, delim_pos)
    if(openDelim !== INLINE_DELIM) return false

    if(delim_pos >= state.posMax) return false

    state.pos = delim_pos

    while(state.pos < state.posMax) {
        let closeDelim = state.src.slice(state.pos, state.pos + INLINE_DELIM.length)

        if(closeDelim === INLINE_DELIM && state.src[state.pos-1] !== '\\') break

        state.md.inline.skipToken(state)
    }

    // failed to find close delim
    if(state.pos >= state.posMax) {
        state.pos = delim_pos - INLINE_DELIM.length
        return false
    }

    let token = state.push('katex_inline', 'katex', 0)
    token.content = state.src.slice(delim_pos, state.pos)
    token.markup = INLINE_DELIM

    state.pos += INLINE_DELIM.length

    return true
}

function katex_block_rule(state, startLine, endLine, silent) {

    // TODO silent mode
    if(silent) return false

    // start position with white space
    let start = state.bMarks[startLine] + state.tShift[startLine]

    let delim_pos = start + BLOCK_DELIM.length
    if(delim_pos > state.eMarks[startLine]) return false

    let openDelim = state.src.slice(start, delim_pos)

    if(openDelim !== BLOCK_DELIM) return false

    let currLine = startLine

    while(currLine <= endLine) {
        let begin = state.bMarks[currLine] + state.tShift[currLine],
            end = state.eMarks[currLine]

        currLine ++

        // only open delim line
        if(currLine == startLine + 1 && end - begin <= BLOCK_DELIM.length) continue

        let line = state.src.slice(begin, end)

        if(line.trim().slice(-BLOCK_DELIM.length) !== BLOCK_DELIM) continue

        // TODO negative indent

        // found close delim, reset line
        currLine --
        break
    }

    let end_pos = state.bMarks[currLine]

    // found close delim
    if(currLine <= endLine) {
        end_pos += state.src.slice(
            state.bMarks[currLine],
            state.eMarks[currLine]
        ).lastIndexOf(BLOCK_DELIM)
    } else {
        // no close delim, auto closed
        end_pos = state.eMarks[endLine]
    }

    let token = state.push('katex_block', 'katex', 0)
    token.block = true
    token.content = state.src.slice(delim_pos, end_pos)
    token.markup = BLOCK_DELIM

    // set to the next line if not reach the end
    state.line = currLine < endLine ? currLine + 1 : endLine

    return true
}

function katex_inline(tokens, idx) {
    try {
        return katex.renderToString(tokens[idx].content.trim())
    } catch(err) {
        return '<span class="katex-error">${err.message}</span>'
    }
}

function katex_block(tokens, idx) {
    console.log(tokens[idx])
    try {
        return katex.renderToString(tokens[idx].content.trim(), {
            displayMode: true,
        })
    } catch(err) {
        return '<div class="katex-error">${err.message}</div>'
    }
}

const md = markdownit()


md.use(function(md, options) {
    md.block.ruler.after('blockquote', 'katex_block', katex_block_rule)
    md.inline.ruler.before('escape', 'katex_inline', katex_inline_rule)

    md.renderer.rules.katex_inline = katex_inline
    md.renderer.rules.katex_block = katex_block
})

$el = document.querySelector('#article')


if(article_id) {
    fetch(`/md/${article_id}.md`, {
        credentials: 'include'
    }).then(response => {
        return response.text()
    }).then(text => {
        $el.innerHTML = md.render(text)
    })
} else {
    console.log("no article")
}
