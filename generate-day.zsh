#!/bin/zsh

if [ -n "${SOME_VARIABLE+1}" ];
then
  echo "\$1 is unset";
fi


cp -r ./daytemplate/ days/day$1
