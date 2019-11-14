likes(bell,sports).
likes(mary,music).
likes(mary,sports).
likes(jane ,smith).

friend(john,X):-likes(X,reading),likes(X,music).
friend(john,X):-likes(X,sports),likes(X,music).