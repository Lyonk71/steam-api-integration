# Steam Backend Service

The purpose of this repository is to demonstrate how to make the necessary backend calls to Steam to support things like microtransactions/purchases.

To be able to run this locally you need to:

* Generate a SteamWorks API key
* Have a valid Steam App ID provided by Steamworks (this is the id that represents your game)

## Setting up and running locally

### Install go

First you need to install golang: https://go.dev/doc/install

### Download dependencies
`make deps`

### Create config
Create a new `config.env` file that matches `config.env.sample` and replace the values as necessary.

### Run the service
`make run`

### Hit the service
#### Initiate a microtransaction
```
curl -X POST \
  http://localhost:8090/v1/api/microtxn/initTxn \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/x-www-form-urlencoded' \
  -d 'SteamAccountID=76561198006253851&OrderID=100002&ItemID=1'
```

#### List all friends for a given user
```
curl "http://localhost:8090/v1/api/steamuser/getfriendlist?steamid=76561198351398889&relationship=friend" 
```

# Steam Microtransaction Flow 

![Steam Microtransaction Flow ](https://github.com/Rushdown-Studios/steam-api-integration/blob/main/docs/steam-microtxn-flow.png?raw=true)
