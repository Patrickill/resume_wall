function showmessage()
{
  document.write('<script src="https://cdn.staticfile.org/axios/0.18.0/axios.min.js"></script>')
  alert("你好");
  const axios=require(axios)
  var conflg={
    method:'get',
    url:'http://180.76.167.8:8888/api/message/getR',
  };
  axios(conflg)
  .then(function(response){
    var back=response.data;
     document.getElementById('pusher').innerHTML=back.name;
     document.getElementById('message').innerHTML=back.message;
  })
  .catch(function(error)
  {
    console.error();
  });
}
function flush()
{
  showmessage();
}
submit
{
  
}