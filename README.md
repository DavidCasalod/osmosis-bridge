# Osmosis-Thornode
In this GitHub repository, there are 2 approaches. Within the local net folder, you will find the necessary code to create and launch a mock network, allowing you to work in a development environment. Inside the bridge folder is the modified bifrost code aiming to achieve communication between BTC and Osmosis, For now, I've focused on establishing the connection between ThorNode and Osmosis, with the intention of later trying to establish the Osmosis to Bitcoin connection. Inside the 'bridge' folder, there is a Readme.md with further explanation.
It also integrates with gaia since it's the only one within cosmos that, so far, allows working with ATOM to BTC and vice versa.



Below is an explanation of how to work with the local net to make different transactions.

## Local-net

First and foremost, you will need to install dependencies; you may skip packages you already have. To do this, follow the instructions from the repository at https://gitlab.com/thorchain/thornode.git.

Once we have completed the installation, we proceed:
```sh
make run-mocknet
```
Results: 
```sh
david.casalod@bip-dev-04:~/nitka/osmosis-thornode$ make run-mocknet
[+] Running 16/16
 ✔ Network docker_default            Created                                0.2s 
 ✔ Volume "docker_bifrost"           Created                                0.0s 
 ✔ Volume "docker_thornode"          Created                                0.0s 
 ✔ Container docker-ethereum-1       Start...                               0.2s 
 ✔ Container docker-binance-1        Starte...                              0.2s 
 ✔ Container docker-midgard-db-1     Sta...                                 0.2s 
 ✔ Container docker-gaia-1           Started                                0.2s 
 ✔ Container docker-binance-smart-1  Started                                0.2s 
 ✔ Container docker-bitcoin-cash-1   S...                                   0.2s 
 ✔ Container docker-bitcoin-1        Starte...                              0.2s 
 ✔ Container docker-avalanche-1      Star...                                0.2s 
 ✔ Container docker-litecoin-1       Start...                               0.2s 
 ✔ Container docker-dogecoin-1       Start...                               0.2s 
 ✔ Container docker-thornode-1       Start...                               0.0s 
 ✔ Container docker-bifrost-1        Starte...                              0.0s 
 ✔ Container docker-midgard-1        Starte...                              0.0s 
```
Now, let’s check pools - expect array to be empty 
```sh
curl localhost:1317/thorchain/pools
```
Result:
```sh
david.casalod@bip-dev-04:~/nitka/osmosis-thornode$ curl localhost:1317/thorchain/pools
[]
```
Boostrap vaults and pools - this will populate vaults, pools for testing
```sh
make bootstrap-mocknet
```
Result:
```sh
david.casalod@bip-dev-04:~/nitka/osmosis-thornode$ make bootstrap-mocknet
I[2023-10-21 07:27:06,316]  2     MASTER => PROVIDER-1 [SEED] 10.00000000 BNB.BNB, 800.00000000 BNB.LOK-3C0
I[2023-10-21 07:27:06,339]  5     MASTER => PROVIDER-1 [SEED] 5.00000000 BTC.BTC
I[2023-10-21 07:27:06,620]  7     MASTER => PROVIDER-1 [SEED] 2,000.00000000 DOGE.DOGE
I[2023-10-21 07:27:06,728]  9     MASTER => PROVIDER-1 [SEED] 500,000.00000000 GAIA.ATOM
I[2023-10-21 07:27:07,865] 12     MASTER => PROVIDER-1 [SEED] 2.00000000 BCH.BCH
I[2023-10-21 07:27:07,991] 15     MASTER => PROVIDER-1 [SEED] 2.00000000 LTC.LTC
I[2023-10-21 07:27:08,330] 18     MASTER => PROVIDER-1 [SEED] 400,000,000,000.00000000 ETH.ETH
I[2023-10-21 07:27:10,101] 20     MASTER => PROVIDER-1 [SEED] 4,000,000,000,000.00000000 ETH.TKN-0X52C84043CD9C865236F11D9FC9F56AA003C1F922
I[2023-10-21 07:27:15,167] 22 PROVIDER-1 => VAULT      [ADD:BNB.BNB:PROVIDER-1] 1,000.00000000 THOR.RUNE
I[2023-10-21 07:27:18,388] 23 PROVIDER-1 => VAULT      [ADD:BNB.BNB:PROVIDER-1] 2.50000000 BNB.BNB
I[2023-10-21 07:27:23,233] 24 PROVIDER-1 => VAULT      [ADD:DOGE.DOGE:PROVIDER-1] 10.00000000 THOR.RUNE
I[2023-10-21 07:27:28,477] 25 PROVIDER-1 => VAULT      [ADD:DOGE.DOGE:PROVIDER-1] 1,500.00000000 DOGE.DOGE
I[2023-10-21 07:27:33,356] 26 PROVIDER-1 => VAULT      [ADD:GAIA.ATOM:PROVIDER-1] 10.00000000 THOR.RUNE
I[2023-10-21 07:27:38,557] 27 PROVIDER-1 => VAULT      [ADD:GAIA.ATOM:PROVIDER-1] 1,500.00000000 GAIA.ATOM
I[2023-10-21 07:27:58,696] 33 PROVIDER-1 => VAULT      [ADD:BTC.BTC:PROVIDER-1] 1,000.00000000 THOR.RUNE
I[2023-10-21 07:28:03,756] 34 PROVIDER-1 => VAULT      [ADD:BTC.BTC:PROVIDER-1] 2.50000000 BTC.BTC
I[2023-10-21 07:28:13,588] 41 PROVIDER-1 => VAULT      [ADD:BCH.BCH:PROVIDER-1] 500.00000000 THOR.RUNE
I[2023-10-21 07:28:18,839] 42 PROVIDER-1 => VAULT      [ADD:BCH.BCH:PROVIDER-1] 1.50000000 BCH.BCH
I[2023-10-21 07:28:28,642] 43 PROVIDER-1 => VAULT      [ADD:LTC.LTC:PROVIDER-1] 500.00000000 THOR.RUNE
I[2023-10-21 07:28:33,931] 44 PROVIDER-1 => VAULT      [ADD:LTC.LTC:PROVIDER-1] 1.50000000 LTC.LTC
I[2023-10-21 07:28:38,867] 45 PROVIDER-1 => VAULT      [ADD:ETH.ETH:PROVIDER-1] 500.00000000 THOR.RUNE
I[2023-10-21 07:28:44,011] 46 PROVIDER-1 => VAULT      [ADD:ETH.ETH:PROVIDER-1] 4,000,000,000.00000000 ETH.ETH
I[2023-10-21 07:29:03,936] 47 PROVIDER-1 => VAULT      [ADD:ETH.TKN-0X52C84043CD9C865236F11D9FC9F56AA003C1F922:PROVIDER-1] 500.00000000 THOR.RUNE
I[2023-10-21 07:29:09,188] 48 PROVIDER-1 => VAULT      [ADD:ETH.TKN-0X52C84043CD9C865236F11D9FC9F56AA003C1F922:PROVIDER-1] 40,000,000,000.00000000 ETH.TKN-0X52C84043CD9C865236F11D9FC9F56AA003C1F922
I[2023-10-21 07:29:34,350] 49 PROVIDER-1 => VAULT      [ADD:BNB.LOK-3C0:PROVIDER-1] 400.00000000 BNB.LOK-3C0
I[2023-10-21 07:29:39,239] 50 PROVIDER-1 => VAULT      [ADD:BNB.LOK-3C0:PROVIDER-1] 500.00000000 THOR.RUNE
I[2023-10-21 07:29:44,423] 52 PROVIDER-1 => VAULT      [ADD:] 0.10000000 BNB.BNB
[]
```

