function showmessage()
{
  var conflg={
    method:'get',
    url:'http://180.76.167.8:8888/api/message/getR',
    headers: {
      'Authorization': localStorage.getItem('token')
  }
}
  axios(conflg)
  .then(function(response){
      document.getElementById('from').innerHTML=response.data.data.info.name;
      document.getElementById('to_content').innerHTML=response.data.data.info.message;
  })
  .catch(function(error)
  {
    console.error();
  });
}