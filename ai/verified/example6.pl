p(a).
p(b).
q(b).

r1(X):-p(X),q(X).
r1(c).

r2(X):-!,p(X),q(X).
r2(c).

r3(X):-p(X),!,q(X).
r3(c).

r4(X):-p(X),q(X),!.
r4(c).