

foreach($day in 8..24){
    $path = "./day$($day)"

    $header = "package day$($day)"

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

    $files= @(
        "day$($day)_dataparser.go",
        "day$($day)_test.go",
        "day$($day)_util.go"
    )
    
    foreach($fn in $files){
        echo $header > $path\$fn
    }    

    echo $solnTemplate.Replace("##DAY##",$day) > "$path\\day$($day).go"

}
