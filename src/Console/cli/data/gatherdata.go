package data

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"github.com/moti/console/helpers"
	"io/ioutil"
	"log"
	"os"
	"errors"
	"regexp"
	"strings"
)

func OrchestrateGatherTestData() (bool, []helpers.PaternData){
	getFilesError, patternDataList := GetAllTargetFiles()
	if getFilesError != nil{
		return false, nil
	}
	return true, patternDataList
}

func GetAllTargetFiles() (error error, patternMapList []helpers.PaternData){
	var patternDataList []helpers.PaternData

	files, err := ioutil.ReadDir(".")
	if err != nil {
		color.Set(color.FgRed, color.Bold)
		fmt.Printf("Error Occured Reading Working Directory \n")
		color.Unset()
		return errors.New("ERROR: Cant read files, Could be a permissions issue"), nil
	}
	for _, f := range files {
		if strings.Contains(f.Name(), ".tf"){
			patternError, patterData := findPattern(f.Name())
			if patternError != nil{
				color.Set(color.FgRed, color.Bold)
				fmt.Printf("Error Occured Reading file %s...... \n", f.Name())
				color.Unset()
				return errors.New("Could Open File, should be a permissions issue"), nil
			}
			patternDataList  = append(patternDataList, patterData)
		}
	}
	fmt.Print(patternDataList)
	return nil, patternDataList
}


func findPattern(target string) (error error, pattern helpers.PaternData){
	var patternData helpers.PaternData
	file, err := os.Open(target)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	line := 0
	for scanner.Scan() {
		line += 1
		matched, _ := regexp.MatchString(`#!/\s*@Test`, scanner.Text())
		if matched {
			patternData.File = target
			patternData.Line = line
			patternData.InstructionString = strings.TrimSpace(scanner.Text())
			syntaxErr1 := SyntaxAndValidations(&patternData)
			if syntaxErr1 != nil{
				return syntaxErr1, patternData
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return err, patternData
	}
	return nil, patternData
}

func SyntaxAndValidations(pattern *helpers.PaternData) error{
	var declaredInstruction helpers.DeclaredInstruction
	instructionString := strings.TrimSpace(strings.Trim(pattern.InstructionString, "#!/@Test"))
	iVerbs := strings.Split(instructionString, "::")
	if len(iVerbs) != 3 {
		pattern.ErrorMessage = fmt.Sprintf("Eror: Invalid Syntax in file %s at line number %s. Malformed Test Case\n", pattern.File, pattern.Line)
		color.Set(color.FgRed, color.Bold)
		fmt.Printf("Eror: Invalid Syntax in file %s at line number %s. Malformed Test Case\n", pattern.File, pattern.Line)
		color.Unset()
		return errors.New("A Syntax error occured")
	}

	declaredInstruction.Resource = strings.TrimSpace(strings.Trim(iVerbs[0], "()"))
	if len(declaredInstruction.Resource) == 0 {
		pattern.ErrorMessage = fmt.Sprintf("Eror: Invalid Syntax in file %s at line number %s\n. No Resouces defined", pattern.File, pattern.Line)
		color.Set(color.FgRed, color.Bold)
		fmt.Printf("Eror: Invalid Syntax in file %s at line number %s. No Resouces defined\n", pattern.File, pattern.Line)
		color.Unset()
		return errors.New("A Syntax error occured")
	}

	declaredInstruction.Condition = strings.TrimSpace(iVerbs[1])
	if len(declaredInstruction.Condition) == 0 {
		pattern.ErrorMessage = fmt.Sprintf("Eror: Invalid Syntax in file %s at line number %s\n. No Resouces defined", pattern.File, pattern.Line)
		color.Set(color.FgRed, color.Bold)
		fmt.Printf("Eror: Invalid Syntax in file %s at line number %s. No Resouces defined\n", pattern.File, pattern.Line)
		color.Unset()
		return errors.New("A Syntax error occured")
	}
	conditionSupported := false
	for _, v := range helpers.AllowedTestConditions{
		if v == declaredInstruction.Condition{
			conditionSupported = true
		}
	}
	if !conditionSupported {
		pattern.ErrorMessage = fmt.Sprintf("Eror: Invalid Syntax in file %s at line number %s\n. Condition Not Supported or Incorrect", pattern.File, pattern.Line)
		color.Set(color.FgRed, color.Bold)
		fmt.Printf("Eror: Invalid Syntax in file %s at line number %s. Condition Not Supported or Incorrect\n", pattern.File, pattern.Line)
		color.Unset()
		return errors.New("A Syntax error occured")
	}
	helpers.AllowedTestConditions
	return nil
}
