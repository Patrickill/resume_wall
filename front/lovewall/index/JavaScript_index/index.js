window.onload=showmessage();
function $(id) {
  return typeof id === "string" ? document.getElementById(id) : null;
}
function showmessage()
{
  document.getElementById('comment_area').innerHTML="";
  var conflg={
    method:'get',
    url:'http://180.76.167.8:8888/api/message/getR',
    headers: {
      'Authorization': sessionStorage.getItem('token')
  }
}
  axios(conflg)
  .then(function(response){
      console.log(response.data);
      document.getElementById('from').innerHTML=response.data.data.info.name;
      document.getElementById('to_content').innerHTML=response.data.data.info.message;
      document.getElementById('show_time').innerHTML=response.data.data.info.time;
      comment_ul.appendChild(new_comment);
       var message_comment=response.data.comment;
       for(comment in message_comment)
       {
         let comment_ul=getElementById('comment_area');
         let new_comment=document.createElement('li');
         new_comment.innerHTML=comment.message+"    "+comment.time;
         comment_ul.appendChild(new_comment);
      }
      sessionStorage.setItem('comments',response.data.data.comments);
      sessionStorage.setItem('message_id',response.data.data.id);
      })
  .catch(function(error)
  {
    console.error();
    alert(error);
  });
}
function send_comment()
{
  var data=JSON.stringify(
    {
        "message_id":sessionStorage.getItem('message_id'),
        "message": document.getElementById('comment').value,
    }
);
var config={
    method: 'POST',
    url: 'http://180.76.167.8:8888/api/comment/add',
    data: data,
    headers: {
        'Authorization': sessionStorage.getItem('token')
    }
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
