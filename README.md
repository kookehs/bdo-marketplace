# bdo-marketplace
💱 collecting data and executing trades through Black Desert Online's Central Market API

## Features
- Regional markets
- Dump to CSV
- Data from `market.blackdesertonline.com` API
  - 🛍️ Your bid listings
  - 🎒 Your Central Market inventory
  - 🔥 View hot listings
  - 📒 Search market listings
  - 💹 Get an item's order book and price history
- Trade actions
  - Buy items
  - Sell items
  - Collect funds
  - Cancel listings

## Obtaining credentials for API calls
You can use the developer tools provided by your browser to look at network requests for the site.

0. You may need to spoof `User-Agent` as a mobile browser. Not sure if they have removed the restriction on viewing marketplace from mobile only.
1. https://market.blackdesertonline.com/ or other regional URL
    - Select your region and sign in.
2. Open `Developer Tools` for your browser and select `Network`.
    - For Chrome it's `Ctrl+Shift+I` and click `Network` tab.
    - For Firefox it's `Ctrl+Shift+E` to open `Network` tab.
3. Click the search icon. Search for an item with enhancement levels.
4. Click on the listing to open up details for that item.
5. Check for a request to `GetItemSellBuyInfo` in `Network` tab.
    - `cookie` is found in `Request Headers` under `Cookie` -> `__RequestVerificationToken`.
    - `session` is found in `Request Headers` under `Cookie` -> `ASP.NET_SessionId`.
    - `token` is found in `Form Data` under `__RequestVerificationToken`.

## Usage
This repository can be used as a library. 

Executable included to simply fetch and dump data for a list of items to CSV.

1. Download the latest release [here](https://github.com/kookehs/bdo-marketplace/releases).
- Grab the `config.json` and the appropriate executable for your system.
- The `items.json` file is a list of some item IDs for reference. You can create your own with relevant IDs.
2. Configure config.json
```
{
	"input": "filename with list of items to get prices for",
	"output": "filename of where to dump CSV data",
	"cookie": "__RequestVerificationToken obtained from session",
	"token": "__RequestVerificationToken obtained from making a request",
	"agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36",
	"url": "regional URL"
}
```
3. Execute binary for your system
```
bdo-marketplace-darwin-arm64
bdo-marketplace-darwin-amd64
bdo-marketplace-linux-386
bdo-marketplace-linux-amd64
bdo-marketplace-windows-386.exe
bdo-marketplace-windows-amd64.exe
```
4. Upload CSV dump to your spreadsheet
```
id,name,grade,enhancement,maximum,minimum,price,count
719898,Fallen God's Armor,1,0,37700000000,32500000000,35100000000,0
719898,Desperate Fallen God's Armor,1,1,40500000000,34900000000,37700000000,0
719898,Distorted Fallen God's Armor,1,2,53500000000,46300000000,50000000000,1
719898,Silent Fallen God's Armor,1,3,65500000000,56500000000,61000000000,0
719898,Wailing Fallen God's Armor,1,4,122000000000,106000000000,114000000000,0
719898,Obliterating Fallen God's Armor,1,5,303000000000,261000000000,282000000000,0
7304,Strawberry,1,0,585,545,585,0
9065,Milk,0,0,10800,10000,10800,0
705511,Manos Ring,1,0,244000000,210000000,227000000,144
705511,Manos Ring,1,1,381000000,329000000,355000000,2
705511,Manos Ring,1,2,1130000000,985000000,1060000000,6
705511,Manos Ring,1,3,3830000000,3310000000,3570000000,5
705511,Manos Ring,1,4,17500000000,15100000000,16300000000,0
705511,Manos Ring,1,5,59000000000,55000000000,59000000000,0
```
5. Profit

## Roadmap to 0.1
- Support for multiple input and output files.
- Switch provided executable over to new library.
- Have provided executable dump price history.
- Proxy support for `centralmarket.Client`.
- Better logging of errors.
