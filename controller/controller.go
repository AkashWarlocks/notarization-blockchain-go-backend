package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"example.com/main/types"
	"example.com/main/utils"
)

func SetData(c *gin.Context) {
	var setDataParameter types.SetData
	
	c.BindJSON((&setDataParameter))
	//var parameterArray utils.ParametersType
	datatype := "string"	
	Document := types.ParametersType{ 
		Datatype: datatype, 
		Value: setDataParameter.Documemt, 
	} 

	User := types.ParametersType{
		Datatype:"string", 
		Value:setDataParameter.User,
	}
	DocumentName := types.ParametersType{
		Datatype:"string", 
		Value:setDataParameter.DocumentName,
	}

	parametersArray := []types.ParametersType{ Document,User, DocumentName}
	
	fmt.Println(parametersArray)
	setDataReturn := utils.SetData(parametersArray)
	fmt.Println(setDataReturn)
	// returnData, err := json.Marshal(&setDataReturn) 
	// fmt.Println(returnData)
	// if err != nil {
    //     fmt.Println(err)
        
    // }
	c.JSON(http.StatusOK, setDataReturn)
	
}