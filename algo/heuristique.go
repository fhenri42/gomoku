package main

func lol() {
  return
}

/*
1) current alignments into account                                                                              Pois: 500
2) check whether an alignment has enough space to develop into a 5-in-a-row                                     Pois: 1000
3) heuristic weigh an alignment according to its freedom                                                        Pois: 900
4) potential captures into account                                                                              Pois: 800
5) heuristic take current captured stones into account                                                          Pois: 800
6) heuristic check for advanteageous combinations                                                               Pois: 700
7) heuristic take both players into account                                                                     Pois: 600
HARDCORE
8)heuristic take past player actions into account to identify patterns and weigh board states accordingl        Pois:
*/
