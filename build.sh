appName="es-client"
builtAt="$(date +'%F %T %z')"
goVersion=$(go version | sed 's/go version //')
gitAuthor="liusanp"
gitCommit=$(git log --pretty=format:"%h" -1)

if [ "$1" = "dev" ]; then
  version="dev"
else
  version=$(git describe --abbrev=0 --tags)
fi

echo "backend version: $version"

ldflags="\
-w -s \
-X 'github.com/liusanp/es-client/tree/main/commons.BuiltAt=$builtAt' \
-X 'github.com/liusanp/es-client/tree/main/commons.GoVersion=$goVersion' \
-X 'github.com/liusanp/es-client/tree/main/commons.GitAuthor=$gitAuthor' \
-X 'github.com/liusanp/es-client/tree/main/commons.GitCommit=$gitCommit' \
-X 'github.com/liusanp/es-client/tree/main/commons.Version=$version' \
"

FetchWebDev() {
  curl -L https://codeload.github.com/liusanp/es-client-web/tar.gz/refs/heads/dev -o web-dist-dev.tar.gz
  tar -zxvf web-dist-dev.tar.gz
  rm -rf public/dist
  mv -f web-dist-dev/dist public
  rm -rf web-dist-dev web-dist-dev.tar.gz
}

FetchWebRelease() {
  curl -L https://github.com/liusanp/es-client-web/releases/latest/download/dist.tar.gz -o dist.tar.gz
  tar -zxvf dist.tar.gz
  rm -rf public/dist
  mv -f dist public
  rm -rf dist.tar.gz
}

BuildWinArm64() {
  echo building for windows-arm64
  chmod +x ./wrapper/zcc-arm64
  chmod +x ./wrapper/zcxx-arm64
  export GOOS=windows
  export GOARCH=arm64
  export CC=$(pwd)/wrapper/zcc-arm64
  export CXX=$(pwd)/wrapper/zcxx-arm64
  export CGO_ENABLED=1
  go build -o "$1" -ldflags="$ldflags" -tags=jsoniter .
}

BuildDev() {
  go build -o ./build -ldflags="$ldflags" -tags=jsoniter .
}

BuildRelease() {
  rm -rf .git/
  mkdir -p "build"
  BuildWinArm64 ./build/es-client-windows-arm64.exe
  xgo -out "$appName" -ldflags="$ldflags" -tags=jsoniter .
  # why? Because some target platforms seem to have issues with upx compression
  upx -9 ./es-client-linux-amd64
  cp ./es-client-windows-amd64.exe ./es-client-windows-amd64-upx.exe
  upx -9 ./es-client-windows-amd64-upx.exe
  mv es-client-* build
}

MakeRelease() {
  cd build
  mkdir compress
  for i in $(find . -type f -name "$appName-linux-*"); do
    cp "$i" es-client
    tar -czvf compress/"$i".tar.gz es-client
    rm -f es-client
  done
  for i in $(find . -type f -name "$appName-windows-*"); do
    cp "$i" es-client.exe
    zip compress/$(echo $i | sed 's/\.[^.]*$//').zip es-client.exe
    rm -f es-client.exe
  done
  cd compress
  find . -type f -print0 | xargs -0 md5sum >"$1"
  cat "$1"
  cd ../..
}

if [ "$1" = "dev" ]; then
#   FetchWebDev
  BuildDev
elif [ "$1" = "release" ]; then
  FetchWebRelease
  BuildRelease
  MakeRelease "md5.txt"
else
  echo -e "Parameter error"
fi