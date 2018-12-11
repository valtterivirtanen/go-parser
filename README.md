# go-parser


PPPPPPPPPPPPPPPPP
P::::::::::::::::P
P::::::PPPPPP:::::P
PP:::::P     P:::::P
  P::::P     P:::::Paaaaaaaaaaaaa  rrrrr   rrrrrrrrr       ssssssssss       eeeeeeeeeeee    rrrrr   rrrrrrrrr
  P::::P     P:::::Pa::::::::::::a r::::rrr:::::::::r    ss::::::::::s    ee::::::::::::ee  r::::rrr:::::::::r
  P::::PPPPPP:::::P aaaaaaaaa:::::ar:::::::::::::::::r ss:::::::::::::s  e::::::eeeee:::::eer:::::::::::::::::r
  P:::::::::::::PP           a::::arr::::::rrrrr::::::rs::::::ssss:::::se::::::e     e:::::err::::::rrrrr::::::r
  P::::PPPPPPPPP      aaaaaaa:::::a r:::::r     r:::::r s:::::s  ssssss e:::::::eeeee::::::e r:::::r     r:::::r
  P::::P            aa::::::::::::a r:::::r     rrrrrrr   s::::::s      e:::::::::::::::::e  r:::::r     rrrrrrr
  P::::P           a::::aaaa::::::a r:::::r                  s::::::s   e::::::eeeeeeeeeee   r:::::r
  P::::P          a::::a    a:::::a r:::::r            ssssss   s:::::s e:::::::e            r:::::r
PP::::::PP        a::::a    a:::::a r:::::r            s:::::ssss::::::se::::::::e           r:::::r
P::::::::P        a:::::aaaa::::::a r:::::r            s::::::::::::::s  e::::::::eeeeeeee   r:::::r
P::::::::P         a::::::::::aa:::ar:::::r             s:::::::::::ss    ee:::::::::::::e   r:::::r
PPPPPPPPPP          aaaaaaaaaa  aaaarrrrrrr              sssssssssss        eeeeeeeeeeeeee   rrrrrrr

****************************************************************************************************

Author: Valtteri Virtanen
Mail: **
Student ID: **
Motive: Excercise for course TIETS02 Automaatit I 2018 - 2019, University of Tampere
****************************************************************************************************
Instructions for usage:

        1.      Provide propositional logic proposition. Allowed symbols are:
                        p _ 1 2 3 4 5 6 7 8 9 ( ) ~ ∼ & v - > <
        2.      Parser will perform lexical analysis on the proposition and turn
                it into sequence of tokens. Each token has type and value.
                Possible (token) values are:
                        proposition variable: p_i where i has to be positive integer
                        negation: ~
                        and-operator: &
                        or-operator: v
                        implication: ->
                        equivalence: <->
        3.      Parser will then ask user whether that wants only to know if the
                proposition is tautology, or assign each proposition variable a
                truth value and find out the proposition's truth value.
        4.      Parser then forms a parse tree to verify that proposition
                is solid. The rules are following:
                        Let PROPS be the smallest set such that
                        -For all p_i where i is positive integer, p_i ∈ PROPS
                        -If A ∈ PROPS, then ~A ∈ PROPS
                        -If A,B ∈ PROPS, then (A∘B) ∈ PROPS where ∘ ∈ {&,v,->,<->}
        5.      Parser then translates parse tree and comes out with a truth value
                which will be output to user.
        6.      Parser will run until user tells it to stop
        Example proposition could be something like:
        ( (p_1 & (p_1 -> p_3) ) -> p_3)
****************************************************************************************************