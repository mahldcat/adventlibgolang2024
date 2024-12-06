

foreach($day in 8..24){
    $path = "./day$($day)"

    $header = "package day$($day)"

    $files= @(
        "day$($day).go",
        "day$($day)_dataparser.go",
        "day$($day)_test.go",
        "day$($day)_util.go"
    )
    
    foreach($fn in $files){
        echo $header > $path\$fn
    }    
}
