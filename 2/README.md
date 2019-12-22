# Statefulset

This is a _very_ raw statefulset for Litecoind. There would be many additions I would add, but don't understand the application well enough to tackle them now:

- Healthcheck/liveness checks: How do I konw the application is healthy/running?
- Connectivity - Do these applications need to talk to eachother, or is the use of a Statefulset simply to manage PVC claims more easily?
- I would extract the config used by the daemon into a Secret to prevent injecting secrets in the Docker file

## Running:

```bash
$ minikube start
😄  minikube v1.5.0 on Darwin 10.15.2
🔥  Creating virtualbox VM (CPUs=4, Memory=12000MB, Disk=30000MB) ...
🐳  Preparing Kubernetes v1.16.2 on Docker 18.09.9 ...
🚜  Pulling images ...
🚀  Launching Kubernetes ...
⌛  Waiting for: apiserver proxy etcd scheduler controller dns
🏄  Done! kubectl is now configured to use "minikube"

$ kubectl apply -f 2/litecoind-statefulset.yaml
service/litecoind created
statefulset.apps/litecoind created

$ kubectl get pods,pvc
NAME              READY   STATUS    RESTARTS   AGE
pod/litecoind-0   1/1     Running   0          69s

NAME                                     STATUS   VOLUME                                     CAPACITY   ACCESS MODES   STORAGECLASS   AGE
persistentvolumeclaim/data-litecoind-0   Bound    pvc-472dd0b6-16fb-4810-b986-4d4ca9ab3c22   25Gi       RWO            standard       69s

$ kubectl logs litecoind-0
2019-12-22T16:08:48Z Litecoin Core version v0.17.1 (release build)
2019-12-22T16:08:48Z InitParameterInteraction: parameter interaction: -bind set -> setting -listen=1
2019-12-22T16:08:48Z InitParameterInteraction: parameter interaction: -whitelistforcerelay=1 -> setting -whitelistrelay=1
2019-12-22T16:08:48Z Assuming ancestors of block a601455787cb65ffc325dda4751a99cf01d1567799ec4b04f45bb05f9ef0cbde have valid signatures.
2019-12-22T16:08:48Z Setting nMinimumChainWork=0000000000000000000000000000000000000000000001488d32719b8eb30150
2019-12-22T16:08:48Z Using the 'sse4(1way),sse41(4way),avx2(8way)' SHA256 implementation
2019-12-22T16:08:48Z Using RdRand as an additional entropy source
2019-12-22T16:08:48Z Default data directory /home/litecoin/.litecoin
2019-12-22T16:08:48Z Using data directory /data/litecoin
2019-12-22T16:08:48Z Using config file /data/litecoin/litecoin.conf
2019-12-22T16:08:48Z Using at most 125 automatic connections (1073741816 file descriptors available)
2019-12-22T16:08:48Z Using 16 MiB out of 32/2 requested for signature cache, able to store 524288 elements
2019-12-22T16:08:48Z Using 16 MiB out of 32/2 requested for script execution cache, able to store 524288 elements
2019-12-22T16:08:48Z Using 4 threads for script verification
2019-12-22T16:08:48Z scheduler thread start
2019-12-22T16:08:48Z WARNING: option -rpcbind was ignored because -rpcallowip was not specified, refusing to allow everyone to connect
2019-12-22T16:08:48Z libevent: getaddrinfo: address family for nodename not supported
2019-12-22T16:08:48Z Binding RPC on address ::1 port 9332 failed.
2019-12-22T16:08:48Z HTTP: creating work queue of depth 16
2019-12-22T16:08:48Z Config options rpcuser and rpcpassword will soon be deprecated. Locally-run instances may remove rpcuser to use cookie-based auth, or may be replaced with rpcauth. Please see share/rpcauth for rpcauth auth generation.
2019-12-22T16:08:48Z HTTP: starting 4 worker threads
2019-12-22T16:08:48Z Using wallet directory /data/litecoin
2019-12-22T16:08:48Z init message: Verifying wallet(s)...
2019-12-22T16:08:48Z Using BerkeleyDB version Berkeley DB 4.8.30: (April  9, 2010)
2019-12-22T16:08:48Z Using wallet wallet.dat
2019-12-22T16:08:48Z BerkeleyEnvironment::Open: LogDir=/data/litecoin/database ErrorFile=/data/litecoin/db.log
2019-12-22T16:08:48Z Cache configuration:
2019-12-22T16:08:48Z * Using 2.0MiB for block index database
2019-12-22T16:08:48Z * Using 56.0MiB for transaction index database
2019-12-22T16:08:48Z * Using 8.0MiB for chain state database
2019-12-22T16:08:48Z * Using 384.0MiB for in-memory UTXO set (plus up to 286.1MiB of unused mempool space)
2019-12-22T16:08:48Z init message: Loading block index...
2019-12-22T16:08:48Z Opening LevelDB in /data/litecoin/blocks/index
2019-12-22T16:08:48Z Opened LevelDB successfully
2019-12-22T16:08:48Z Using obfuscation key for /data/litecoin/blocks/index: 0000000000000000
2019-12-22T16:08:48Z LoadBlockIndexDB: last block file = 0
2019-12-22T16:08:48Z LoadBlockIndexDB: last block file info: CBlockFileInfo(blocks=0, size=0, heights=0...0, time=1970-01-01...1970-01-01)
2019-12-22T16:08:48Z Checking all blk files are present...
2019-12-22T16:08:48Z Initializing databases...
2019-12-22T16:08:48Z Pre-allocating up to position 0x1000000 in blk00000.dat
2019-12-22T16:08:48Z Opening LevelDB in /data/litecoin/chainstate
2019-12-22T16:08:48Z Opened LevelDB successfully
2019-12-22T16:08:48Z Wrote new obfuscate key for /data/litecoin/chainstate: f69d494cd63479d7
2019-12-22T16:08:48Z Using obfuscation key for /data/litecoin/chainstate: f69d494cd63479d7
2019-12-22T16:08:48Z init message: Rewinding blocks...
2019-12-22T16:08:48Z  block index               9ms
2019-12-22T16:08:48Z Opening LevelDB in /data/litecoin/indexes/txindex
2019-12-22T16:08:48Z Opened LevelDB successfully
2019-12-22T16:08:48Z Using obfuscation key for /data/litecoin/indexes/txindex: 0000000000000000
2019-12-22T16:08:48Z init message: Loading wallet...
2019-12-22T16:08:48Z txindex thread start
2019-12-22T16:08:48Z txindex is enabled
2019-12-22T16:08:48Z txindex thread exit
2019-12-22T16:08:48Z [default wallet] nFileVersion = 170100
2019-12-22T16:08:48Z [default wallet] Keys: 0 plaintext, 0 encrypted, 0 w/ metadata, 0 total. Unknown wallet records: 0
2019-12-22T16:08:48Z [default wallet] Performing wallet upgrade to 169900
2019-12-22T16:08:49Z [default wallet] keypool added 2000 keys (1000 internal), size=2000 (1000 internal)
2019-12-22T16:08:49Z [default wallet] Wallet completed loading in             898ms
2019-12-22T16:08:49Z [default wallet] setKeyPool.size() = 2000
2019-12-22T16:08:49Z [default wallet] mapWallet.size() = 0
2019-12-22T16:08:49Z [default wallet] mapAddressBook.size() = 0
2019-12-22T16:08:49Z UpdateTip: new best=12a765e31ffd4059bada1e25190f6e98c99d9714d334efa41a195a7e7e04bfe2 height=0 version=0x00000001 log2_work=20.000022 tx=1 date='2011-10-07T07:31:05Z' progress=0.000000 cache=0.0MiB(0txo)
2019-12-22T16:08:49Z Failed to open mempool file from disk. Continuing anyway.
2019-12-22T16:08:49Z mapBlockIndex.size() = 1
2019-12-22T16:08:49Z nBestHeight = 0
2019-12-22T16:08:49Z torcontrol thread start
2019-12-22T16:08:49Z Bound to 127.0.0.1:9333
2019-12-22T16:08:49Z init message: Loading P2P addresses...
2019-12-22T16:08:49Z ERROR: DeserializeFileDB: Failed to open file /data/litecoin/peers.dat
2019-12-22T16:08:49Z Invalid or missing peers.dat; recreating
2019-12-22T16:08:49Z init message: Loading banlist...
2019-12-22T16:08:49Z ERROR: DeserializeFileDB: Failed to open file /data/litecoin/banlist.dat
2019-12-22T16:08:49Z Invalid or missing banlist.dat; recreating
2019-12-22T16:08:49Z init message: Starting network threads...
2019-12-22T16:08:49Z net thread start
2019-12-22T16:08:49Z dnsseed thread start
2019-12-22T16:08:49Z addcon thread start
2019-12-22T16:08:49Z Loading addresses from DNS seeds (could take a while)
2019-12-22T16:08:49Z init message: Done loading
2019-12-22T16:08:49Z opencon thread start
2019-12-22T16:08:49Z msghand thread start
2019-12-22T16:08:50Z New outbound peer connected: version: 70015, blocks=1757717, peer=0
2019-12-22T16:08:52Z New outbound peer connected: version: 70015, blocks=1757717, peer=1
2019-12-22T16:08:53Z New outbound peer connected: version: 70015, blocks=1757717, peer=2
2019-12-22T16:08:53Z New outbound peer connected: version: 70015, blocks=1757717, peer=3
2019-12-22T16:08:54Z New outbound peer connected: version: 70015, blocks=1757717, peer=4
2019-12-22T16:08:56Z 95 addresses found from DNS seeds
2019-12-22T16:08:56Z dnsseed thread exit
2019-12-22T16:08:57Z New outbound peer connected: version: 70015, blocks=1757717, peer=5
```
