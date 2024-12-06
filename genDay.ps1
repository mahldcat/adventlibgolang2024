

foreach($day in 6..24){
    $path = "./day$($day)"

    $solnUtilTemplate = 
@'
package day##DAY##
'@

    $solnParserTemplate=
@'
package day##DAY##

import (
	"errors"
)

func day5DataParser(rawData string) (error) {
    return errors.New("not implemented")
}
'@

    $solnTestTemplate=
@'
package day##DAY##

import (
	"testing"
)

var exampleRaw = ""

func TestParser(t *testing.T) {

}

func TestDay##DAY##Part1(t *testing.T) {
	expected := -1
	sln := SolveDay##DAY##Part1(exampleRaw)

	if sln != expected {
		t.Errorf("mismatched solution value expected %d, actual %d", expected, sln)
	}

}

func TestDay##DAY##Part2(t *testing.T) {
	expected := -1
	sln := SolveDay##DAY##Part2(exampleRaw)

	if sln != expected {
		t.Errorf("mismatched solution value expected %d, actual %d", expected, sln)
	}
}
'@

    $solnTemplate = 
@"
package day##DAY##


func SolveDay##DAY##Part1(rawData string) int {
    solnResult:=-1

    return solnResult
}

func SolveDay##DAY##Part2(rawData string) int {
    solnResult:=-1

    return solnResult
}

"@

    rm "$path/*"

    echo $solnTemplate.Replace("##DAY##",$day) > "$path\\day$($day).go"
    echo $solnTestTemplate.Replace("##DAY##",$day) > "$path\\day$($day)_test.go"
    echo $solnParserTemplate.Replace("##DAY##",$day) > "$path\\day$($day)_dataparser.go"
    echo $solnUtilTemplate.Replace("##DAY##",$day) > "$path\\day$($day)_util.go"
 

}
