
var createURL = function(str) {
    var infoDiv = document.getElementById('info');

    axios({
        url: 's/',
        data: {
            link: str,
            recaptcha: grecaptcha.getResponse()
        },
        method: 'post',
        headers: {
            'Content-Type': 'application/json'
        }
    }).then(function(response) {
        infoDiv.style.display = 'block';
        infoDiv.style.color = 'black';
        infoDiv.innerHTML = 'Success! Your new URL is '
            + '<a href="' + response.data + '">' + response.data + '</a>';
    }).catch(function(error) {
        infoDiv.style.display = 'block';
        infoDiv.style.color = 'red';
        infoDiv.innerHTML = 'Error: ' + error.response.data;
    });
}

var onLinkFormSubmit = function(form) {
    event.preventDefault();
    createURL(form.link.value);
}
