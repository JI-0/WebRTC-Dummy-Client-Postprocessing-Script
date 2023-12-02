# Dummy Client Timestamp Postprocessor
This is a postprocessing script application for the timestamp output of the [WebRTC Dummy client testing suite](https://github.com/JI-0/WebRTC-Dummy-Client). This script takes the .csv file created with the dummy client application and creates two additional .csv files:
* A `Processed_` .csv file, which has dummy clients as columns (client 0 is A, 1 is B...) and timestamps of whatever the time set in the Dummy client application was as rows; the data in the cells is the number of packets received in a certain second.
* A `Sum_` .csv file, which has a timestamps as rows like the previous file, but only has 3 columns - a minimum, maximum, and average of packets received per client of the previous file.

## Instructions
To use this script:
1. Clone the repo.
2. Place the source data file inside the repo directory.
3. Change the `maxPeers` variable if the test max peers was set to a different number.
4. Run the script with `go run .`.