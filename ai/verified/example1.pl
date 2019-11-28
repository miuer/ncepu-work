love(nizhen,wangfei).
love(nizhen,zhouhuimin).
love(wangfei,liyapeng).
love(zhouhuimin,nizhen).
love(liyapeng,wangfei).
love(liyapeng,zhouhuimin).
love(liudehua,zhouhuimin).

lovers(X,Y):-love(X,Y),love(Y,X).

rival_in_love(X,Y):-love(X,Z),not(love(Z,X)),love(Z,Y).
