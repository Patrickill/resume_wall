
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
    console.log(response.data);
      document.getElementById('from').innerHTML=response.data.data.info.name;
      document.getElementById('to_content').innerHTML=response.data.data.info.message;
      document.getElementById('show_time').innerHTML=response.data.data.info.time;
      localStorage.setItem('message_id','example');
      localStorage.removeItem('message_id');
      localStorage.setItem('message_id',response.data.data.info.id);
  })
  .catch(function(error)
  {
    console.error();
  });
}
function clean()
{
  getElementById('comment').innerHTML="nihao";
}