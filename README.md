# known01

接口说明：  
1：分页获取信息列表  
URI:IP/v1/news  
请求方式:GET  
参数列表：  
参数名称 类型  含义   必传     标注  
page    int  页码    是     从0开始（一页多少条数据，服务端决定）  
返回值格式：  
{  
 "code": 0;   //非0表示异常    
 “msg”: "****" //接口状态信息，如果code为非0，该值为错误信息  
 "data":{  
    "list": [  
         {"id":1,"title":"****","flag":true,"publish":"2020-10-28","jump":"http://aaaa.com/*.html"},  //其中单条数据维度数据  
        ......
         {"id":2,"title":"****","flag":false,"publish":"2020-10-29","jump":"http://aaaa.com/*.html"},  //其中单条数据维度数据  
  ],  
   "end":false,      //表示当前页数据是否为最后一页。  
 }    
}     

2：根据id获取信息详情  
URI:IP/v1/detail  
请求方式:GET  
参数列表：  
参数名称  类型  含义  必传   标注  
id      int   消息id  是    为上述接口返回的负载数据中的id值。  
返回值格式：    
 {    
  "code": 0;   //非0表示异常  
  “msg”: "****" //接口状态信息，如果code为非0，该值为错误信息  
  "data": {"content":"****"},  //单条数据html内容，可能需要其他字段。  
 }  
 
 
 
 
 {"code":0,"data":[{"id":1,"title":"11月3日，解放军攻台？","status":1,"url":"https://dy.163.com/article/FN2UDNH905504DP0.html","is_real":2,"publish":"2020-11-09T15:37:40+08:00","author":"环球时报评论"},{"id":2,"title":"11月3日，解放军攻台？","status":1,"url":"https://dy.163.com/article/FN2UDNH905504DP0.html","is_real":2,"publish":"2020-11-09T15:39:45+08:00","author":"环球时报评论"},{"id":3,"title":"11月3日，解放军攻台？","status":1,"url":"https://dy.163.com/article/FN2UDNH905504DP0.html","is_real":2,"publish":"2020-11-09T15:41:03+08:00","author":"环球时报评论"},{"id":4,"title":"11月3日，解放军攻台？","status":1,"url":"https://dy.163.com/article/FN2UDNH905504DP0.html","is_real":2,"publish":"2020-11-09T15:41:51+08:00","author":"环球时报评论"},{"id":5,"title":"11月3日，解放军攻台？","status":1,"url":"https://dy.163.com/article/FN2UDNH905504DP0.html","is_real":2,"publish":"2020-11-09T15:42:43+08:00","author":"环球时报评论"}],"is_end":true}
 
 
3：用户提交查询  
 URI:IP/v1/brain  
 请求方式:GET  
 参数列表：  
 参数名称    类型     含义  必传   标注  
 content  string   提交的文本消息  是    用户编辑后的文本数据。  
 返回值格式：    
  {  
   "code": 0;   //非0表示异常  
   “msg”: "****" //接口状态信息，如果code为非0，该值为错误信息  
   "data": {"suggest":"****","flag":true},   
  }  
 