Check for pools - for example: 
```sh
curl localhost:1317/thorchain/pool/GAIA.ATOM
```
Result:
```sh
david.casalod@bip-dev-04:~/nitka/osmosis-thornode$ curl localhost:1317/thorchain/pool/GAIA.ATOM
{
  "asset": "GAIA.ATOM",
  "short_code": "g",
  "status": "Available",
  "decimals": 6,
  "pending_inbound_asset": "0",
  "pending_inbound_rune": "0",
  "balance_asset": "150000000000",
  "balance_rune": "1000000000",
  "pool_units": "1000000000",
  "LP_units": "1000000000",
  "synth_units": "0",
  "synth_supply": "0",
  "savers_depth": "0",
  "savers_units": "0",
  "synth_mint_paused": false,
  "synth_supply_remaining": "105000000000",
  "loan_collateral": "0",
  "loan_cr": "0",
  "derived_depth_bps": "0"
}
```

Get thornode cli so we can interact with mocknet:
```sh
make cli-mocknet
```
Result:
```sh
david.casalod@bip-dev-04:~/nitka/osmosis-thornode$ make cli-mocknet
>>> THORNode Mocknet CLI <<<
1. The passphrase for all keys is "password"
2. The key named "dog" is an admin MIMIR key for the mocknet cluster
root@e688368a1b44:~# 
```
There are various accounts that you can use to interact with mocknet.
You can find it here and the mnemoni for them: https://gitlab.com/thorchain/thornode/-/blob/develop/build/docker/README.md#keys.

