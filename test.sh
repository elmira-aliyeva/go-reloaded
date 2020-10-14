cd ./cat
go build

if [ ! -f quest8.txt ]; then
    cp /home/kali/div-01/go-reloaded/cat/quest8.txt .
fi

if [ ! -f quest8T.txt ]; then
    cp /home/kali/div-01/go-reloaded/cat/quest8T.txt .
fi

PASS=true

OUTPUT1="$(./cat quest8.txt)"
OUTPUT2="$(cat quest8.txt)"

if [ "$OUTPUT1" != "$OUTPUT2" ]; then
    echo "FAIL: ./cat quest8.txt"
    PASS=false
    diff  <(echo "$OUTPUT1" ) <(echo "$OUTPUT2")
fi

OUTPUT1="$(./cat quest8T.txt)"
OUTPUT2="$(cat quest8T.txt)"

if [ "$OUTPUT1" != "$OUTPUT2" ]; then
    echo "FAIL: ./cat quest8T.txt"
    PASS=false
    diff  <(echo "$OUTPUT1" ) <(echo "$OUTPUT2")
fi

OUTPUT1="$(./cat quest8.txt quest8T.txt)"
OUTPUT2="$(cat quest8.txt quest8T.txt)"

if [ "$OUTPUT1" != "$OUTPUT2" ]; then
    echo "FAIL: ./cat quest8.txt quest8T.txt"
    PASS=false
    diff  <(echo "$OUTPUT1" ) <(echo "$OUTPUT2")
fi

OUTPUT1="$(./cat scddf)"
OUTPUT2="$(cat scddf 2>&1)"

if [ "$OUTPUT1" != "$OUTPUT2" ]; then
    echo "FAIL: ./cat scddf"
    PASS=false
    diff  <(echo "$OUTPUT1" ) <(echo "$OUTPUT2")
fi

cd ../doop
go build

OUTPUT1="$(./doop 861 + 870)"
OUTPUT2="1731"

if [ "$OUTPUT1" != "$OUTPUT2" ]; then
    echo "FAIL: ./doop 861 + 870"
    PASS=false
    diff  <(echo "$OUTPUT1" ) <(echo "$OUTPUT2")
fi

OUTPUT1="$(./doop 861 - 870)"
OUTPUT2="-9"

if [ "$OUTPUT1" != "$OUTPUT2" ]; then
    echo "FAIL: ./doop 861 - 870"
    PASS=false
    diff  <(echo "$OUTPUT1" ) <(echo "$OUTPUT2")
fi

OUTPUT1="$(./doop 861 "*" 870)"
OUTPUT2="749070"

if [ "$OUTPUT1" != "$OUTPUT2" ]; then
    echo "FAIL: ./doop 861 "*" 870"
    PASS=false
    diff  <(echo "$OUTPUT1" ) <(echo "$OUTPUT2")
fi

OUTPUT1="$(./doop 861 % 870)"
OUTPUT2="861"

if [ "$OUTPUT1" != "$OUTPUT2" ]; then
    echo "FAIL: ./doop 861 % 870"
    PASS=false
    diff  <(echo "$OUTPUT1" ) <(echo "$OUTPUT2")
fi

OUTPUT1="$(./doop  hello + 1)"
OUTPUT2="0"

if [ "$OUTPUT1" != "$OUTPUT2" ]; then
    echo "FAIL: ./doop  hello + 1"
    PASS=false
    diff  <(echo "$OUTPUT1" ) <(echo "$OUTPUT2")
fi

OUTPUT1="$(./doop 1 p 1)"
OUTPUT2="0"

if [ "$OUTPUT1" != "$OUTPUT2" ]; then
    echo "FAIL: ./doop 1 p 1"
    PASS=false
    diff  <(echo "$OUTPUT1" ) <(echo "$OUTPUT2")
fi

OUTPUT1="$(./doop 1 / 0)"
OUTPUT2="No division by 0"

if [ "$OUTPUT1" != "$OUTPUT2" ]; then
    echo "FAIL: ./doop 1 / 0"
    PASS=false
    diff  <(echo "$OUTPUT1" ) <(echo "$OUTPUT2")
fi

OUTPUT1="$(./doop 1 % 0)"
OUTPUT2="No modulo by 0"

if [ "$OUTPUT1" != "$OUTPUT2" ]; then
    echo "FAIL: ./doop 1 % 0"
    PASS=false
    diff  <(echo "$OUTPUT1" ) <(echo "$OUTPUT2")
fi

OUTPUT1="$(./doop  1 "*" 1)"
OUTPUT2="1"

if [ "$OUTPUT1" != "$OUTPUT2" ]; then
    echo "FAIL: ./doop  1 "*" 1"
    PASS=false
    diff  <(echo "$OUTPUT1" ) <(echo "$OUTPUT2")
fi

OUTPUT1="$(./doop 9223372036854775807 + 1)"
OUTPUT2="0"

if [ "$OUTPUT1" != "$OUTPUT2" ]; then
    echo "FAIL: ./doop 9223372036854775807 + 1"
    PASS=false
    diff  <(echo "$OUTPUT1" ) <(echo "$OUTPUT2")
fi

OUTPUT1="$(./doop 9223372036854775809 - 3)"
OUTPUT2="0"

