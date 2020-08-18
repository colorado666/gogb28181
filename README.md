#### Usage

运行程序，进入交互试shell，输入要打开的子应用：

有如下几个子应用：
传输层：tcps、tcpc、udps、udpc
TU层：uas 和 uac 

```
➜  gogb28181 go run main.go
gb28181 test
>>> help

Commands:
  clear      clear the screen
  exit       exit the program
  help       display help
  tcpc       tcp client
  tcps       tcp server
  uac        sip user agent client
  uas        sip user agent server
  udpc       udp client
  udps       udp server

```



#### 说明1

对比了很多sip库的实现，有的与标准sip差距太大，有的实现的太复杂。osip2比较容易懂。所以这个代码参考（抄袭）的osip2 和 exosip2的相关实现。比如状态机的对应管理连函数名都没改。


#### 说明2

这是一个gb28181的golang实现。仅实现了部分基础库。


1、sip标准兼容性

sip协议的实现github上其实有好几个go库了，但都实现的过于复杂了。所以我这边重新造了个轮子，毕竟go造轮子太简单了。

```
Implementations MUST be able to process multiple header field rows with the same name in any combination of the 
single-value-per-line or comma-separated value forms.
```
上面是RFC3261的内容。
因为gb28181对于sip的使用也是很初级的（我的理解），并且后续扩展的话也很方便。
我们这个sip包里面并没有完全实现RFC，很多header字段没有支持、且header字段不处理多值的情况。

