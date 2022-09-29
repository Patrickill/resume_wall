function login(){
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
   localStorage.setItem('token',response.data.data.token);
   alert('登陆成功');
   location.href="../../index/html/index_new.html";
})
.catch(function (error) {
   console.log(error);
   alert(error);
});
}