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
bdo-marketplace-darwin32
bdo-marketplace-darwin64
bdo-marketplace-linux32
bdo-marketplace-linux64
bdo-marketplace-windows32.exe
bdo-marketplace-windows64.exe
```
3. Upload CSV dump to your spreadsheet
```
id,name,grade,enhancement,maximum,minimum,price,count
13414,Yuria Crescent Pendulum,1,0,103000,89500,96500,82
13414,Yuria Crescent Pendulum,1,8,1970000,1710000,1840000,5
13414,Yuria Crescent Pendulum,1,11,2710000,2350000,2530000,0
13414,Yuria Crescent Pendulum,1,13,5600000,4860000,5250000,0
13414,Yuria Crescent Pendulum,1,16,27900000,24100000,26000000,0
13414,Yuria Crescent Pendulum,1,17,41600000,35800000,38700000,0
13414,Yuria Crescent Pendulum,1,18,100000000,86500000,93500000,0
13414,Yuria Crescent Pendulum,1,19,354000000,306000000,330000000,0
13414,Yuria Crescent Pendulum,1,20,795000000,685000000,740000000,0
```
4. Profit
