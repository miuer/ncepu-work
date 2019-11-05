/*
 * Revision History:
 *     Initial:                 2019/11/05      Miuer
 *     Test environment:        Swi-prolog 7.6.4
 */

% You can customize the directed graph.
road(a, b).
road(a, e).
road(b, c).
road(b, d).
road(c, g).
road(d, c).
road(e, d).
road(e, f).
road(f, g).

% One-dimensional space solution.
path(Y, Z) :- road(Y, Z), write(Y), write(" --> "), write(Z).

% Dimension upgrades
path(X, Z) :- road(Y, Z), path(X, Y), write(" --> "), write(Z).

goal(X, Z) :- path(X, Z).