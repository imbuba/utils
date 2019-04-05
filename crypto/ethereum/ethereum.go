package ethereum

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"math/big"
	"strconv"
	"strings"
	"sync"
	"time"

	gabi "github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	cmath "github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/finnan444/utils/math/ints"
	"github.com/finnan444/utils/transport"
	"github.com/onrik/ethrpc"
)

// Some kind of constants
var (
	Pattern        = "999999999999999999999999999999999999"
	PrecisionFloat = big.NewFloat(math.Pow10(18))
	Precision      = big.NewInt(1000000000000000000)
	Zero           = big.NewInt(0)
	ZeroFloat      = big.NewFloat(0)
	EmptyAddress   common.Address
)

var (
	client                  *ethrpc.EthRPC
	config                  *Config
	lastBlocksLocker        sync.RWMutex
	lastBlocks              map[string]uint64
	abis                    map[string]*abiWraper
	blockDifferenceReporter BlockDifferenceReporter
)

type abiWraper struct {
	Abi          *gabi.ABI
	Events       map[string]string
	Address      string
	AddressSlice []string
	EventTopics  map[string][][]string
}

// SignMessageFunc describes sign message function
type SignMessageFunc func(from, signerName string, values ...interface{}) (message string, err error)

// RecoverSignFunc describes recover sign function
type RecoverSignFunc func(message, sign string) (address string, err error)

// BlockDifferenceReporter describes function to report block difference
type BlockDifferenceReporter func(delta, blockNumber, worldBlockNumber int64)

// IteratorFunc process iteration
type IteratorFunc func(index int, log ethrpc.Log) bool

// SaverFunc saves blockNumber with a given tag
type SaverFunc func(blockNumber uint64, tag string) error

// IsAddressValid checks address for validity
func IsAddressValid(address string) bool {
	return len(address) == 42 && address[:2] == "0x"
}

func parseHex(hexValue string) string {
	if hexValue == "0x" {
		return "0"
	}
	return strings.Replace(hexValue, "0x", "", -1)
}

func getIntFromHex(hexValue string) *big.Int {
	result, _ := new(big.Int).SetString(parseHex(hexValue), 16)
	return result
}

// Initialize initialized ethereum
func Initialize(c *Config, initLastBlocks func() map[string]uint64, bdr BlockDifferenceReporter) (SignMessageFunc, RecoverSignFunc, error) {
	connectAddress := fmt.Sprintf("%s://%s:%d", c.Protocol, c.Host, c.Port)
	client = ethrpc.NewEthRPC(connectAddress)
	if version, err := client.Web3ClientVersion(); err != nil {
		return nil, nil, err
	} else {
		log.Printf("Connected to RPC at: %s, version: %s", connectAddress, version)
	}
	client.Debug = c.Logging
	abis = make(map[string]*abiWraper)
	lastBlocks = initLastBlocks()
	for k, v := range c.Addresses {
		if !IsAddressValid(v) {
			return nil, nil, fmt.Errorf("[Ethereum] %s address is invalid", k)
		}
	}
	config = c
	blockDifferenceReporter = bdr
	for k, v := range c.Definitions {
		if !IsAddressValid(v.Address) {
			return nil, nil, fmt.Errorf("[Ethereum] %s address is invalid", k)
		}
		wraper := &abiWraper{Address: v.Address, Events: map[string]string{}, AddressSlice: []string{v.Address}, EventTopics: map[string][][]string{}}
		if abi, err := gabi.JSON(strings.NewReader(v.Abi)); err != nil {
			return nil, nil, fmt.Errorf("[Ethereum] error parsing abi definition %s", k)
		} else {
			wraper.Abi = &abi
			for _, event := range v.Events {
				if ev, ok := abi.Events[event]; ok {
					eventAddress := ev.Id().Hex()
					wraper.Events[event] = eventAddress
					wraper.EventTopics[event] = [][]string{[]string{eventAddress}}
				} else {
					return nil, nil, fmt.Errorf("[Ethereum] no such event %s in abi definition %s", event, v.Abi)
				}
			}
			abis[k] = wraper
		}
	}
	go checkBlockDiff()
	return signMessage, recoverSign, nil
}

