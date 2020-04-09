# Tank

<a name="9YDQQ"></a>
## requirement
1. 使用多线程编制一个程序，实现多线程任务的调度
2. 高阶要求：当线程比较多时如何管理性能；
3. 开发语言不限、程序形式不限；
4. 程序、运行结果+简要说明文档（运行环境、设计思路）


<a name="m1WuA"></a>
## design
本次程序设计选择多线程塔克大战为题，体验并发编程，了解进程，携程在程序运行中的调度。选择 JavaScript 进行编程，结合 html + css 进行游戏页面显示。


将程序分为两个线程，一个是坦克和背景的绘制，另一个是子弹的移动。在程序中使用两个画布进行渲染，主进程通过计时器调用线程，通过不断的画布渲染实现坦克大战游戏的实现。


1. 使用画布构建画出 Boss
1. 使用二维数组存储墙，草，水的位置（x，y），图片像素选择 40*40
1. 在画布中使用矩形构建英雄坦克和敌军坦克
1. 构建敌军坦克子弹数组，并让每一颗子弹运行
1. 构建坦克与建筑物碰撞判定函数
1. 构建子弹与坦克碰撞的判定函数
1. 构建爆炸显示函数
1. 构建胜利判定函数
1. 构建计时器定时刷新画布，实现页面动态显示



上诉的步骤是实现坦克大战的思路，在并发的设计时采用非抢占式方式，使用 FCFS 算法进行坦克的移动和子弹的移动,使用相同时间间隔的计时器同时刷新画布，实现页面动态显示。


<a name="2h930"></a>
## summary
测试环境：

- 操作系统：macOS Catalina 10.15.2
- 开发软件：Visual Studio Code 1.43
- Web浏览器：Apple Safari



直接使用 Web 浏览器打开 tank.html 文件即可运行程序，运行结果如下图所示
![tank1.jpg](https://cdn.nlark.com/yuque/0/2020/jpeg/532901/1586437097498-84b3e306-b8a2-417f-a622-0c05860c2f27.jpeg#align=left&display=inline&height=624&name=tank1.jpg&originHeight=624&originWidth=780&size=110887&status=done&style=none&width=780)![tank2.jpg](https://cdn.nlark.com/yuque/0/2020/jpeg/532901/1586437103997-31ad5807-f62d-4a5b-a463-126b95562770.jpeg#align=left&display=inline&height=624&name=tank2.jpg&originHeight=624&originWidth=780&size=102699&status=done&style=none&width=780)


通过本次作业，关于并发编程有了更深的理解，回想自己在编程过程中遇到的各种问题之后感受颇多，而且在完成之后发现还有很多可以提高性能的地方。第一点是可以将每个坦克的移动和子弹的移动创建新的协程，原来的方案是将坦克和子弹放入数组中遍历运行，将线程转换为协程可以更高的提高系统的性能。其次是关于协程调度算法的问题，由于之前的并发编程当中，从未考虑过调度算法的问题，如果能够选择适当的调度算法同样可以提高并发系统的性能。最后是关于代码复用的问题，其实
画墙，草，水的方法基本一致，应该将其抽象出来，使得代码更加工程化。

项目地址：https://github.com/miuer/ncepu-work/tree/master/architecture/tank

