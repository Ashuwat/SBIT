#### Sim_1 is a simulation that shows segregation of nodes through emergence.
- The simulation consists of n number of nodes that are attracted to nodes that have have similar vectors (done by checking vector similarity)
- These nodes only follow 1 rule: move closer to the nodes that are more alike
-   There are variations in movement, if similarity is higher, then they move faster to them.
- The emergence patterns look like this:

These are the node positions before:
<img src="https://github.com/user-attachments/assets/085e682c-fdd1-4f37-8d97-285ada0ad110" width="300">
These are the positions after: 
<img src="https://github.com/user-attachments/assets/a476aa25-94d5-44fe-8adf-d2d366e4b11c" width="300">

The purpose of this project was to see how social behaviours create groups (or clusters) of various information and are biased towards that information. For instance when information was passed along, there were only select groups that took the information, and many rejected it. 
The implications of this model can be shown in financial markets, to understand what groups may be going long, or short in a specific market.

*This wasn't inspired by the schelling's model, however after I made this model I realized that it was very similar structure to that model, where the model shows similarities between how people behave.


#### Sim_2 is a simulation that shows equity market performance with stochasticity. 
This one is a lot cooler.

- The simulation consists of a number of traders that can either buy/sell/hold and that is given by a random uniform probability.
- Given this, there is an orderbook that takes in these orders as 'tickets'.
- These tickets are then handled by a clearing house that, first checks whether this trade is even possible, and then sees if there is a trade that works for it.
- Orders get filled, and the cycle continues.

There was a lot of analysis that happened here, it can be seen in the ```interpretation``` folder. 
