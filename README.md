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
3. Click the search icon. Search for any item.
4. Click on the listing to open up details for that item.
5. Check for a request to `GetWorldMarketSubList` in `Network` tab.
    - `cookie` is found in ` Request Headers` under `Cookie` -> `__RequestVerificationToken`.
   - `token` is found in `Form Data` under `__RequestVerificationToken`.

## Usage
1. Configure config.json
```
{
	"input": "filename with list of items to get prices for",
	"output": "filename of where to dump CSV data",
	"cookie": "__RequestVerificationToken obtained from session",
	"token": "__RequestVerificationToken obtained from making a sub list request",
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
id,name,grade,price,count
9213,Beer,1,1050,98870
9283,Cold Draft Beer,2,2030,34326
```
4. Profit
