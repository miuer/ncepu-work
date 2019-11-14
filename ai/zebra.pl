/*
 * Revision History:
 *    Initial:                 2019/11/14    Miuer
 *    Test environment:        Swi-prolog 7.6.4
 */

member(X,[X|T]).
member(X,[_|T]):-
      member(X,T).

color(h(C,N,P,Y,D),C). 
nation(h(C,N,P,Y,D),N). 
pet(h(C,N,P,Y,D),P).
yan(h(C,N,P,Y,D),Y).
drink(h(C,N,P,Y,D),D). 


next(A,B,[A,B,C,D,E]).
next(B,C,[A,B,C,D,E]). 
next(C,D,[A,B,C,D,E]). 
next(D,E,[A,B,C,D,E]). 
next(B,A,[A,B,C,D,E]). 
next(C,B,[A,B,C,D,E]).
next(D,C,[A,B,C,D,E]). 
next(E,D,[A,B,C,D,E]). 

right(A,B,[A,B,_,_,_]).
right(A,B,[_,A,B,_,_]).
right(A,B,[_,_,A,B,_]).
right(A,B,[_,_,_,A,B]).


middle(X,[_,_,X,_,_]). 
first(A,[A|X]). 


solve(X,TT,TTT):-

X=[h(C1,N1,P1,Y1,D1),h(C2,N2,P2,Y2,D2),h(C3,N3,P3,Y3,D3),h(C4,N4,P4,Y4,D4),h(C5,N5,P5,Y5,D5)],

    member(Z1,X), 
    color(Z1,red), 
    nation(Z1,englishman),

    member(Z2,X),
    pet(Z2,dog),
    nation(Z2,spaniard),

    member(Z3,X),
    first(Z3,X),
    nation(Z3,norwegian),

    member(Z4,X),
    yan(Z4,kools),
    color(Z4,yellow),

    member(Z5,X),
    pet(Z5,fox),%交叉处用变量A……I表示
    next(Z6,Z5,X),
    yan(Z6,chesterfields),    

    member(Z7,X),
    color(Z7,blue),
    next(Z8,Z7,X),
    nation(Z8,norwegian),

    member(Z9,X),
    yan(Z9,winston),
    pet(Z9,snails),

    member(Z10,X),
    drink(Z10,orangeJuice),
    yan(Z10,luckyStrike),

    member(Z11,X),
    nation(Z11,ukrainian),
    drink(Z11,tea),

    member(Z12,X),
    nation(Z12,japanese),
    yan(Z12,parliaments),

    member(Z13,X),
    pet(Z13,horse),
    next(Z14,Z13,X),
    yan(Z14,kools),
    
    member(Z15,X),
    color(Z15,green),
    drink(Z15,coffee),

    member(Z16,X),
    color(Z16,ivory),
    color(Z17,green),
    right(Z16,Z17,X),

    middle(Z18,X),
    drink(Z18,milk),

    member(TT,X),
    pet(TT,zebra),

    member(TTT,X),
    drink(TTT,water).