func checkBlockDiff() {
	ticker := time.NewTicker(5 * time.Minute)
	for range ticker.C {
		go checkBlockDifference()
	}
}

// UnpackEvent unpacks data to a given event
func UnpackEvent(abiName, eventName string, event interface{}, data string) error {
	if abi, ok := abis[abiName]; ok {
		return abi.Abi.Unpack(event, eventName, common.FromHex(data))
	}
	return fmt.Errorf("[Ethereum] no such abi %s", eventName)
}

func balanceOf(fromAddressName, toAddressName, checkingAddress string) (*big.Int, error) {
	if fromAddress, ok := config.Addresses[fromAddressName]; ok {
		if contractAddress, ok := config.Addresses[toAddressName]; ok {
			balanceOfHash := "0x70a08231"
			callData := fmt.Sprintf("%s000000000000000000000000%s", balanceOfHash, parseHex(checkingAddress))
			strResult, err := client.EthCall(ethrpc.T{From: fromAddress, To: contractAddress, Data: callData}, "latest")
			if err != nil {
				log.Printf("[Ethereum] balanceOf: rpc.client.EthCall return error: %v", err)
				return nil, err
			}
			result := getIntFromHex(strResult)
			if err != nil {
				log.Printf("[Ethereum] balanceOf: strconv.ParseInt return error: %v", err)
				return nil, err
			}
			return result, nil
		}
		return nil, fmt.Errorf("[Ethereum] no such contract address %s", toAddressName)
	}
	return nil, fmt.Errorf("[Ethereum] no such from address %s", fromAddressName)
}

func signMessage(from, signerName string, values ...interface{}) (string, error) {
	if signer, ok := config.Addresses[signerName]; ok {
		fromAddress := common.HexToAddress(from)
		args := make([][]byte, 0)
		args = append(args, fromAddress.Bytes())
		if config.Logging {
			log.Printf("Sign arg #%d: %x", len(args), args[len(args)-1])
		}
		for _, val := range values {
			if bi, ok := val.(*big.Int); ok {
				args = append(args, cmath.PaddedBigBytes(bi, 32))
				if config.Logging {
					log.Printf("Sign arg #%d: %x", len(args), args[len(args)-1])
				}
			} else if intVal, ok := val.(int); ok {
				args = append(args, cmath.PaddedBigBytes(big.NewInt(int64(intVal)), 32))
				if config.Logging {
					log.Printf("Sign arg #%d: %x", len(args), args[len(args)-1])
				}
			} else if uintVal, ok := val.(uint); ok {
				args = append(args, cmath.PaddedBigBytes(big.NewInt(int64(uintVal)), 32))
				if config.Logging {
					log.Printf("Sign arg #%d: %x", len(args), args[len(args)-1])
				}
			} else if str, ok := val.(string); ok {
				args = append(args, []byte(str))
				if config.Logging {
					log.Printf("Sign arg #%d: %x", len(args), args[len(args)-1])
				}
			}
		}
		str := fmt.Sprintf("0x%x", crypto.Keccak256(args...))
		if config.Logging {
			log.Printf("[Ethereum] SignMessage keccak is %s", str)
		}
		return client.EthSign(signer, str)
	}
	return "", fmt.Errorf("[Ethereum] no such signer %s", signerName)
}

func recoverSign(msg, sign string) (string, error) {
	if config.Logging {
		log.Printf("[Ethereum] Start recovering msg %s; sign %s", msg, sign)
	}
	if len(msg) > 2 && strings.ToLower(msg[:2]) != "0x" {
		msg = fmt.Sprintf("0x%x", []byte(msg))
	}
	return client.PersonalEcRecover(msg, sign)
}

func getLastBlock(id string) uint64 {
	lastBlocksLocker.RLock()
	if res, ok := lastBlocks[id]; ok {
		lastBlocksLocker.RUnlock()
		return res
	}
	lastBlocksLocker.RUnlock()
	return 1
}

