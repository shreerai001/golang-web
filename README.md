# Run sigma backend

## Clone the project 
    git clone https://github.com/Aavash/sigma.git
#### Install dependencies
    go mod vendor 
    or
    go mod build
####  Copy env sample as .env and configure your settings in .env file
    cp env.sample .env

####  Run
    go run main.go