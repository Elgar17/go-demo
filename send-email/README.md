# Go 语言发送邮件

发送邮件一般至少需要一下信息：

- 发送人：发送邮件的人，唯一的
- 收件人：发送个邮箱的人，一个或多个
- 邮件的主题：字符串
- 邮件内容： 可以是文字，HTML 等内容。
- 发送人的秘钥：秘钥相当于密码，这里以 qq 邮箱为例，登录后“设置” => “账户” => “生成授权码” 获取

这里使用到 "gopkg.in/gomail.v2" 这个包发送邮件。
