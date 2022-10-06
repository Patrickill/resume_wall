
function showmessage()
{
  document.getElementById('comment_area').innerHTML="评论:";
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
        var message_comment=response.data.data.comment;
        for(let comment in message_comment)
        {
          alert(JSON.stringify(comment));
          let comment_ul=document.getElementById('comment_area');
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
  });
}
function submit()
{
  let comment_ul=document.getElementById('comment_area');
  let new_comment=document.createElement('li');
  var myDate = new Date();
  var my_data=myDate.toLocaleString( ); //获取日期与时间
 new_comment.innerHTML=document.getElementById('comment').value+"    "+my_data;
  comment_ul.appendChild(new_comment);
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
})
.catch(function(error)
{
    console.log(error);
    alert(error);
})
}
