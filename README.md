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
3. ### Building the Docker Image

To build the Docker image, navigate to the project directory containing the Dockerfile and run the following command:

Replace `myimage` with the name of the Docker image you specified during the build phase.

## Dependencies

The Dockerfile installs the following dependencies:

- `software-properties-common`
- `ethereum`

## Guide to Running a Docker Image for Simulating Enclave

This guide provides detailed instructions on how to run a Docker image to simulate the enclave, as well as the necessary steps to install additional dependencies and set up the environment.

### Explanation of Steps

#### Installation

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
