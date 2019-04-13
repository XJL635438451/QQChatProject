# QQChatProject
该项目是一个类似于QQ交流软件，分为client端和server端，目前只在后台的consel端实现一对多的通信交流，注册的账户存存入redis中。

要运行该程序：
1. 需要安装redis，如果redis安装在Linux上需要更改程序里面redis的初始化程序；
2. 先启动server端，然后可以启动多个client端；
3. 启动client端之后需要输入用户ID（数字）和用户密码，如果不存在则会在server端注册。
注意：注册之后程序不会退出，但是可能会和其他用户通信有问题，此时退出，重新登录即可。

后续工作：
1. 目前的代码只是简单实现，架构还需要改善，特别是处理大并发；
2. 只实现一对多通信，后续会实现点对点，包括QQ的加群，群聊等功能。
3. 将后台搬到前端来实现交流，有时间会考虑做APP。
