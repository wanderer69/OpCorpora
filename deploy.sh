#!/bin/bash
MODE="test"
#MODE="RELEASE"
NAME="OpCorporaService"
if [[ "$OSTYPE" == "linux-gnu"* ]]; then
        # Linux
	if [[ "$MODE" == "RELEASE" ]]; then
		DIR_DEPLOY="/home/user/Deploymented" #Deploy_test or /home/user/Deploymented
	else
		DIR_DEPLOY="/home/user/Deploy_test" #Deploy_test or /home/user/Deploymented
	fi
	FILE_TARGET="/home/user/Go_projects/OpCorpora/service/OpCorporaServer.exe"
	#FILE_CLIENT1="/home/user/Go_projects/Servicer/Client/app.js"
	#FILE_CLIENT2="/home/user/Go_projects/Servicer/Client/app.js.map"
	#FILE_CLIENT3="/home/user/Go_projects/Servicer/Client/index.html"
	FILE_CONFIG="/home/user/Go_projects/MyPageService/settings.json"
	FILE_BLOG="/home/user/Go_projects/MyPageService/BlogDataBase.json"
elif [[ "$OSTYPE" == "darwin"* ]]; then
        # Mac OSX
        exit -4
elif [[ "$OSTYPE" == "cygwin" ]]; then
        # POSIX compatibility layer and Linux environment emulation for Windows
        exit -4
elif [[ "$OSTYPE" == "msys" ]]; then
        # Lightweight shell and GNU utilities compiled for Windows (part of MinGW)
	DIR_DEPLOY="C:/Deploymented" #or /home/user/Deploymented
	FILE_TARGET="C:/Development/Go_projects/OpCorpora/service/OpCorporaServer.exe"
	#FILE_CLIENT1="/home/user/Go_projects/Servicer/Client/app.js"
	#FILE_CLIENT2="/home/user/Go_projects/Servicer/Client/app.js.map"
	#FILE_CLIENT3="/home/user/Go_projects/Servicer/Client/index.html"
	#FILE_CONFIG="C:/Development/Go_projects/OpCorpora/settings.json"
	#FILE_BLOG="C:/Development/Go_projects/MyPageService/BlogDataBase.json"
elif [[ "$OSTYPE" == "win32" ]]; then
        # I'm not sure this can happen.
        exit -4
elif [[ "$OSTYPE" == "freebsd"* ]]; then
        # FreeBSD
        exit -4
else
        # Unknown.
        exit -4
fi
DIR_TARGET=$DIR_DEPLOY/$NAME
DIR_BIN=$DIR_TARGET/bin
#DIR_CLIENT=$DIR_TARGET/clients
#DIR_CONFIG="config"
#FILE_CONFIG_N=$DIR_BIN/"settings.json"
if [ -d "$DIR_DEPLOY" ]; then
  # Take action if $DIR exists. #
  echo "Installing explanatory base trynar files in ${DIR_DEPLOY}..."
else
  mkdir $DIR_DEPLOY
fi
if [ -d "$DIR_TARGET" ]; then
  # Take action if $DIR exists. #
  if [ "$(ls -A $DIR_TARGET)" ]; then
     echo "Take action $DIR_TARGET is not Empty"
     echo "Rename explanatory base trynar files in ${DIR_TARGET}..."
     OLD=$(date +%Y-%m-%d_%s)
     mv ${DIR_TARGET} ${DIR_TARGET}_${OLD}
     echo "Create ${DIR_TARGET}..."
     mkdir $DIR_TARGET
  else
     echo "$DIR_TARGET is Empty"
  fi
else
  echo "Create ${DIR_TARGET}..."
  mkdir $DIR_TARGET
fi
mkdir $DIR_BIN
if [ -f "$FILE_TARGET" ]; then
    echo "$FILE_TARGET exists."
else
    echo "Please make $FILE_TARGET"
    exit -2
fi
cp $FILE_TARGET $DIR_BIN
#cp $FILE_CONFIG $FILE_CONFIG_N
#cp $FILE_BLOG $DIR_BIN
#mkdir $DIR_CLIENT
#if [ -f "$FILE_CLIENT1" ]; then
#    echo "$FILE_CLIENT1 exists."
#else
#    echo "Please make $FILE_CLIENT1"
#    exit -2
#fi
#cp $FILE_CLIENT1 $DIR_CLIENT
#cp $FILE_CLIENT2 $DIR_CLIENT
#cp $FILE_CLIENT3 $DIR_CLIENT
exit 0