You can also add a new account using: 
```sh
thornode keys add 
```
Check for accounts 
```sh
thornode keys list
```
Result:
```sh
root@e688368a1b44:~# thornode keys list
Enter keyring passphrase:
- name: dog
  type: local
  address: tthor1zf3gsk7edzwl9syyefvfhle37cjtql35h6k85m
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"AmF4AUTWZEUSBtgqiR5n2Lgic/Yrr1mWupMo5TAubNRO"}'
  mnemonic: ""
- name: dvd
  type: local
  address: tthor10qtcngxapfhf3jcqvxzl564et0w63xwlucq90d
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"AzcmtsInkEuu6T2ZLgJAUxjAzfX9iDrgN1SaljIZbtBJ"}'
  mnemonic: ""
- name: fox
  type: local
  address: tthor13wrmhnh2qe98rjse30pl7u6jxszjjwl4f6yycr
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"Aw/2MBvAhnLEifCInxlpTjCXxV0I/nE8pI5jNI+Zblx6"}'
  mnemonic: ""
```

Now we have a local net, funded pools and three accounts.

### Testing
Fox user swaps rune for btc:
```sh
thornode tx thorchain deposit 1000000000 rune swap:btc/btc --from fox $TX_FLAGS
```
Result: 
```sh
root@e688368a1b44:~# thornode tx thorchain deposit 1000000000 rune swap:btc/btc --from fox $TX_FLAGS
Enter keyring passphrase:
{"body":{"messages":[{"@type":"/types.MsgDeposit","coins":[{"asset":"THOR.RUNE","amount":"1000000000","decimals":"0"}],"memo":"swap:btc/btc","signer":"tthor13wrmhnh2qe98rjse30pl7u6jxszjjwl4f6yycr"}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[],"fee":{"amount":[],"gas_limit":"200000","payer":"","granter":""}},"signatures":[]}

confirm transaction before signing and broadcasting [y/N]: y
{"height":"351","txhash":"DDCA5371A3A495183E44179A0F2FB45CD1555D898583C5CC41353D27E20B1E66","codespace":"","code":0,"data":"0A130A112F74797065732E4D73674465706F736974","raw_log":"[{\"events\":[{\"type\":\"coin_received\",\"attributes\":[{\"key\":\"receiver\",\"value\":\"tthor1g98cy3n9mmjrpn0sxmn63lztelera37nrytwp2\"},{\"key\":\"amount\",\"value\":\"1000000000rune\"}]},{\"type\":\"coin_spent\",\"attributes\":[{\"key\":\"spender\",\"value\":\"tthor13wrmhnh2qe98rjse30pl7u6jxszjjwl4f6yycr\"},{\"key\":\"amount\",\"value\":\"1000000000rune\"}]},{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"deposit\"},{\"key\":\"sender\",\"value\":\"tthor13wrmhnh2qe98rjse30pl7u6jxszjjwl4f6yycr\"}]},{\"type\":\"transfer\",\"attributes\":[{\"key\":\"recipient\",\"value\":\"tthor1g98cy3n9mmjrpn0sxmn63lztelera37nrytwp2\"},{\"key\":\"sender\",\"value\":\"tthor13wrmhnh2qe98rjse30pl7u6jxszjjwl4f6yycr\"},{\"key\":\"amount\",\"value\":\"1000000000rune\"}]}]}]","logs":[{"msg_index":0,"log":"","events":[{"type":"coin_received","attributes":[{"key":"receiver","value":"tthor1g98cy3n9mmjrpn0sxmn63lztelera37nrytwp2"},{"key":"amount","value":"1000000000rune"}]},{"type":"coin_spent","attributes":[{"key":"spender","value":"tthor13wrmhnh2qe98rjse30pl7u6jxszjjwl4f6yycr"},{"key":"amount","value":"1000000000rune"}]},{"type":"message","attributes":[{"key":"action","value":"deposit"},{"key":"sender","value":"tthor13wrmhnh2qe98rjse30pl7u6jxszjjwl4f6yycr"}]},{"type":"transfer","attributes":[{"key":"recipient","value":"tthor1g98cy3n9mmjrpn0sxmn63lztelera37nrytwp2"},{"key":"sender","value":"tthor13wrmhnh2qe98rjse30pl7u6jxszjjwl4f6yycr"},{"key":"amount","value":"1000000000rune"}]}]}],"info":"","gas_wanted":"0","gas_used":"156763","tx":null,"timestamp":"","events":[{"type":"tx","attributes":[{"key":"ZmVl","value":"","index":true}]},{"type":"coin_spent","attributes":[{"key":"c3BlbmRlcg==","value":"dHRob3IxM3dybWhuaDJxZTk4cmpzZTMwcGw3dTZqeHN6amp3bDRmNnl5Y3I=","index":true},{"key":"YW1vdW50","value":"MjAwMDAwMHJ1bmU=","index":true}]},{"type":"coin_received","attributes":[{"key":"cmVjZWl2ZXI=","value":"dHRob3IxZGhleWNkZXZxMzlxbGt4czJhNnd1dXp5bjRhcXhodmUzaGhtbHc=","index":true},{"key":"YW1vdW50","value":"MjAwMDAwMHJ1bmU=","index":true}]},{"type":"transfer","attributes":[{"key":"cmVjaXBpZW50","value":"dHRob3IxZGhleWNkZXZxMzlxbGt4czJhNnd1dXp5bjRhcXhodmUzaGhtbHc=","index":true},{"key":"c2VuZGVy","value":"dHRob3IxM3dybWhuaDJxZTk4cmpzZTMwcGw3dTZqeHN6amp3bDRmNnl5Y3I=","index":true},{"key":"YW1vdW50","value":"MjAwMDAwMHJ1bmU=","index":true}]},{"type":"message","attributes":[{"key":"c2VuZGVy","value":"dHRob3IxM3dybWhuaDJxZTk4cmpzZTMwcGw3dTZqeHN6amp3bDRmNnl5Y3I=","index":true}]},{"type":"tx","attributes":[{"key":"YWNjX3NlcQ==","value":"dHRob3IxM3dybWhuaDJxZTk4cmpzZTMwcGw3dTZqeHN6amp3bDRmNnl5Y3IvMA==","index":true}]},{"type":"tx","attributes":[{"key":"c2lnbmF0dXJl","value":"R0ZueVUydEViN2NQSmlXaVlyalZaR09CaS9XY1pwZ0ZIQW85eTVoT3E1NDFMbHJKQlVnUEdHK21hUU5PRVVqbDVGdjY4dHBPdmZLMnpLY2plbVhtOFE9PQ==","index":true}]},{"type":"message","attributes":[{"key":"YWN0aW9u","value":"ZGVwb3NpdA==","index":true}]},{"type":"coin_spent","attributes":[{"key":"c3BlbmRlcg==","value":"dHRob3IxM3dybWhuaDJxZTk4cmpzZTMwcGw3dTZqeHN6amp3bDRmNnl5Y3I=","index":true},{"key":"YW1vdW50","value":"MTAwMDAwMDAwMHJ1bmU=","index":true}]},{"type":"coin_received","attributes":[{"key":"cmVjZWl2ZXI=","value":"dHRob3IxZzk4Y3kzbjltbWpycG4wc3htbjYzbHp0ZWxlcmEzN25yeXR3cDI=","index":true},{"key":"YW1vdW50","value":"MTAwMDAwMDAwMHJ1bmU=","index":true}]},{"type":"transfer","attributes":[{"key":"cmVjaXBpZW50","value":"dHRob3IxZzk4Y3kzbjltbWpycG4wc3htbjYzbHp0ZWxlcmEzN25yeXR3cDI=","index":true},{"key":"c2VuZGVy","value":"dHRob3IxM3dybWhuaDJxZTk4cmpzZTMwcGw3dTZqeHN6amp3bDRmNnl5Y3I=","index":true},{"key":"YW1vdW50","value":"MTAwMDAwMDAwMHJ1bmU=","index":true}]},{"type":"message","attributes":[{"key":"c2VuZGVy","value":"dHRob3IxM3dybWhuaDJxZTk4cmpzZTMwcGw3dTZqeHN6amp3bDRmNnl5Y3I=","index":true}]}]}
```
Summary:

