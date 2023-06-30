# TrussTakeHome

This repo aims to solve the CSV Normalization problem.

#  How to run locally 

For this takehome, I chose Golang to create the application. For my own testing purposes(using a Windows based machine), I have included a set of docker commands which requires the docker desktop app is installed and the latest version of Go(1.20.5) on your local machine. Otherwise, the application should just require the latest version of Go(1.20.5) to run the below commands on linux based machines: 

1. To download external dependencies, run `make deps`

3. To run the service locally, run `make run INPUT={input_file_path} OUTPUT={output_file_path}`. If no input file or output file are provided, the program will default to using sample.csv in the test data directory and will output to output.csv. Input path should be relative to src folder.

4. To run unit tests locally, run `make test`

# References

https://docs.docker.com/engine/install/
https://go.dev/dl/

# Additional features I would've liked to implement 
2. More thorough testing
3. Telemetry
4. A makefile that's portable 