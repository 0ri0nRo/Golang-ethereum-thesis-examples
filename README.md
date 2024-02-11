# Guide to Running a Docker Image for Simulating Enclave

This guide provides detailed instructions on how to run a Docker image to simulate the enclave, as well as the necessary steps to install additional dependencies and set up the environment.

## Explanation of Steps

### Installation 

1. **Install the Docker image to simulate the enclave:**
    ```bash
    docker run -it -v $(pwd):/myapp -p 8080:8080 ghcr.io/edgelesssys/ego-dev
    ```
2. **Install Ethereum and Solc:**
    - [This link](https://stackoverflow.com/questions/47257800/error-when-installing-ethereum-package-ethereum-has-no-installation-candidate) contains detailed instructions on how to install Ethereum. Summarizing:
        ```bash
        sudo apt-get install software-properties-common
        sudo add-apt-repository -y ppa:ethereum/ethereum
        sudo apt-get update
        sudo apt-get install ethereum
        ```
    - Next, install `solc` to compile Solidity contracts.

## Running the Dockerfile from Terminal

Follow these steps if you want to execute the Dockerfile from the terminal:

1. Make sure you have Docker installed on your system. If you haven't installed it yet, you can follow the official instructions to [Install Docker](https://docs.docker.com/get-docker/).

2. Ensure you are in the directory of your project that contains the Dockerfile.

3. Use the `docker build` command to build the Docker image from the Dockerfile. You can assign a name to the image using the `-t` option. For example, if you want to name the image `myimage`, the command would be similar to this:

   ```
   docker build -t myimage .
   ```

   Make sure to include the dot `.` at the end of the command to indicate the current directory as the build context.

4. Once the image is built, you can run the Docker container using the `docker run` command. For example:

   ```
   docker run -it myimage
   ```

   Replace `myimage` with the name you gave to the image during the build phase.

By following these commands, you should be able to build and run the Docker container using the provided Dockerfile. Make sure you have sufficient privileges to execute Docker commands; you may need to use `sudo` if you are not a root user or do not have adequate permissions.



### Creating the Smart Contract

1. **Create the smart contract in Solidity:** 
    - Write your smart contract in Solidity and save it in a `.sol` file.

2. **Convert it with `abigen` and `solc` into a `.go` file:**
    - Use the `abigen` and `solc` tools to convert the Solidity file into a Go file. Make sure you have Go dependencies installed.

3. **Perform the following operations:**
    - `go mod init`
    - `go mod tidy`
    - `go get github.com/ethereum/go-ethereum`
    - `go get all`

These steps will install the necessary dependencies to run and test your Ethereum contract within the Go environment.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

[MIT](https://choosealicense.com/licenses/mit/)
