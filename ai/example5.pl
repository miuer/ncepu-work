data(one). 
data(two).
data(three).

cut_test_a(X) :- data(X). 
cut_test_a('last clause').
go1:-cut_test_a(X), write(X), nl, fail.
