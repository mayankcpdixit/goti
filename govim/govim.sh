extRemoved=$(echo ${1//\.go/})
lCase=$(echo "$extRemoved" | tr '[:upper:]' '[:lower:]' )
snakeCase=$(echo ${lCase// /_})
fileName=$snakeCase.go

if test -f "$fileName"; then
    echo "$fileName already exists."
    vim $fileName
else
    cp $(PWD)/template.go $fileName
    vim $fileName
fi
