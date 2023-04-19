# Goroutine

## 并发和并行

## 特点介绍

* goroutine是由Go的运行时（runtime）调度和管理的，在语言层面内置了调度和上下文切换机制
* goroutine奉行的是，通过通信来共享内存，而不是共享内存来通信
* 在Go语言编程中你不需要去自己写进程、线程、协程，你的技能包里只有一个技能–goroutine

## 用法

```
go 函数（）
```

* 这样就会去启动一个goroutine，去执行这个函数
* 程序启动时，会为main()函数创建一个goroutine，当main（）函数默认的goroutine结束的时候，在main函数里面的goroutine会一同结束
