# CoinTracker  Assignment
Sepandar Sepehr

## Prerequisite
Make sure you have the latest version of Go installed
## Running the App
First run `go build cointracker_app.go` to make sure dependencies are installed

Then run `go run cointracker_app.go` 

Go to `http://localhost:8080/wallet` on your browser.
On this page, you can add a wallet ID, remove one, get transactions of, and synchronize transactions for all stored wallets. 

## TODO
* On remove wallet it doesn't clean transactions. We need to do that to avoid the problem of adding that wallet again and syncing it which causes storing transactions to fail.
* Build repository instead of interacting with DB directly
* Separate the struct of client's transaction response from DB models
* Add logging and metrics
* Pagination for API call to get transactions. Right now you have to manually call sync to fetch next pages
* Build cron job to kick off syncing automatically