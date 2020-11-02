# bdo-marketplace
💱 collecting data from Black Desert Online's marketplace

## Features
- Prices from `market.blackdesertonline.com` API
- Regional Markets
- Dump to CSV

## Obtaining credentials for API calls
You can use the developer tools provided by your browser to look at network requests for the site.

0. You may need to spoof `User-Agent` as a mobile browser. Not sure if they have removed the restriction on viewing marketplace from mobile only.
1. https://market.blackdesertonline.com/
    - Select your region and sign in.
2. Open `Developer Tools` for your browser and select `Network`.
    - For Chrome it's `Ctrl+Shift+I` and click `Network` tab.
    - For Firefox it's `Ctrl+Shift+E` to open `Network` tab.
3. Click the search icon. Search for an item with enhancement levels.
4. Click on the listing to open up details for that item.
5. Check for a request to `GetItemSellBuyInfo` in `Network` tab.
    - `cookie` is found in ` Request Headers` under `Cookie` -> `__RequestVerificationToken`.
    - `token` is found in `Form Data` under `__RequestVerificationToken`.

## Usage
1. Configure config.json
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
2. Execute binary for your system
```
bdo-marketplace-darwin-arm64
bdo-marketplace-darwin-amd64
bdo-marketplace-linux-x8632
bdo-marketplace-linux-amd64
bdo-marketplace-windows-x8632.exe
bdo-marketplace-windows-amd64.exe
```
3. Upload CSV dump to your spreadsheet
```
id,name,grade,enhancement,maximum,minimum,price,count
719898,Fallen God's Armor,1,0,30100000000,25900000000,28000000000,0
719898,Desperate Fallen God's Armor,1,1,30500000000,26300000000,28400000000,0
719898,Distorted Fallen God's Armor,1,2,32200000000,27800000000,30000000000,0
719898,Silent Fallen God's Armor,1,3,40500000000,34900000000,37700000000,0
719898,Wailing Fallen God's Armor,1,4,122000000000,106000000000,114000000000,0
719898,Obliterating Fallen God's Armor,1,5,303000000000,261000000000,282000000000,0
7304,Strawberry,1,0,585,545,585,0
```
4. Profit
