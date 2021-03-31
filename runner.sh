#!/bin/sh
FILE="chest-of-drawers"
if [ -e $FILE ]
then
    echo "Found execute file."
    read -sp 'Password : ' pass
    echo
    echo $pass | sudo -S chmod 755 $FILE
    ./$FILE
    echo
else
    echo "nok"
fi