package main

import (
	"fmt"

	student ".."
)

func main() {
	str := "|=choumi=|which|=choumi=|itself|=choumi=|used|=choumi=|cards|=choumi=|and|=choumi=|a|=choumi=|central|=choumi=|computing|=choumi=|unit.|=choumi=|When|=choumi=|the|=choumi=|machine|=choumi=|was|=choumi=|finished,"
	fmt.Println(student.Split(str, "|=choumi=|"))
	str = "!==!which!==!was!==!making!==!all!==!kinds!==!of!==!punched!==!card!==!equipment!==!and!==!was!==!also!==!in!==!the!==!calculator!==!business[10]!==!to!==!develop!==!his!==!giant!==!programmable!==!calculator,"
	fmt.Println(student.Split(str, "!==!"))
	str = "AFJCharlesAFJBabbageAFJstartedAFJtheAFJdesignAFJofAFJtheAFJfirstAFJautomaticAFJmechanicalAFJcalculator,"
	fmt.Println(student.Split(str, "AFJ"))
	str = "<<==123==>>In<<==123==>>1820,<<==123==>>Thomas<<==123==>>de<<==123==>>Colmar<<==123==>>launched<<==123==>>the<<==123==>>mechanical<<==123==>>calculator<<==123==>>industry[note<<==123==>>1]<<==123==>>when<<==123==>>he<<==123==>>released<<==123==>>his<<==123==>>simplified<<==123==>>arithmometer,"
	fmt.Println(student.Split(str, "<<==123==>>"))

}

// [ which itself used cards and a central computing unit. When the machine was finished,]
// [ which was making all kinds of punched card equipment and was also in the calculator business[10] to develop his giant programmable calculator,]
// [ Charles Babbage started the design of the first automatic mechanical calculator,]
// [ In 1820, Thomas de Colmar launched the mechanical calculator industry[note 1] when he released his simplified arithmometer,]
