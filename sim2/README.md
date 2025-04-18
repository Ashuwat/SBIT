*true and false is buy and sell respectively*

The assumptions being made on this model are: 

- Every node either buys, sells, or holds
- There is max, one trade that happens at a single point in time
- Each node has a limit to their buy or sell
    Ex: Node wants to buy at a max price of 100
    Ex: Node wants to sell at a min price of 120
- There are 3 types of traders:
    - Random Trader - random - majority of the traders -> don't have many shares nor investment
    - The fundamentalist - using markov chains - some traders -> have some shares/investment
    - The big players - influence a lot of people  -> save a lot of shares/investment
    - The market makers - can tank markets -> have the most shares/investment
- Everyone that 
- Communication happens via a heirarchy:
    - big players -> the fundamentalists -> random traders
    - fundamentalists must have at least 2 big players in their group
    - the random traders must have at least 4 fundamentalists in their group
    - market makers are anonymous - make decisions for themselves
