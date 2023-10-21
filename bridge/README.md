# BRIDGE
In this directory is the modified code of ThorNode bifrost to establish a connection between Osmosis and ThorChain. The code to connect BTC and Gaia is also included. For now, I have conducted tests on various functionalities related to accounts, transactions, and cryptographic operations within the context of the Osmosis package and associated ThorChain processes.

## Osmosis chain client

This code defines an Osmosis client to interact with a Cosmos-based blockchain. The code is essentially designed to facilitate interactions with a Cosmos-based blockchain. The OsmosisClient provides the necessary methods to start and stop the client, sign transactions, fetch account details, etc. 

### TEST
 I have create some tests to test functionalities related to the osmosis package. Here are the main points:

Initial Setup:
A suite named CosmosTestSuite is set up, which includes fields like a Thorchain bridge, metrics, and keys.
Within the suite, a method SetUpSuite is defined, which sets up configurations, initializes a Thorchain bridge, and keys.
The suite also contains a TearDownSuite method to clean up after the tests, including unsetting environment variables and removing directories.
GetMetricForTest: A method that provides metrics for testing. If metrics aren't initialized, it creates a new one with a certain configuration.

TestGetAddress:
This test checks the function's ability to retrieve an account based on an address or public key.
The test uses mock bank and account service clients to simulate the behavior.
Various assertions ensure that the retrieved account details match the expected values.

TestProcessOutboundTx:
This test is focused on processing an outbound transaction.
It sets up a mock HTTP server and creates a new Osmosis client.
The test then simulates processing an outbound transaction and checks various transaction properties, ensuring they match expected values.

TestSign:
This test revolves around signing a message.
It involves converting cryptographic keys, creating a local key manager, setting up mock services, and processing an outbound transaction.
The main focus is to ensure the construction of an unsigned transaction, verify its properties, and then check the signature post-signing.

Overall, the code tests various functionalities related to accounts, transactions, and cryptographic operations within the context of the osmosis package and associated Thorchain processes.

Results:
OOPS: 7 passed, 2 FAILED