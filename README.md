# Etherium block analyzer
This parser is using etherscan.io to find the wallet that had its balance
changed the most for the selected amount of last blocks.
## Usage
1. Obtain the API Key on https://etherscan.io/
2. Build the Docker image
3. Launch the container with the following environment variables
### Variables
* APIKEY - etherscan.io API Key
* RATE - max requests per second
* BLOCKAMOUNT - the amount of last blocks that will be scanned