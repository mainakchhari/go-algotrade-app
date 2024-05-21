# go-algotrade-app

## Premise
Hi! Never mind the title. I know this is not algo trading! As a hobbyist/enthusiast in stock market movements, 
I was recently curious how a simple MA crossover strategy would perform as compared to a random decision based strategy. 
As an example I implemented a basic stream processing application. 
This is more of an exercise in software development principles than trading. 

**WARNING: Please do not apply this with real money (or do so at your own risk!)

## Introduction
We read data as events from a binance provided websocket [aggTrade streams](https://developers.binance.com/docs/derivatives/usds-margined-futures/websocket-market-streams/Aggregate-Trade-Streams). 
Two wallets are initialized (with virtual money) and each is subjected to a different strategy. One is a simple moving averages based
crossover strat that compares two moving averages (long and short) which are evaluated on each event and makes a BUY/SELL decision upon crossover 
([More here..](https://www.investopedia.com/articles/active-trading/052014/how-use-moving-average-buy-stocks.asp)) and decides to HOLD otherwise. The other is a random decision strategy 
which randomly chooses either of BUY/HOLD/SELL for each input event, agnostic of event parameters. 

The strategies are also agnostic of any wallet and make their decisions based on aggregated trade signals only. 
An executor connects a strategy to a wallet and executes an actual transaction on the wallet depending on the strategy output. 
For simplification, a fixed amount executor is implemented that decides to buy a fixed amount 
of an asset corresponding to a BUY/SELL signal. As output, we display the wallet balances (both asset holding and money) which
are refreshed with each event
