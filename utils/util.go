package utils

import (
	"fmt"

	"example.com/main/configs"
	"example.com/main/types"
	"github.com/hashgraph/hedera-sdk-go/v2"
)


type SetDataParameter struct {
	file_data string
	user string
	file_name string
}

func ConnectHedera() *hedera.Client {
	var client *hedera.Client
	var err error

	client = hedera.ClientForTestnet()
	if err != nil {
		println(err.Error(), ": error creating client")
	}
	configOperatorID := configs.EnvHederaOperatorID()
	configOperatorKey := configs.EnvHederaOperatorPrivateKey()

	//client.SetOperator(configOperatorID, configOperatorKey)

	if configOperatorID != "" && configOperatorKey != "" {
		fmt.Println("In if")
		operatorAccountID, err := hedera.AccountIDFromString(configOperatorID)
		if err != nil {
			println(err.Error(), ": error converting string to AccountID")
		}

		operatorKey, err := hedera.PrivateKeyFromString(configOperatorKey)
		if err != nil {
			println(err.Error(), ": error converting string to PrivateKey")
		}
		// fmt.Println(operatorAccountID.Realm)
		// fmt.Println(operatorKey.String())
		client.SetOperator(operatorAccountID, operatorKey)
	}
	fmt.Println("Connected to Hedera")
	return client

}

var hederaClient *hedera.Client = ConnectHedera()
var contractId string = configs.EnvContractID() 

func executeTransaction(functionName string, parametersArray []types.ParametersType) *hedera.ContractFunctionResult {
	// functionVariables := 
	FinalFunctionVariables := addFunctionParameters(parametersArray, &*hedera.NewContractFunctionParameters() ,0)
	//fmt.Println(&FinalFunctionVariables)
	// functionVariables :=  hedera.NewContractFunctionParameters().AddString("Test doc2").AddString("user1").AddString("test_doc")
	newContractId,err := hedera.ContractIDFromString(contractId)
	//newfunctionVariables := *hedera.ContractFunctionParameters
	transaction := hedera.NewContractExecuteTransaction().
					   SetContractID(newContractId).
					   SetGas(10000000).
					   SetFunction(functionName, FinalFunctionVariables)
	//Sign with the client operator privateclient key to pay for the transaction and submit the query to a Hedera network
	txResponse, err := transaction.Execute(hederaClient)
	fmt.Println("tx Response")
	fmt.Println(txResponse)
	if err != nil {
		panic(err)
	}

	// Get Transaction Record 
	txRecord, err := txResponse.GetRecord(hederaClient)

	// Get Contract Result
	contractResult, err := txRecord.GetContractExecuteResult()
	
	
	fmt.Println(txRecord.CallResult.LogInfo[0].Data)
	
	//Request the receipt of the transaction
	txReceipt, err := txResponse.GetReceipt(hederaClient)

	if err != nil {
		panic(err)
	}
	fmt.Println((txReceipt))
	//Get the transaction consensus status
	transactionStatus := txReceipt.Status

	fmt.Printf("The transaction consensus status %v\n", transactionStatus)

	return &contractResult

}

func addFunctionParameters(parameterArray []types.ParametersType, functionParameter *hedera.ContractFunctionParameters, index int) *hedera.ContractFunctionParameters {
	//newFunctionParameter *hedera.ContractFunctionParameters := nil
	if(index >= len(parameterArray)) {
		return functionParameter
	}
	fmt.Println(parameterArray[index])

	switch parameterArray[index].Datatype {
	case "string":
		functionParameter = functionParameter.AddString(parameterArray[index].Value)
		break
	case "int32":
		functionParameter = functionParameter.AddInt32(200)	
	}

	return addFunctionParameters(parameterArray, functionParameter, index+1)
}

func SetData(parametersArray []types.ParametersType) types.SetDataOutput {
	// Execute Transaction
	contractResult := executeTransaction("setData", parametersArray)

	// Get Output Parameters
	hash := contractResult.GetBytes32(0)
	hashString := string(hash[:])
	
	timestamp := contractResult.GetUint256(1)
	timestampString := string(timestamp[:])

	// jsonFile, err := os.Open("./notarization.json")

	//  // if we os.Open returns an error then handle it
    // if err != nil {
    //     fmt.Println(err)
    // }

	// defer jsonFile.Close()

	// byteValue, _ := ioutil.ReadAll(jsonFile)

    // var result map[string]interface{}
    
	// json.Unmarshal([]byte(byteValue), &result)

    // fmt.Println(result)
	//data := *hedera.ContractFunctionResultFromBytes(hash)
	//fmt.Println(data)
	setDataOutput := types.SetDataOutput{ Hash:hashString, Timestamp:timestampString }

	return setDataOutput
}