A transaction was sent to deposit 1,000 RUNE with an exchange-related memo for BTC. The transaction was successful (as code:0 indicates no errors) and was included in the block at height "351"

Check amount of synth btc user fox has:
```sh
thornode query bank balances tthor13wrmhnh2qe98rjse30pl7u6jxszjjwl4f6yycr
```
Result:
```sh
root@e688368a1b44:~# thornode query bank balances tthor13wrmhnh2qe98rjse30pl7u6jxszjjwl4f6yycr
balances:
- amount: "45636286"
  denom:  btc/btc
- amount: "199978982000000"
  denom:  rune
pagination:
  next_key:  null
  total:  "0"
```

Now fox deposits btc into vault:

```sh
thornode tx thorchain deposit 45636286 btc/btc ADD:btc/btc:tthor13wrmhnh2qe98rjse30pl7u6jxszjjwl4f6yycr --from fox  $TX_FLAGS
```
Result confirmed transaction.

Now to verify we check the btc synth utilization by calculating from pool endpoint: 

```sh
curl http://localhost:1317/thorchain/pool/BTC.BTC
```
Result:
```sh
david.casalod@bip-dev-04:~/nitka/osmosis-thornode$ curl http://localhost:1317/thorchain/pool/BTC.BTCC
{
  "asset": "BTC.BTC",
  "short_code": "b",
  "status": "Available",
  "pending_inbound_asset": "0",
  "pending_inbound_rune": "0",
  "balance_asset": "250000000",
  "balance_rune": "100993050240",
  "pool_units": "100496500145",
  "LP_units": "100000000000",
  "synth_units": "496500145",
  "synth_supply": "2470236",
  "savers_depth": "0",
  "savers_units": "0",
  "synth_mint_paused": false,
  "synth_supply_remaining": "172529764",
  "loan_collateral": "0",
  "loan_cr": "0",
  "derived_depth_bps": "0"
}
```
As we can see, compared to the beginning, the balance and the synth units have changed.