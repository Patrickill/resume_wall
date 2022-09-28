function reg(){
   axios.post('http://180.76.167.8:8888/api/user/register', {
      "name": document.getElementById('name').value,
      "email": document.getElementById('email').value,
      "password": document.getElementById('password').value
   })
       .then(function (response) {
          alert("你好");
          console.log(JSON.stringify(response.data));
          alert(JSON.stringify(response.data));
       })
       .catch(function (error) {
          console.log(error);
       });
}