if [ "$OUTPUT1" != "$OUTPUT2" ]; then
    echo "FAIL: ./doop 9223372036854775809 - 3"
    PASS=false
    diff  <(echo "$OUTPUT1" ) <(echo "$OUTPUT2")
fi

OUTPUT1="$(./doop 9223372036854775807 "*" 3)"
OUTPUT2="0"

if [ "$OUTPUT1" != "$OUTPUT2" ]; then
    echo "FAIL: ./doop 9223372036854775807 "*" 3"
    PASS=false
    diff  <(echo "$OUTPUT1" ) <(echo "$OUTPUT2")
fi

cd ../rotatevowels
go build

OUTPUT1="$(./rotatevowels "Hello World")"
OUTPUT2="Hollo Werld"

if [ "$OUTPUT1" != "$OUTPUT2" ]; then
    echo "FAIL: ./rotatevowels \"Hello World\""
    PASS=false
    diff  <(echo "$OUTPUT1" ) <(echo "$OUTPUT2")
fi

OUTPUT1="$(./rotatevowels "HEllO World" "problem solved")"
OUTPUT2="Hello Werld problom sOlvEd"

if [ "$OUTPUT1" != "$OUTPUT2" ]; then
    echo "FAIL: ./rotatevowels \"HEllO World\" \"problem solved\""
    PASS=false
    diff  <(echo "$OUTPUT1" ) <(echo "$OUTPUT2")
fi

OUTPUT1="$(./rotatevowels "str" "shh" "psst")"
OUTPUT2="str shh psst"

if [ "$OUTPUT1" != "$OUTPUT2" ]; then
    echo "FAIL: ./rotatevowels \"str\" \"shh\" \"psst\""
    PASS=false
    diff  <(echo "$OUTPUT1" ) <(echo "$OUTPUT2")
fi

OUTPUT1="$(./rotatevowels "happy thoughts" "good luck")"
OUTPUT2="huppy thooghts guod lack"

if [ "$OUTPUT1" != "$OUTPUT2" ]; then
    echo "FAIL: ./rotatevowels \"happy thoughts\" \"good luck\""
    PASS=false
    diff  <(echo "$OUTPUT1" ) <(echo "$OUTPUT2")
fi

OUTPUT1="$(./rotatevowels "al's elEphAnt is overly underweight!")"
OUTPUT2="il's elephunt es ovirly AndErweaght!"

if [ "$OUTPUT1" != "$OUTPUT2" ]; then
    echo "FAIL: ./rotatevowels \"al's elEphAnt is overly underweight!\""
    PASS=false
    diff  <(echo "$OUTPUT1" ) <(echo "$OUTPUT2")
fi

cd ../ztail
go build

if [ ! -f quest8.txt ]; then
    cp /home/kali/div-01/go-reloaded/ztail/quest8.txt .
fi

if [ ! -f quest8T.txt ]; then
    cp /home/kali/div-01/go-reloaded/ztail/quest8T.txt .
fi

OUTPUT1="$(./ztail -c 20 quest8.txt)"
OUTPUT2="$(tail -c 20 quest8.txt)"

if [ "$OUTPUT1" != "$OUTPUT2" ]; then
    echo "FAIL: ./ztail -c 20 quest8.txt"
    PASS=false
    diff  <(echo "$OUTPUT1" ) <(echo "$OUTPUT2")
fi

OUTPUT1="$(./ztail quest8.txt -c 23)"
OUTPUT2="$(tail quest8.txt -c 23)"

if [ "$OUTPUT1" != "$OUTPUT2" ]; then
    echo "FAIL: ./ztail quest8.txt -c 23"
    PASS=false
    diff  <(echo "$OUTPUT1" ) <(echo "$OUTPUT2")
fi

OUTPUT1="$(./ztail -c jelrjq 15)"
OUTPUT1=${OUTPUT1%????}
OUTPUT2="$(tail -c jelrjq 15 2>&1)"
OUTPUT2=${OUTPUT2%????}

if [ "$OUTPUT1" != "$OUTPUT2" ]; then
    echo "FAIL: ./ztail -c jelrjq 15"
    diff  <(echo "$OUTPUT1" ) <(echo "$OUTPUT2")
    PASS=false
fi

OUTPUT1="$(./ztail -c 11 quest8.txt jfdklsa)"
OUTPUT2="$(tail -c 11 quest8.txt jfdklsa 2>&1)"

if [ "$OUTPUT1" != "$OUTPUT2" ]; then
    echo "FAIL: ./ztail -c 11 quest8.txt jfdklsa"
    diff  <(echo "$OUTPUT1" ) <(echo "$OUTPUT2")
    PASS=false
fi

OUTPUT1="$(./ztail 11 -c quest8.txt )"
OUTPUT1=${OUTPUT1%????}
OUTPUT2="$(tail 11 -c quest8.txt  2>&1)"
OUTPUT2=${OUTPUT2%????}

if [ "$OUTPUT1" != "$OUTPUT2" ]; then
    echo "FAIL: ./ztail 11 -c quest8.txt "
    diff  <(echo "$OUTPUT1" ) <(echo "$OUTPUT2")
    PASS=false
fi

if [ "$PASS" = true ] ; then
    echo 'PASS'
fi

