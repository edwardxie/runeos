#! /bin/env sh
APP="runeos"
ADDR=":8080"

# echo "RuneOS"
# echo "Rune operating system"

# if [[ ${1} == "nb" && -f bin/${APP} ]]; then
#     echo "Runing ${APP}"
#     bin/$APP 
#     # -addr=${ADDR}
#     exit 0
# fi

PWD=`pwd`
# GO_INSTALL_LIST="utils fsnotify golua core ide cmd"
# PKG="${PWD}/pkg"
export RUNEOS_ROOT=${PWD}
export GOPATH=${PWD}:$GOPATH

# goinstall() {
#     pkg="${1}"
#     # if [[ ! -d "${PKG}" ]]; then
#         GOCMD="go install ${APP}/${pkg}"
#         # echo ${GOCMD}
#         ${GOCMD}
#     # fi 
# }


# for pkg in ${GO_INSTALL_LIST}; do
#     goinstall ${pkg}
# done

# if [ -f bin/${APP} ]; then
#     rm bin/${APP}
# fi

# echo "Building ${APP}"
go build -ldflags="-s -w" -o bin/$APP ./src
# echo "GOPATH: "$GOPATH
echo "[${APP} is Running]================================"

# if [[ -f src ]]; then
    # mv ./src ${APP}
    # echo "Runing ${APP} ========================================"
    # echo "[Exec][${APP} ${1} ${2} ${3} ${4} ${5} ${6}]"
    bin/${APP} ${1} ${2} ${3} ${4} ${5} ${6}
    # -addr=${ADDR}
# fi