func checkBlockDifference() {
	httpClient := transport.GetHTTPClient()
	sc, body, err := httpClient.GetTimeout(nil, config.CheckWorldBlockURL, time.Second*5)
	transport.PutHTTPClient(httpClient)
	if err != nil {
		log.Printf("[Ethereum] error getting world stats %v", err)
	} else {
		if sc == 200 {
			var nodeInfo struct {
				Result string
			}
			err = json.Unmarshal(body, &nodeInfo)
			if err != nil {
				log.Printf("[Ethereum] error unmarshalling block number %v", err)
			} else {
				if worldBlockNumber, err := strconv.ParseInt(parseHex(nodeInfo.Result), 16, 64); err != nil {
					log.Printf("[Ethereum] error parsing block number %v", err)
				} else {
					if blockNumberInt, err := client.EthBlockNumber(); err != nil {
						log.Printf("[Ethereum] error getting our block number %v", err)
					} else {
						blockDifferenceReporter(ints.AbsInt64(int64(blockNumberInt)-worldBlockNumber), int64(blockNumberInt), worldBlockNumber)
					}
				}
			}
		} else {
			log.Printf("[Ethereum] Check world block failed with %d", sc)
		}
	}
}

// ProcessingEthereumTransactions processed ethereum transactions for events
func ProcessingEthereumTransactions(wg *sync.WaitGroup, abiName, eventName, tag string, iterator IteratorFunc, saverFunc SaverFunc, tillBlock uint64) {
	if abi, ok := abis[abiName]; ok {
		if eventAddress, ok := abi.Events[eventName]; ok {
			if config.Logging {
				log.Printf("[Ethereum] Start processing logs from block: %d; %s address %s; %s topic %s", getLastBlock(tag), abiName, abi.Address, eventName, eventAddress)
			}
			lastBl := getLastBlock(tag)
			var fp ethrpc.FilterParams
			if lastBl+50000 < tillBlock {
				fp = ethrpc.FilterParams{
					FromBlock: fmt.Sprintf("0x%x", lastBl),
					ToBlock:   fmt.Sprintf("0x%x", lastBl+50000),
					Address:   abi.AddressSlice,
					Topics:    abi.EventTopics[eventName],
				}
			} else {
				fp = ethrpc.FilterParams{
					FromBlock: fmt.Sprintf("0x%x", lastBl),
					ToBlock:   "latest",
					Address:   abi.AddressSlice,
					Topics:    abi.EventTopics[eventName],
				}
			}
			if logs, err := client.EthGetLogs(fp); err != nil {
				log.Printf("[Ethereum][%s] Error getting logs, err %v", tag, err)
			} else {
				if config.Logging {
					log.Printf("[Ethereum] Logs received %d", len(logs))
				}
				earliestBlockNumber := 0
				checkFurther := true
				for idx, lg := range logs {
					processed := iterator(idx, lg)
					if checkFurther {
						if !processed {
							checkFurther = false
						}
						earliestBlockNumber = lg.BlockNumber
					}
				}
				if checkFurther && earliestBlockNumber > 0 {
					earliestBlockNumber++
				}
				if len(logs) == 0 && fp.ToBlock != "latest" {
					earliestBlockNumber = int(lastBl + 50000)
				}
				lll := getLastBlock(tag)
				if earliestBlockNumber > 0 && lll != uint64(earliestBlockNumber) {
					log.Printf("[Ethereum][Log][%s] presave earliest = %d and saved = %d", tag, earliestBlockNumber, lll)
					if err := saverFunc(uint64(earliestBlockNumber), tag); err != nil {
						log.Printf("[Ethereum][Log][%s] save last transaction block error %v", tag, err)
					} else {
						lastBlocksLocker.Lock()
						lastBlocks[tag] = uint64(earliestBlockNumber)
						lastBlocksLocker.Unlock()
					}
				}
			}
		}
	}
	wg.Done()
}
