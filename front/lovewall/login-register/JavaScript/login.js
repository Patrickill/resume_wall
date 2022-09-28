function login(){
   alert("你好");
   var axios = require('axios');
    var data = JSON.stringify({
   "email": document.getElementById('email').value,
   "password": document.getElementById('password').value
});

var config = {
   method: 'post',
   url: 'http://180.76.167.8:8888/api/user/login',
   data : data
};

axios(config)
.then(function (response) {
   console.log(JSON.stringify(response.data));
   alert(JSON.stringify(response.data));
})
.catch(function (error) {
   console.log(error);
});
}