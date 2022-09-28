function register()
{  
   var myHeaders = new Headers();
   myHeaders.append("User-Agent", "Apifox/1.0.0 (https://www.apifox.cn)");
   myHeaders.append("Content-Type", "application/json");
   var axios = require('axios');
   var data = JSON.stringify({
   "name": document.getElementById('name').value,
   "email": document.getElementById('email').value,
   "password": document.getElementById('password').value
});
var config = {
   method: 'post',
   url: 'http://180.76.167.8:8888/api/user/register',
   headers: { 
      'User-Agent': 'Apifox/1.0.0 (https://www.apifox.cn)', 
      'Content-Type': 'application/json'
   },
   data : data
};

axios(config)
.then(function (response) {
   alert("你好");
   console.log(JSON.stringify(response.data));
   alert(JSON.stringify(response.data));
})
.catch(function (error) {
   console.log(error);
});
alert("你好");
}