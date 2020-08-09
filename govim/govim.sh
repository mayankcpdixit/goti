[ -z "$1" ] && echo -e "Usage:\n govim filename" && exit 1

extRemoved=$(echo ${1//\.go/})
lCase=$(echo "$extRemoved" | tr '[:upper:]' '[:lower:]' )
snakeCase=$(echo ${lCase// /_})
fileName=$snakeCase.go

if test -f "$fileName"; then
    echo "$fileName already exists."
    vim $fileName
else
    cp $GOVIM_TPL_PATH/template.go $fileName
    vim $fileName
fi
