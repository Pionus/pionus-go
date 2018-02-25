$(() => {
    $("#form").on("submit", e => {
        let email = $("#email").val()
        let content = $("#content").val()

        data = new
        $.ajax({
            type: 'POST',
            url: '/dreams',
            data: {
                email,
                content,
            },
            dataType: 'json',
            success(data) {
                if(data.code == 200) {
                    alert("OK!")
                    window.location.reload()
                }
            }
        })

        return false
    })
})
