const { default: axios } = require("axios");
function send()
{
    var data=JSON.stringify(
        {
            "name": document.getElementById('#').value,
            "message": document.getElementById('#').value,
        }
    );
    var config={
        method: 'POST',
        url: 'http://180.76.167.8:8888/api/message/add',
        data: data
    };
    
    axios(config)
    .then(function(response)
    {
        console.log(JSON.stringify(response.data));
        alert("发送成功！");
    })
    .catch(function(error)
    {
        console.log(error);
        alert(error);
    })
}