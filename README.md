# known01

known01.exe  -c D:\gowork\src\known01\conf\known01_windows.yml -log_dir=log

需要把配置文件+分词词表（dictionary.txt）和二进制 放到同一个目录



接口说明：

3：用户提交查询  
 URI:IP:PORT/message/brain 
 请求方式:POST
 参数列表：  
 参数名称    类型     含义        必传   标注  
 content  string   提交的文本消息  是    用户编辑后的文本数据。  
 sender    string   短信发送者号码 否    如95595 

请求参数示例：
{"sender":"","content":"光大银行信用卡促销性价，畅销十年，利率最低价，详情请登录http://girls.com/index.html,请咨询手机：13812348888."}
 返回值格式：    
 {"code":0,"data":{"flag":2,"hotline":"95595","score":0.14,"website":"www.cebbank.com"},"msg":"succ"}
