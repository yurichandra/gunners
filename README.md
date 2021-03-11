# Gunners

A livescore API utilize Twitter stream API.

## How to use

- Setup mongoDB in your local.
- Add your Twitter API token on env.
- Run server by type command `refresh run main.go`.
- Store match that you want to keep livescore up to date by hit `POST /matches` endpoint with example payload below.

```JSON
    {
        "homeTeam": "Liverpool",
        "awayTeam": "Everton",
        "date": "2021-02-21"
    }
```

- Just wait, if a team that you watch is score, it will print `Liverpool is scored` in the console.

## Main Services

### `TwitterService`

TwitterService will responsible to be layer that abstract Twitter API functionalities, especially for stream API. Functionalities in this service includes `SetRules`, `GetRules` and `Stream`. Set rules is for managing keywords, tweets from a specific twitter account or any rules to listen or stream tweets from Twitter API. In this project, this layer will store rules for matches hastags and specify account handler that are going to be source of tweets. `Stream` will be method that allowing this API to stream tweets from Twitter API based on rules that we specify on `SetRules` functionality. Stream will always opened and will utilize Golang concurrency to keep HTTP connection open to Twitter API.

Any goalscoring events, will returning messages from Twitter API stream. Data for next process will be sent via `TwitterChan` of Golang channel.

### `WebsocketService`

WebsocketService will responsible to be layer that sending data into websocket connection. This service, will range over `TwitterChan` channel. If there are messages coming from `TwitterChan` channel, service will proxying data into opened websocket connection, which is in this case is livescore web client.

### API Design

![Screenshot](livescore-api.png)

In case you guys have any feedbacks, please don't hesitate to raise issues. Thank you.
