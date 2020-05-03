package _194_转置文件

/*
cat file.txt | awk -v ORS='' '{for(f=1;f<=NF;f++)m[NR,f]=$f} END{for(f=1;f<=NF;f++)for(r=1;r<=NR;r++){print m[r,f];print r==NR?"\n":" "}}'
 */
